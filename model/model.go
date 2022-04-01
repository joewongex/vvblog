package model

import (
	"database/sql"
	"fmt"
	"time"
)

type Model struct {
	ID        uint         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type Paginator struct {
	Page     uint `json:"page" form:"page" binding:"required,min=1"`
	PageSize uint `json:"page_size" form:"page_size" binding:"required,min=1,max=50"`
}

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf(`"%s"`, time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
