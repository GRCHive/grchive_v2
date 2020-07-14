package fusionauth

import (
	"errors"
	"fmt"
	"html"
)

const loginEndpoint = "/oauth2/authorize"
const logoutEndpoint = "/oauth2/logout"

func (c FusionAuthClient) GetOAuthRedirectEndpoint() string {
	return c.cfg.RedirectDomain + "/oauth2callback"
}

func (c FusionAuthClient) GetOauthLoginEndpoint(state string) string {
	return fmt.Sprintf("%s%s?client_id=%s&response_type=code&redirect_uri=%s&state=%s&scope=%s",
		c.cfg.ExternalHost,
		loginEndpoint,
		c.cfg.ClientId,
		html.EscapeString(c.GetOAuthRedirectEndpoint()),
		state,
		html.EscapeString("openid offline_access"),
	)
}

func (c FusionAuthClient) ExchangeOAuthCodeForAccessToken(code string) (*AccessToken, error) {
	token, oauthError, err := c.FusionAuthClient.ExchangeOAuthCodeForAccessToken(
		code,
		c.cfg.ClientId,
		c.cfg.ClientSecret,
		c.GetOAuthRedirectEndpoint(),
	)

	if err != nil {
		return nil, err
	}

	if oauthError != nil {
		// TODO: #9geac2
		// Implement 2FA and handle the oauth error.
		return nil, errors.New(fmt.Sprintf("[FA OAuth Error] %s - %s [%s]",
			oauthError.Error,
			oauthError.ErrorDescription,
			oauthError.ErrorReason,
		))
	}

	return &AccessToken{token}, nil
}

func (c FusionAuthClient) ExchangeRefreshTokenForAccessToken(refreshToken string) (*AccessToken, error) {
	token, oauthError, err := c.FusionAuthClient.ExchangeRefreshTokenForAccessToken(
		refreshToken,
		c.cfg.ClientId,
		c.cfg.ClientSecret,
		"",
		"",
	)

	if err != nil {
		return nil, err
	}

	if oauthError != nil {
		return nil, errors.New(fmt.Sprintf("[FA OAuth Error] %s - %s [%s]",
			oauthError.Error,
			oauthError.ErrorDescription,
			oauthError.ErrorReason,
		))
	}

	return &AccessToken{token}, nil
}

func (c FusionAuthClient) GetOAuthLogoutUrl() string {
	return fmt.Sprintf("%s%s?client_id=%s&tenantId=%s",
		c.cfg.ExternalHost,
		logoutEndpoint,
		c.cfg.ClientId,
		c.cfg.TenantId,
	)
}
