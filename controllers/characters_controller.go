package controllers

import (
    "strconv"

    "github.com/guswitch/spiderman-api.git/models"
    "github.com/guswitch/spiderman-api.git/database"
    "github.com/gin-gonic/gin"
)

func AllCharacters(c *gin.Context) {
   	db := database.GetDatabase()
	var p []models.Character
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find characters: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func CreateCharacter(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Character

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create character: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func ShowCharacter(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()
	var p models.Character
	err = db.First(&p, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find character by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func DeleteCharacter(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Character{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete character: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func EditCharacter(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Character

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create character: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}