package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "til2_go_gin_gorm/services"
)

type JsonRequest struct {
    UserId  string `json:"user_id"`
    Password  string    `json:"password"`
}

func SignUp(c *gin.Context) {
    // リクエストの解析
    var json JsonRequest
    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error" : err.Error(),
        })
        return
    }
    userId := json.UserId
    pw := json.Password

    // サインアップ処理の実行
    if err := services.SignUp(userId, pw); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error" : err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message" : "サインアップ成功",
    })
}

func SignIn(c *gin.Context) {

    // リクエストの解析
    var json JsonRequest
    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error" : err.Error(),
        })
        return
    }
    userId := json.UserId
    pw := json.Password

    // ログイン処理の実行
    token, err := services.SignIn(userId, pw);
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error" : err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message" : "ログイン成功",
        "accessToken" : token,
    })
}
