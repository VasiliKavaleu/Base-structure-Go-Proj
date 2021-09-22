package controllers

import (
	"goproj/config"
	"goproj/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type detailCategory struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

func GetCategories(c *gin.Context) {
	category := models.Category{}

	categories, err := category.FindAllCategories(config.DB)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})

}

func CreateCategory(c *gin.Context) {
	category := models.Category{}

	// err := c.ShouldBindJSON(&category)
	err := c.ShouldBind(&category)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err = category.SaveCategory(config.DB)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"category": category})
}

func GetCategory(c *gin.Context) {

	categoryId := c.Param("id")
	cid, err := strconv.ParseUint(categoryId, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request",
		})
		return
	}

	category := models.Category{}
	categoryReceived, err := category.FindCategoryById(config.DB, cid)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": categoryReceived,
	})

}

func DeleteCategory(c *gin.Context) {
	categoryId := c.Param("id")
	cid, err := strconv.ParseUint(categoryId, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request",
		})
		return
	}

	category := models.Category{}
	err = config.DB.Debug().Model(models.Category{}).Where("id = ?", cid).Take(&category).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = category.DeleteACategory(config.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Please try again later",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})

}

func UpdateCategory(c *gin.Context) {

	categoryId := c.Param("id")
	cid, err := strconv.ParseUint(categoryId, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request",
		})
		return
	}

	originCategory := models.Category{}
	err = config.DB.Debug().Model(models.Category{}).Where("id = ?", cid).Take(&originCategory).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	category := models.Category{}
	err = c.ShouldBindJSON(&category)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	category.ID = originCategory.ID

	categoryUpdated, err := category.UpdateACategory(config.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Please try again later",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": categoryUpdated,
	})

}
