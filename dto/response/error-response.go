package response

import ()

type ErrorResponse struct {
	Message *string `json:"message,omitempty"`
	Code    *int    `json:"code,omitempty"`
	Reason  *string `json:"reason,omitempty"`
}
