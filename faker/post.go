package faker

import (
	"github.com/brianvoe/gofakeit"
	"time"
)

type Post struct {
	Id uint8 `json:"id"`
	Author Author `json:"author"`
	PostedAt time.Time `json:"posted_at"`
	Content Content `json:"content"`
}

func NewPost() *Post {
	return &Post{
		Id:       gofakeit.Uint8(),
		Author:   Author{
			Id: gofakeit.Uint8(),
			Name: gofakeit.Name(),
			Username: gofakeit.Username(),
			Image: gofakeit.ImageURL(256, 256),
		},
		PostedAt: time.Now(),
		Content:  Content{
			Text:  gofakeit.Paragraph(1, 2, 10, "."),
			Image: gofakeit.ImageURL(1920, 1080),
		},
	}
}