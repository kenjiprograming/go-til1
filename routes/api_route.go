package routes

import (
    "github.com/gin-gonic/gin"
    "til2_go_gin_gorm/controllers"
)

func GetApiRouter() *gin.Engine {
    r := gin.Default()
    // r.LoadHTMLGlob("view/*html")

    r.POST("/signup", controllers.SignUp)
    r.POST("/signin", controllers.SignIn)

    r.GET("/tokenvalid", controllers.TokenValid)
    return r
}
