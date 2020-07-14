package utility

import (
	"errors"
	"github.com/lestrrat-go/jwx/jwt"
	"gitlab.com/grchive/grchive-v2/shared/fusionauth"
	"net/http"
	"time"
)

func ValidateJWT(token string, fa *fusionauth.FusionAuthClient) (jwt.Token, error) {
	// Do a two-step process of validating the JWT:
	// 	1) Validate with FusionAuth. We know for sure that this will check the signature.
	// 	2) Manually verify some of the claims (e.g. expiration, issuer).
	resp, err := fa.ValidateJWT(token)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Non-200 status code for validating JWT with FusionAuth.")
	}

	tkn, err := jwt.ParseString(token)
	if err != nil {
		return nil, err
	}

	err = jwt.Verify(tkn,
		jwt.WithAcceptableSkew(time.Minute*time.Duration(1)),
		jwt.WithIssuer("grchive.com"),
	)

	if err != nil {
		return nil, err
	}

	return tkn, nil
}
