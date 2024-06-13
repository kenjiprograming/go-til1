package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "til2_go_gin_gorm/services"
    "strings"
)

func TokenValid(c *gin.Context) {

    // アクセストークン取得
    access_token := GetAccessToken(c)
    if access_token == "" {
        MakeJson(c, http.StatusUnauthorized, "Bearerが存在しません。")
        return;
    }

    // アクセストークン検証
    msg, err := services.AccessTokenValidation(access_token)
    if err != nil {
        MakeJson(c, http.StatusUnauthorized, err.Error())
        return;
    }
    MakeJson(c, http.StatusOK, msg)
    return;
}

// 返却用Json作成
func MakeJson(c *gin.Context, code int, msg string) {
    c.JSON(code, gin.H{
        "message" : msg,
    })
}

// Bearerからアクセストークン取得
func GetAccessToken(c *gin.Context) string {
    authorizationHeader := c.Request.Header.Get("Authorization")
    if authorizationHeader != "" {
        ary := strings.Split(authorizationHeader, " ")
        if len(ary) == 2 {
            // Bearer値を解析する
            if ary[0] == "Bearer" {
                return ary[1]
            }
        }
    }
    return ""
}
