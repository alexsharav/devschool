package models

type Course struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Topics      []Topic `json:"topics"`
}

type Topic struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Order   int      `json:"order"`
	Lessons []Lesson `json:"lessons"`
}

type Lesson struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Order   int    `json:"order"`
}
