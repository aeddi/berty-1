package orbitutil

import (
	"context"
	"errors"

	"berty.tech/go-orbit-db/iface"
	"berty.tech/go-orbit-db/stores/basestore"
	"berty.tech/go-orbit-db/stores/operation"
	"berty.tech/go/internal/group"
	"berty.tech/go/pkg/errcode"
	"github.com/libp2p/go-libp2p-core/crypto"
)

type SecretStore interface {
	iface.Store

	// GetDeviceSecret gets secret device
	GetDeviceSecret(remoteMemberPubKey crypto.PubKey) (*group.DeviceSecret, error)

	// SendSecret sends secret of this device to another group member
	SendSecret(ctx context.Context, localDevicePrivKey crypto.PrivKey, remoteMemberPubKey crypto.PubKey, secret *group.DeviceSecret) (operation.Operation, error)
}

type secretStore struct {
	basestore.BaseStore

	groupContext *GroupContext
}

func (s *secretStore) GetDeviceSecret(remoteDevicePubKey crypto.PubKey) (*group.DeviceSecret, error) {
	bytes, err := remoteDevicePubKey.Raw()
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	value := s.Index().Get(string(bytes))
	if value == nil {
		return nil, errors.New("unable to get secret for this device")
	}

	secret, ok := value.(*group.DeviceSecret)
	if !ok {
		return nil, errors.New("unable to cast entry to secret")
	}

	return secret, nil
}

func (s *secretStore) SendSecret(ctx context.Context, localDevicePrivKey crypto.PrivKey, remoteMemberPubKey crypto.PubKey, secret *group.DeviceSecret) (operation.Operation, error) {
	// Nonce doesn't need to be secret, random nor unpredictable, it just needs
	// to be used only once for a given {sender, receiver} set and we will send
	// only one SecretEntryPayload per {localDevicePrivKey, remoteMemberPubKey}
	// So we can reuse groupID as nonce for all SecretEntryPayload and save
	// 24 bytes of storage and bandwidth for each of them.
	//
	// See https://pynacl.readthedocs.io/en/stable/secret/#nonce
	// See Security Model here: https://nacl.cr.yp.to/box.html
	var nonce [24]byte

	gid, err := s.groupContext.Group.GroupIDAsString()
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	copy(nonce[:], []byte(gid))

	payload, err := group.NewSecretEntryPayload(localDevicePrivKey, remoteMemberPubKey, secret, &nonce)
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	env, err := group.SealStorePayload(payload, s.groupContext.Group, localDevicePrivKey)
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	op := operation.NewOperation(nil, "ADD", env)

	e, err := s.AddOperation(ctx, op, nil)
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	op, err = operation.ParseOperation(e)
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	return op, nil
}

var _ SecretStore = (*secretStore)(nil)
