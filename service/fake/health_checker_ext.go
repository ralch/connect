package fake

type FakeHealthCheckRequest struct {
	Service string `json:"service"`
}

type FakeHealthCheckResponse struct {
	Status string `json:"service"`
}

type FakeHealthErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
