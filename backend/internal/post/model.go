package post

import (
	"math/rand"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Uid       string `json:"uid" gorm:"uniqueIndex"`
	Name      string `json:"name"`
	Text      string `json:"text"`
	AuthorUid string `json:"author_uid"`
	TemeUid   string `json:"themeUid"`
	// Files           []*file.File `gorm:"foreignKey:ProductUid;references:Uid"`
}

func NewPost(name, text string) *Post {
	post := &Post{
		Name: name,
		Text: text,
	}
	post.GenerateHash()
	return post
}

func (post *Post) GenerateHash() {
	post.Uid = RandStringRunes(20)
}

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
