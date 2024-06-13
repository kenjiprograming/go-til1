package models

type Users struct {
    Id int
    UserId   string
    Password string
}

// 1件取得用
func GetOneUsers(userId string) (data Users) {
    Db.Where("user_id = ?", userId).First(&data)
    return
}

// 1件登録用
func RegistUser(userId string, encryptPw string) {
    user := Users{}
    user = Users{UserId: userId, Password: encryptPw}
    Db.Create(&user)
}
