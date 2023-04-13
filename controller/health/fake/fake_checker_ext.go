package fake

type FakeCheckRequest struct {
	Service string `json:"service"`
}

type FakeCheckResponse struct {
	Status string `json:"service"`
}

type FakeErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
