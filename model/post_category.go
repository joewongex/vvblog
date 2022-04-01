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
