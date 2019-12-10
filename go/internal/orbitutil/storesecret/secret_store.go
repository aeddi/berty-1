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
	GetDeviceSecret(destMemberPubKey crypto.PubKey, senderDevicePubKey crypto.PubKey) (*group.DeviceSecret, error)

	// SendSecret sends secret of this device to another group member
	SendSecret(ctx context.Context, localDevicePrivKey crypto.PrivKey, remoteMemberPubKey crypto.PubKey, secret *group.DeviceSecret) (operation.Operation, error)
}

type secretStore struct {
	basestore.BaseStore

	groupContext *GroupContext
}

func (s *secretStore) GetDeviceSecret(destMemberPubKey crypto.PubKey, senderDevicePubKey crypto.PubKey) (*group.DeviceSecret, error) {
	key, err := formatSecretMapKey(destMemberPubKey, senderDevicePubKey)
	if err != nil {
		return nil, errcode.TODO
	}

	value := s.Index().Get(key)
	if value == nil {
		return nil, errors.New("unable to get secret for this device")
	}

	casted, ok := value.(*secretsMapEntry)
	if !ok {
		return nil, errors.New("unable to cast interface to map entry")
	}

	if !casted.exists {
		return nil, errors.New("device secret does not exist")
	}

	return casted.secret, nil
}

func (s *secretStore) SendSecret(ctx context.Context, localDevicePrivKey crypto.PrivKey, remoteMemberPubKey crypto.PubKey, secret *group.DeviceSecret) (operation.Operation, error) {
	payload, err := group.NewSecretEntryPayload(localDevicePrivKey, remoteMemberPubKey, secret, s.groupContext.Group)
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
