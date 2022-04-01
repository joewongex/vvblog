package model

import (
	"time"
)

type Post struct {
	Model
	Title      string         `gorm:"type:varchar(100);not null" json:"title"`
	Content    string         `gorm:"type:text;not null" json:"content"`
	Draft      uint8          `gorm:"type:tinyint unsigned not null" json:"draft"`
	Categories []PostCategory `gorm:"many2many:post_2_category"`
	PostedAt   time.Time      `json:"posted_at"`
}

type CreatePostReq struct {
	Title       string `json:"title" binding:"required,max=100"`
	Content     string `json:"content"`
	CategoryIds []uint `json:"category_ids" binding:"dive,min=1"`
	Draft       *uint8 `json:"draft" binding:"required,min=0,max=1"`
}

type UpdatePostReq struct {
	Title       *string `json:"title" binding:"max=100"`
	Content     *string `json:"content"`
	CategoryIds []uint `json:"category_ids" binding:"dive,min=1"`
	Draft       *uint8 `json:"draft" binding:"min=0,max=1"`
}

type ListPostReq struct {
	Paginator
	Keyword string `form:"keyword"`
}

type ListPostRes struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	CreatedAt  JSONTime  `json:"created_at"`
	UpdatedAt  JSONTime  `json:"updated_at"`
	PostedAt   *JSONTime `json:"posted_at"`
	Draft      uint8     `json:"draft"`
	Categories []string  `json:"categories"`
}

type PostDetailRes struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	Content        string `json:"content"`
	Draft          uint8  `json:"draft"`
	CategoryIdsStr string `json:"-"`
	CategoryIds    []uint `json:"category_ids"`
}
