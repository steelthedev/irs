package data

type AppHttpErr struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *AppHttpErr) Error() string {
	return e.Message
}

type AppErr struct {
	Message string
}

func (e *AppErr) Error() string {
	return e.Message
}
