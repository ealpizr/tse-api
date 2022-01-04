package models

type Record struct {
	ID        string `json:"id,omitempty"`
	FullName  string `json:"fullName,omitempty"`
	BirthDate string `json:"birthDate,omitempty"`
	Age       string `json:"age,omitempty"`
}
