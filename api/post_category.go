package api

import (
	"vvblog/errors/vcode"
	"vvblog/errors/verror"
	"vvblog/model"
	"vvblog/service"
	"vvblog/utils"

	"github.com/gin-gonic/gin"
)

var PostCategory = new(postCategoryApi)

type postCategoryApi struct{}

func (p *postCategoryApi) Index(c *gin.Context) {
	var categories []model.PostCategoryListItemRes
	if err := model.DB.Model(&model.PostCategory{}).Order("sort, id DESC").Scan(&categories).Error; err != nil {
		panic(verror.WrapCode(vcode.CodeDbOperationError, err))
	}
	utils.JsonSuccess(c, gin.H{
		"list": categories,
	})
}

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

func (p *postCategoryApi) Update(c *gin.Context) {
	id := c.Param("id")
	service.V.CheckPositiveInt(id, "id")
	var req model.PostCategoryUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(verror.WrapCode(vcode.CodeInvalidParameter, err))
	}

	data := map[string]interface{}{}
	if req.Name != nil {
		data["name"] = req.Name
	}
	if req.Sort != nil {
		data["sort"] = req.Sort
	}

	if len(data) > 0 {
		if err := model.DB.Model(&model.PostCategory{}).Where("id", id).UpdateColumns(data).Error; err != nil {
			panic(verror.WrapCode(vcode.CodeDbOperationError, err))
		}
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
