package swagger

type NotFoundErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
