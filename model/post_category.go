package model

type PostCategory struct {
	Model
	Name  string `gorm:"type:varchar(10);not null" json:"name"`
	Sort  uint8  `json:"sort"`
	Posts []Post `gorm:"many2many:post_2_category"`
}

type PostCategoryCreateReq struct {
	Name string `json:"name" binding:"required,max=10"`
	Sort *uint8 `json:"sort" binding:"required,min=0,max=255"`
}

type PostCategoryUpdateReq struct {
	Name *string `json:"name" binding:"omitempty,max=10"`
	Sort *uint8  `json:"sort" binding:"omitempty,min=0,max=255"`
}

type PostCategoryListItemRes struct {
	Id        uint     `json:"id"`
	Name      string   `json:"name"`
	Sort      uint8    `json:"sort"`
	CreatedAt JSONTime `json:"created_at"`
	UpdatedAt JSONTime `json:"updated_at"`
}
