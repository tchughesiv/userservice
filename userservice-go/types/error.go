package types

type Error struct {
	Error string `json:"error,omitempty"`

	ErrorDescription string `json:"error_description,omitempty"`
}
