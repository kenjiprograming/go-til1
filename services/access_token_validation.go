package services

import (
    "til2_go_gin_gorm/models"
    "til2_go_gin_gorm/libraries/crypto"
    "github.com/golang-jwt/jwt/v5"
    "fmt"
    "time"
)

var Layout = "2006-01-02 15:04:05"

func AccessTokenValidation(access_token string) (string, error) {

    token, err := jwt.Parse(access_token, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }

        return []byte("ACCESS_SECRET_KEY"), nil
    })

    var msg string

    // エラーチェック
    if err != nil {
        msg = "tokenエラーです。"
        return "", fmt.Errorf(msg)
    }

     // クレームの取得
     if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
         userId := string(claims["user_id"].(string))
         pw := string(claims["password"].(string) )
         exp := int64(claims["exp"].(float64))

         // ユーザー情報の取得
         data := models.GetOneUsers(userId)
         if data.Id == 0 {
             return "", fmt.Errorf("ユーザーが存在しません。")
         }

         // パスワードの検証
         err := crypto.CompareHashAndPassword(data.Password, pw)
         if err != nil {
             return "", fmt.Errorf("パスワードが一致しません。%s",pw)
         }

         return "トークンの検証に成功しました 。 有効期限 : " + time.Unix(exp, 0).Format(Layout), nil
     } else {
         return "",  fmt.Errorf("クレームの取得に失敗しました。")
     }
     return "", nil
}
