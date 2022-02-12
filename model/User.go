package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"` // 0是管理员 1是阅读者
}

// CheckUser 查询用户是否存在
func CheckUser(username string) int {
	var users User
	//fmt.Println("用户名为", username)
	db.Debug().Select("id").Where("username = ?", username).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPassword(data.Password)
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Debug().Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// EditUser 编辑用户信息(密码以外的内容)
func EditUser(id int, data *User) int {
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Debug().Model(&User{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	err = db.Debug().Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// ScryptPassword 密码加密
func ScryptPassword(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 5, 6, 66, 22, 111}
	HashPassword, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	FinalPassword := base64.StdEncoding.EncodeToString(HashPassword)
	return FinalPassword
}

func (u *User) BeforeCreate() {

}
