package faker

import (
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit"
)

type Post struct {
	Id       uint8     `json:"id"`
	Author   Author    `json:"author"`
	PostedAt time.Time `json:"posted_at"`
	Content  Content   `json:"content"`
}

func NewPost() *Post {
	return &Post{
		Id: gofakeit.Uint8(),
		Author: Author{
			Id:       gofakeit.Uint8(),
			Name:     gofakeit.Name(),
			Username: gofakeit.Username(),
			Image:    gofakeit.ImageURL(getRandImgSize("author")),
		},
		PostedAt: time.Now(),
		Content: Content{
			Text:  gofakeit.Paragraph(1, 3, 10, "."),
			Image: gofakeit.ImageURL(getRandImgSize("content")),
		},
	}
}

func getRandImgSize(imgType string) (w int, h int) {
	switch imgType {
	case "author":
		w, h = rand.Intn(300-200)+200, rand.Intn(300-200)+200
	case "content":
		w, h = rand.Intn(2000-1800)+1800, rand.Intn(1100-1000)+1000
	}

	return w, h
}
