package api

import (
	"vvblog/errors/vcode"
	"vvblog/errors/verror"
	"vvblog/model"
	"vvblog/utils"

	"github.com/gin-gonic/gin"
)

var PostCategory = new(postCategoryApi)

type postCategoryApi struct{}

func (p *postCategoryApi) Create(c *gin.Context) {
	var req model.PostCategoryCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(verror.WrapCode(vcode.CodeInvalidParameter, err))
	}

	var count = new(int64)
	if err := model.DB.Model(&model.PostCategory{}).Where("name = ?", req.Name).Count(count).Error; err != nil {
		panic(verror.WrapCode(vcode.CodeDbOperationError, err))
	}
	if *count > 0 {
		panic(verror.NewCode(vcode.CodeBusinessValidationFailed, "分类名称已存在"))
	}

	if err := model.DB.Create(&model.PostCategory{Name: req.Name, Sort: *req.Sort}).Error; err != nil {
		panic(verror.WrapCode(vcode.CodeDbOperationError, err))
	}

	utils.JsonSuccess(c)
}

func (p *postCategoryApi) Options(c *gin.Context) {
	var options []model.PostCategory
	if err := model.DB.Select("id, name").Find(&options).Error; err != nil {
		panic(verror.WrapCode(vcode.CodeDbOperationError, err))
	}

	utils.JsonSuccess(c, gin.H{
		"list": options,
	})
}
