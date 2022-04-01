package api

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"vvblog/config"
	"vvblog/errors/vcode"
	"vvblog/errors/verror"
	"vvblog/model"
	"vvblog/service"
	"vvblog/utils"

	"github.com/ahmetb/go-linq/v3"
	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gorm.io/gorm"
)

var Post = postApi{}

type postApi struct{}

func (p *postApi) Upload(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		panic(verror.WrapCode(vcode.CodeInvalidRequest, err))
	}

	file, err := fileHeader.Open()
	if err != nil {
		panic(verror.WrapCode(vcode.CodeInternalError, err))
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		panic(verror.WrapCode(vcode.CodeInternalError, err))
	}

	md5Hash := fmt.Sprintf("%x", md5.Sum(fileContent))

	baseUrl := fmt.Sprintf("https://%s.cos.%s.myqcloud.com", config.COS.Bucket, config.COS.Region)
	u, err := url.Parse(baseUrl)
	if err != nil {
		panic(verror.WrapCode(vcode.CodeInternalError, err))
	}

	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.COS.SecretID,
			SecretKey: config.COS.SecretKey,
		},
	})

	key := "posts/" + md5Hash + filepath.Ext(fileHeader.Filename)
	reader := bytes.NewReader(fileContent)
	_, err = client.Object.Put(context.Background(), key, reader, nil)
	if err != nil {
		panic(verror.WrapCode(vcode.CodeInternalError, err))
	}

	utils.JsonSuccess(c, gin.H{
		"url": baseUrl + "/" + key,
	})
}

func (p *postApi) Create(c *gin.Context) {
	var r model.CreatePostReq
	if err := c.ShouldBindJSON(&r); err != nil {
		panic(verror.WrapCode(vcode.CodeInvalidParameter, err))
	}

	categories := []model.PostCategory{}
	for _, id := range r.CategoryIds {
		categories = append(categories, model.PostCategory{Model: model.Model{ID: id}})
	}

	var postedAt time.Time
	if *r.Draft == 0 {
		postedAt = time.Now()
	}
	if err := model.DB.Create(&model.Post{
		Title:      r.Title,
		Content:    r.Content,
		Draft:      *r.Draft,
		Categories: categories,
		PostedAt:   postedAt,
	}).Error; err != nil {
		panic(verror.WrapCode(vcode.CodeDbOperationError, err))
	}

	utils.JsonSuccess(c)
}

func (p *postApi) List(c *gin.Context) {
	var r model.ListPostReq
	if err := c.ShouldBindQuery(&r); err != nil {
		panic(verror.WrapCode(vcode.CodeInvalidParameter, err))
	}

	var posts []model.ListPostRes
	var count int64
	var postIDS []uint
	offset := (r.Page - 1) * r.PageSize
	fmt.Println(offset)

	query := model.DB.Model(&model.Post{})
	if r.Keyword != "" {
		query.Where("title like ?", fmt.Sprintf("%%%s%%", r.Keyword))
	}

	// Session 复制一个新的statement，但是保留查询条件
	if err := query.Session(&gorm.Session{}).Count(&count).Error; err != nil {
		panic(verror.WrapCode(vcode.CodeDbOperationError, err))
	}

	if err := query.Offset(int(offset)).Limit(int(r.PageSize)).Scan(&posts).Error; err != nil {
		panic(verror.WrapCode(vcode.CodeDbOperationError, err))
	}

	linq.From(posts).Select(func(i interface{}) interface{} {
		return i.(model.ListPostRes).ID
	}).ToSlice(&postIDS)

	type PostCategoryNames struct {
		PostID uint
		Names  string
	}
	var categories []PostCategoryNames
	if err := model.DB.Table("post_2_category AS pc").
		Joins("JOIN post_category c ON pc.post_category_id = c.id").
		Where("pc.post_id IN ?", postIDS).
		Select("pc.post_id, GROUP_CONCAT(c.name) names").
		Group("pc.post_id").
		Find(&categories).Error; err != nil {
		panic(verror.NewCode(vcode.CodeDbOperationError, err.Error()))
	}

	namesMap := map[uint][]string{}
	linq.From(categories).ToMapBy(&namesMap, func(i interface{}) interface{} {
		return i.(PostCategoryNames).PostID
	}, func(i interface{}) interface{} {
		return strings.Split(i.(PostCategoryNames).Names, ",")
	})

	for i, post := range posts {
		posts[i].Categories = namesMap[post.ID]
	}

	utils.JsonSuccess(c, gin.H{
		"list":  posts,
		"count": count,
	})
}

func (p *postApi) Show(c *gin.Context) {
	id := c.Param("id")
	service.V.CheckPositiveInt(id, "id")

	var post model.PostDetailRes
	err := model.DB.Table("post AS p").
		Joins("JOIN post_2_category pc ON pc.post_id = p.id").
		Where("p.id = ?", id).
		Select("p.id, p.title, p.content, p.draft, GROUP_CONCAT(pc.post_category_id) category_ids_str").
		Group("p.id").
		Scan(&post).
		Error

	if err != nil {
		panic(verror.WrapCode(vcode.CodeDbOperationError, err))
	}

	if post.ID > 0 {
		ids := strings.Split(post.CategoryIdsStr, ",")
		for _, id := range ids {
			idInt, _ := strconv.Atoi(id)
			post.CategoryIds = append(post.CategoryIds, uint(idInt))
		}
	}

	utils.JsonSuccess(c, gin.H{"post": post})
}

func (p *postApi) Update(c *gin.Context) {
	id := c.Param("id")
	service.V.CheckPositiveInt(id, "id")

	var r model.UpdatePostReq
	if err := c.ShouldBindJSON(&r); err != nil {
		panic(verror.WrapCode(vcode.CodeInvalidParameter, err))
	}

	data := map[string]interface{}{}
	var categoryIds []uint
	if r.Title != nil {
		data["title"] = r.Title
	}
	if r.Content != nil {
		data["content"] = r.Content
	}
	if r.CategoryIds != nil {
		categoryIds = r.CategoryIds
	}
	if r.Draft != nil {
		data["draft"] = r.Draft
	}

	err := model.DB.Transaction(func(tx *gorm.DB) error {
		if len(data) > 0 {
			if err := tx.Model(&model.Post{}).Where("id", id).Updates(data).Error; err != nil {
				return err
			}
		}
		if categoryIds != nil {
			if err := tx.Table("post_2_category").Where("post_id", id).Delete(nil).Error; err != nil {
				return err
			}

			if len(categoryIds) > 0 {
				insert := []map[string]interface{}{}
				for _, cid := range categoryIds {
					insert = append(insert, map[string]interface{}{"post_id": id, "post_category_id": cid})
				}
				if err := tx.Table("post_2_category").Create(&insert).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		panic(verror.WrapCode(vcode.CodeDbOperationError, err))
	}

	utils.JsonSuccess(c)
}
