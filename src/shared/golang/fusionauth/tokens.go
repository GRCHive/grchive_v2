package fusionauth

import (
	"github.com/FusionAuth/go-client/pkg/fusionauth"
)

type AccessToken struct {
	*fusionauth.AccessToken
}

func (a AccessToken) Token() string {
	return a.AccessToken.AccessToken
}
