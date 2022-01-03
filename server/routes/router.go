package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/guswitch/spiderman-api.git/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
    main := router.Group("api")
    {
        characters := main.Group("characters")
        {
            characters.GET("/", controllers.AllCharacters)
            characters.GET("/:id", controllers.ShowCharacter)
            characters.POST("/create", controllers.CreateCharacter)
            characters.PUT("/update/:id", controllers.EditCharacter)
            characters.DELETE("/delete/:id", controllers.DeleteCharacter)
        }
    }
    return router
}