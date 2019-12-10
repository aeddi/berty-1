package orbitutil

import (
	"fmt"
	"sync"

	ipfslog "berty.tech/go-ipfs-log"
	"berty.tech/go-orbit-db/iface"
	"berty.tech/go/internal/group"
	"berty.tech/go/pkg/errcode"
	"github.com/libp2p/go-libp2p-core/crypto"
)

type secretStoreIndex struct {
	groupContext *GroupContext

	secrets     map[string]*group.DeviceSecret
	muSecrets   sync.RWMutex
	processed   map[string]struct{}
	muProcessed sync.RWMutex
}

type secretsMapEntry struct {
	secret *group.DeviceSecret
	exists bool
}

func formatSecretMapKey(destMemberPubKey crypto.PubKey, senderDevicePubKey crypto.PubKey) (string, error) {
	destBytes, err := destMemberPubKey.Raw()
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	senderBytes, err := senderDevicePubKey.Raw()
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	return fmt.Sprintf("%s-%s", string(destBytes), string(senderBytes)), nil
}

func (m *secretStoreIndex) Get(key string) interface{} {
	ret := &secretsMapEntry{}

	m.muSecrets.RLock()
	ret.secret, ret.exists = m.secrets[key]
	m.muSecrets.RUnlock()

	return ret
}

func (m *secretStoreIndex) UpdateIndex(log ipfslog.Log, entries []ipfslog.Entry) error {
	for _, e := range log.Values().Slice() {
		var err error
		entryHash := e.GetHash().String()

		m.muProcessed.RLock()
		_, ok := m.processed[entryHash]
		m.muProcessed.RUnlock()

		if !ok {
			m.muProcessed.Lock()
			m.processed[entryHash] = struct{}{}
			m.muProcessed.Unlock()

			var entryBytes []byte
			payload := &group.SecretEntryPayload{}

			if entryBytes, err = unwrapOperation(e); err != nil {
				continue
			}

			if err = group.OpenStorePayload(payload, entryBytes, m.groupContext.Group); err != nil {
				continue
			}

			if err = payload.CheckStructure(); err != nil {
				continue
			}

			deviceSecret, err := payload.DecryptSecret(m.groupContext.MemberPrivKey, m.groupContext.Group)
			if err != nil /* && err != notForMe */ {
				continue
			}

			destMemberPubKey, err := crypto.UnmarshalEd25519PublicKey(payload.DestMemberPubKey)
			if err != nil {
				continue
			}

			senderDevicePubKey, err := crypto.UnmarshalEd25519PublicKey(payload.SenderDevicePubKey)
			if err != nil {
				continue
			}

			secretMapKey, err := formatSecretMapKey(destMemberPubKey, senderDevicePubKey)
			if err != nil {
				continue
			}

			m.muSecrets.Lock()
			m.secrets[secretMapKey] = deviceSecret
			m.muSecrets.Unlock()
		}
	}

	return nil
}

// NewSecretStoreIndex returns a new index to manage the list of the group member secrets
func NewSecretStoreIndex(gc *GroupContext) iface.IndexConstructor {
	return func(publicKey []byte) iface.StoreIndex {
		return &secretStoreIndex{
			groupContext: gc,
			processed:    map[string]struct{}{},
			secrets:      map[string]*group.DeviceSecret{},
		}
	}
}

var _ iface.StoreIndex = &secretStoreIndex{}
