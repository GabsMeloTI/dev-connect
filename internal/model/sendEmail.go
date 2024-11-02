package model

type EmailResponse struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
