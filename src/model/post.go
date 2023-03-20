package model

import "time"

type Post struct {
	Id        uint64    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserId    uint64    `json:"userId"`
	Likes     uint64    `json:"likes"`
	CreatedAt time.Time `json:"createdAt"`
}
