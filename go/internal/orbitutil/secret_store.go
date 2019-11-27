package orbitutil

import (
	"berty.tech/go-orbit-db/iface"
	"berty.tech/go-orbit-db/stores/basestore"
)

type SecretStore interface {
	iface.Store
}

type secretStore struct {
	basestore.BaseStore

	groupContext *GroupContext
}

var _ SecretStore = (*secretStore)(nil)
