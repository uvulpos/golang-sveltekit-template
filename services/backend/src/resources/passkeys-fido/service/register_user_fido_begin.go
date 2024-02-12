package service

import (
	"encoding/json"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/uvulpos/go-svelte/src/helper/webauthn"
)

func (h *PasskeySvc) RegisterUserFidoBegin(uuid string) (*protocol.CredentialCreation, error) {
	user, userErr := h.userService.GetUserByUUID(nil, uuid)
	if userErr != nil {
		return nil, userErr
	}

	webAuthUser := webauthn.CreateWebAuthNUserWithOutCertificates(
		user.Id.String(),
		user.Username,
		user.Username,
	)

	credential, session, err := h.webAuthN.BeginRegistration(webAuthUser)
	if err != nil {
		return nil, err
	}

	sessionJson, sessionJsonErr := json.Marshal(session)
	if sessionJsonErr != nil {
		return nil, sessionJsonErr
	}

	sessionErr := h.storage.InsertFidoSession(nil, uuid, sessionJson, session.Expires)
	if sessionErr != nil {
		return nil, sessionErr
	}

	return credential, nil
}
