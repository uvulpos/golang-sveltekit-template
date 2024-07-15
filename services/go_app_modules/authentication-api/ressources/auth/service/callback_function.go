package service

import (
	"context"
)

func (s *AuthService) CallbackFunction(authCode string) error {
	authToken, err := s.config.Exchange(context.Background(), authCode)
	if err != nil {
		return err
	}

	client := s.config.Client(context.Background(), authToken)
	client.Get("...")

	return nil
}
