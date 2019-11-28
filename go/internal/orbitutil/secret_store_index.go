package orbitutil

import (
	ipfslog "berty.tech/go-ipfs-log"
	"berty.tech/go-orbit-db/iface"
)

type secretEntry struct {
	secret int
}

type secretStoreIndex struct {
	groupContext *GroupContext

	sent     map[string]bool
	received map[string]*secretEntry
}

func (m *secretStoreIndex) Get(key string) interface{} {
	return nil
}

func (m *secretStoreIndex) UpdateIndex(log ipfslog.Log, entries []ipfslog.Entry) error {
	return nil
}

// NewSecretStoreIndex returns a new index to manage the list of the group member secrets
func NewSecretStoreIndex(gc *GroupContext) iface.IndexConstructor {
	return func(publicKey []byte) iface.StoreIndex {
		return &secretStoreIndex{
			groupContext: gc,
		}
	}
}

var _ iface.StoreIndex = &secretStoreIndex{}
