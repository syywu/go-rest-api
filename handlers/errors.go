package handlers

type ErrResponse struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
	StatusText string `json:"status_text`
	Message    string `json:"message`
}
