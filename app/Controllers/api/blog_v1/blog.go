package blog_v1

import (
	"gintest/app/services"
	"gintest/config"
	"gintest/util/bind"
	"gintest/util/error"
	"gintest/util/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type AddArticleForm struct {
	TagID         int    `form:"tag_id" validate:"required,min=1,max=11"`
	Title         string `form:"title" valid:"required;max=20"`
	Desc          string `form:"desc" valid:"required;max=50"`
	Content       string `form:"content" valid:"required;max=100"`
	State         int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Add article
// @Produce  json
// @Param tag_id body int true "TagID"
// @Param title body string true "Title"
// @Param desc body string true "Desc"
// @Param content body string true "Content"
// @Param created_by body string true "CreatedBy"
// @Param state body int true "State"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/articles [post]
func Add(c *gin.Context) {
	var (
		utilGin = response.Gin{Ctx: c}
		form *AddArticleForm
	)
	//参数验证
	validate := validator.New()
	if _, err := bind.Bind(&form, c);err != nil{
		utilGin.Response(config.InvalidParams,error.NewError(err.Error(), "ERROR").Error(), form)
	}
	if err := validate.Struct(form);err != nil{
		utilGin.Response(config.InvalidParams, err.Error(), form)
	}

	tagService := services.Tag{ID: form.TagID}
	exists, err := tagService.ExistByID()
	if err != nil {
		utilGin.Response(config.ERROR, "", err)
		return
	}
	if !exists {
		utilGin.Response(config.ERROR, "tag 不存在", err)
		return
	}

	articleService := services.Article{
		TagID:         form.TagID,
		Title:         form.Title,
		Desc:          form.Desc,
		Content:       form.Content,
		State:         form.State,
	}
	if err := articleService.Add(); err != nil {
		utilGin.Response(config.ERROR, "", err)
		return
	}

	utilGin.Response(config.Success, "", nil)
}


// @Summary Get a single article
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/articles/{id} [get]
func Detail(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		utilGin.Response(config.InvalidParams, "", nil)
	}

	articleService := services.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		utilGin.Response(config.ERROR, "", err)
		return
	}
	if !exists {
		utilGin.Response(config.Success, "tag 不存在", nil)
		return
	}

	article, err := articleService.Get()
	if err != nil {
		utilGin.Response(config.ERROR, "", err)
	}

	utilGin.Response(config.Success, "", article)
}
