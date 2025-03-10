package model

import (
	"awesomeProject/global"
	"time"
)

type SystemUserInfo struct {
	Id         int       `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id"`
	UserId     int       `gorm:"column:user_id" json:"user_id"`
	UserName   string    `gorm:"column:user_name" json:"user_name"`
	PassWord   string    `gorm:"column:password" json:"password"`
	Phone      string    `gorm:"column:phone" json:"phone"`
	Sex        string    `gorm:"column:sex" json:"sex"`
	CreateBy   string    `gorm:"column:create_by" json:"create_by"`
	UpdateBy   string    `gorm:"column:update_by" json:"update_by"`
	Remark     string    `gorm:"column:remark" json:"remark"`
	Status     int       `gorm:"column:status" json:"status"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;default:NULL" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:NULL" json:"update_time"`
	DeleteTime time.Time `gorm:"column:delete_time;type:timestamp;default:NULL" json:"delete_time"`
}

func (SystemUserInfo) TableName() string {
	return "sys_user_info"
}

func Insert(user SystemUserInfo) (int, error) {
	db := global.GetMysql()
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	result := db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.Id, nil
}

func Update(user SystemUserInfo) int64 {
	db := global.GetMysql()
	user.UpdateTime = time.Now()
	result := db.Model(&user).Updates(user)
	return result.RowsAffected
}

func Delete(user SystemUserInfo) int64 {
	db := global.GetMysql()
	user.DeleteTime = time.Now()
	result := db.Delete(&user)
	return result.RowsAffected
}

func FindById(id int) SystemUserInfo {
	var user SystemUserInfo
	db := global.GetMysql()
	db.Find(&user, "id=?", id)
	return user
}
