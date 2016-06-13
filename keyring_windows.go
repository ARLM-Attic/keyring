package keyring

import (
	"fmt"
	"github.com/SpiderOak/wincred"
)

type windowsProvider struct {
}

func init() {
	defaultProvider = windowsProvider{}
}

func (p windowsProvider) Get(Service, Username string) (string, error) {
	cred, err := wincred.GetGenericCredential(fmt.Sprintf("%s::%s", Service, Username))
	if err != nil {
		return "", err
	}
	return string(cred.CredentialBlob), nil
}

func (p windowsProvider) Set(Service, Username, Password string) error {
	cred := wincred.NewGenericCredential(fmt.Sprintf("%s::%s", Service, Username))
	cred.CredentialBlob = []byte(Password)
	return cred.Write()
}