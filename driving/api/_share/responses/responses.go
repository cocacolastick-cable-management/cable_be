package responses

type ResponseType struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Payload any    `json:"payload,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

