package services

import (
    "til2_go_gin_gorm/models"
    "til2_go_gin_gorm/libraries/crypto"
    "time"
    "github.com/golang-jwt/jwt/v5"
    "fmt"
)

func SignUp(userId string, password string) (error) {

    // ユーザー情報の取得
    data := models.GetOneUsers(userId)
    if data.Id != 0 {
        return fmt.Errorf("そのUserIdは既に登録されています。")
    }

    // パスワードの暗号化
    encryptPw, err := crypto.PasswordEncrypt(password)
    if err != nil {
        return fmt.Errorf("パスワードの暗号化でエラーが発生しました。")
    }

    // DB登録
    models.RegistUser(userId, encryptPw)

    return nil
}

func SignIn(userId string, password string) (string, error) {

    // ユーザー情報の取得
    data := models.GetOneUsers(userId)
    if data.Id == 0 {
        return "", fmt.Errorf("ユーザーが存在しません。")
    }

    // パスワードの検証
    err := crypto.CompareHashAndPassword(data.Password, password)
    if err != nil {
        return "", fmt.Errorf("パスワードが一致しません。")
    }

    // JWTに付与する構造体
    var limit time.Duration = time.Hour * 24 // トークンの有効期限を24時間とする

    claims := jwt.MapClaims{
        "user_id": userId,
        "password": password,
        "exp": time.Now().Add(limit).Unix(),
    }
    // ヘッダーとペイロード生成
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // トークンに署名を付与
    accessToken, _ := token.SignedString([]byte("ACCESS_SECRET_KEY"))

    return accessToken, nil
}
