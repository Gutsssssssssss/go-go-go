package api

const (
	ResponseSuccess ResponseStatus = "success"
	ResponseFailed  ResponseStatus = "failed"
)

type ResponseStatus string

type Response struct {
	ID      string         `json:"id"`
	Status  ResponseStatus `json:"status"`
	Message string         `json:"message"`
}
