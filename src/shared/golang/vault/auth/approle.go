package auth

import (
	"bytes"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

type AppRoleTokenManager struct {
	RoleId   string
	SecretId string

	token string
}

type appRoleLoginResponse struct {
	Auth struct {
		ClientToken string `json:"client_token"`
	} `json:"auth"`
}

func (m *AppRoleTokenManager) Initialize(url string, client *http.Client) {
	jsonBuffer, err := json.Marshal(map[string]string{
		"role_id":   m.RoleId,
		"secret_id": m.SecretId,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create JSON body for Vault AppRole auth.")
	}
	body := bytes.NewBuffer(jsonBuffer)

	req, err := http.NewRequest("POST", url+"/v1/auth/approle/login", body)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create HTTP request for Vault AppRole auth.")
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to execute HTTP request for Vault AppRole auth.")
	}
	defer resp.Body.Close()

	respBodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed Vault AppRole auth login body.")
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal().Msgf("[%d] Failed to auth using AppRole with Vault: %s", resp.StatusCode, string(respBodyData))
	}

	parsedData := appRoleLoginResponse{}
	err = json.Unmarshal(respBodyData, &parsedData)
	if err != nil {
		log.Fatal().Err(err).Msg("Vault AppRole auth login body is not in expected form.")
	}

	m.token = parsedData.Auth.ClientToken
}

func (m *AppRoleTokenManager) Token() string {
	return m.token
}
