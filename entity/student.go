package entity

type Student struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Age    uint     `json:"age"`
	Grade  float64  `json:"grade"`
	Course []string `json:"course"`
}
