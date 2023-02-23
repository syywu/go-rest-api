package handlers

type ErrResponse struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
	StatusText string `json:"status_text`
	Message    string `json:"message`
}

var (
	ErrMethodNotAllowed = &ErrResponse{StatusCode: 405, Message: "Method not allowed"}
	ErrNotFound         = &ErrResponse{StatusCode: 404, Message: "Not found"}
	ErrBadRequest       = &ErrResponse{StatusCode: 400, Message: "Bad request"}
)
