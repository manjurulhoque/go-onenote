package onenote

type PageContent struct {
	Target   string `json:"target"`
	Action   string `json:"action"`
	Position string `json:"position"`
	Content  string `json:"content"`
}
