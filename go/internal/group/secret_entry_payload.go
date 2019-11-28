package group

import (
	"berty.tech/go/pkg/errcode"
	cconv "github.com/agl/ed25519/extra25519"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/crypto/pb"
	"golang.org/x/crypto/nacl/box"
)

const (
	derivationStateSize = 32 // Fixed: see Berty Protocol paper
	counterSize         = 8  // Int64
)

// CheckStructure checks validity of SecretEntryPayload
func (s *SecretEntryPayload) CheckStructure() error {
	_, err := crypto.UnmarshalEd25519PublicKey(s.DestMemberPubKey)
	if err != nil {
		return errcode.TODO.Wrap(err)
	}

	_, err = crypto.UnmarshalEd25519PublicKey(s.SenderDevicePubKey)
	if err != nil {
		return errcode.TODO.Wrap(err)
	}

	if len(s.EncryptedDeviceSecret) < derivationStateSize+counterSize {
		return errcode.TODO.Wrap(err)
	}

	return nil
}

// CheckStructure decrypts and returns encrypted device secret
func (s *SecretEntryPayload) DecryptSecret(localMemberPrivateKey crypto.PrivKey, nonce *[24]byte) (*DeviceSecret, error) {
	destMemberPubKey, err := crypto.UnmarshalEd25519PublicKey(s.DestMemberPubKey)
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	} else if !destMemberPubKey.Equals(localMemberPrivateKey.GetPublic()) {
		// TODO: Custom err to inform caller that this secret isn't for me
		return nil, errcode.TODO.Wrap(err)
	}

	senderDevicePubKey, err := crypto.UnmarshalEd25519PublicKey(s.SenderDevicePubKey)
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	mongPriv, mongPub, err := edwardsToMontgomery(localMemberPrivateKey, senderDevicePubKey)
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	decryptedSecret := &DeviceSecret{}
	decryptedMessage, ok := box.Open(nil, s.EncryptedDeviceSecret, nonce, mongPub, mongPriv)
	if !ok {
		return nil, errcode.TODO
	}

	err = decryptedSecret.Unmarshal(decryptedMessage)
	if err != nil {
		return nil, errcode.TODO
	}

	return decryptedSecret, nil
}

func NewSecretEntryPayload(localDevicePrivKey crypto.PrivKey, remoteMemberPubKey crypto.PubKey, secret *DeviceSecret, nonce *[24]byte) (*SecretEntryPayload, error) {
	remoteMemberPubKeyBytes, err := remoteMemberPubKey.Raw()
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	localDevicePubKeyBytes, err := localDevicePrivKey.GetPublic().Raw()
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	message, err := secret.Marshal()
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	mongPriv, mongPub, err := edwardsToMontgomery(localDevicePrivKey, remoteMemberPubKey)
	if err != nil {
		return nil, errcode.TODO.Wrap(err)
	}

	encryptedSecret := box.Seal(nil, message, nonce, mongPub, mongPriv)

	return &SecretEntryPayload{
		DestMemberPubKey:      remoteMemberPubKeyBytes,
		SenderDevicePubKey:    localDevicePubKeyBytes,
		EncryptedDeviceSecret: encryptedSecret,
	}, nil
}

func edwardsToMontgomery(privKey crypto.PrivKey, pubKey crypto.PubKey) (*[32]byte, *[32]byte, error) {
	var edPriv [64]byte
	var edPub, mongPriv, mongPub [32]byte

	if privKey.Type() != crypto_pb.KeyType_Ed25519 || pubKey.Type() != crypto_pb.KeyType_Ed25519 {
		return nil, nil, errcode.TODO
	}

	privKeyBytes, err := privKey.Raw()
	if err != nil {
		return nil, nil, errcode.TODO.Wrap(err)
	} else if len(privKeyBytes) != 64 {
		return nil, nil, errcode.TODO.Wrap(err)
	}

	pubKeyBytes, err := pubKey.Raw()
	if err != nil {
		return nil, nil, errcode.TODO.Wrap(err)
	} else if len(pubKeyBytes) != 32 {
		return nil, nil, errcode.TODO.Wrap(err)
	}

	copy(edPriv[:], privKeyBytes)
	copy(edPub[:], pubKeyBytes)

	cconv.PrivateKeyToCurve25519(&mongPriv, &edPriv)
	if !cconv.PublicKeyToCurve25519(&mongPub, &edPub) {
		return nil, nil, errcode.TODO.Wrap(err)
	}

	return &mongPriv, &mongPub, nil
}
