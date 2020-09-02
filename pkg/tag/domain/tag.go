package domain

type Tag struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// Tags slice de punteros a tag
type Tags []*Tag
