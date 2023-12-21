package domain

type Response struct {
	Data        any    `json:"data"`
	Message     string `json:"message"`
	ElapsedTime string `json:"elapsed_time"`
}
