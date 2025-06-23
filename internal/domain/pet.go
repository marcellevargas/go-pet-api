package domain

type Pet struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Tutor string `json:"tutor"`
	Photo string `json:"photo"`
}
