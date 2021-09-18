package bertybridge

import (
	"berty.tech/berty/v2/go/pkg/bertyaccount"
)

type DecryptedPush bertyaccount.DecryptedPush

func PushDecryptStandalone(rootDir string, inputB64 string) (*DecryptedPush, error) {
	dp, err := bertyaccount.PushDecryptStandalone(rootDir, inputB64)
	return (*DecryptedPush)(dp), err
}
