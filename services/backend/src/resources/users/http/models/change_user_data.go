package models

type ChangeUserDataPayload struct {
	Username string `json:"username" xml:"username" form:"username"`
	Email    string `json:"email" xml:"email" form:"email"`
}

type ChangeUserDataResponseBody struct {
	Status bool `json:"status" xml:"status" form:"status"`
}
