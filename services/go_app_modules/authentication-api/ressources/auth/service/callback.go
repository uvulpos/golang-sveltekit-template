package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

func (s *AuthService) CallbackFunction(authCode, state, oauthUserinfoURL string) (*AuthentikUserInfo, error) {

	authToken, err := s.config.Exchange(context.Background(), authCode)
	if err != nil {
		fmt.Println("Exchange ERROR")
		return nil, err
	}

	client := s.config.Client(context.Background(), authToken)
	resp, err := client.Get(s.oauthUserInfoEP)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result *AuthentikUserInfo
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
