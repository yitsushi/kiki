package misskey

type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	ID      string `json:"id"`
	Kind    string `json:"kind"`
	Info    struct {
		Param  string `json:"param"`
		Reason string `json:"reason"`
	} `json:"info"`
}