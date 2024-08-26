package httpModels

import (
	"encoding/json"

	"github.com/uvulpos/go-svelte/basic-utils/customerrors"
)

type SelfPermissions struct {
	Permissions []string `json:"permissions"`
}

func NewSelfPermissions(permissions []string) *SelfPermissions {
	return &SelfPermissions{
		Permissions: permissions,
	}
}

func (s *SelfPermissions) ToJson() (string, customerrors.ErrorInterface) {
	json, jsonErr := json.Marshal(s)
	if jsonErr != nil {
		return "", customerrors.NewInternalServerError(jsonErr, "", "could not convert SelfPermissions http model into json")
	}

	return string(json), nil
}
