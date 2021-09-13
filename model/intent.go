package model

// Intent holds properties related to intent entity coming from intent-api
type Intent struct {
	Confidence float64 `json:"confidence"`
	Name       string  `json:"name"`
}
