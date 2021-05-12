package functions

type APIResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error,omitempty"`
}
