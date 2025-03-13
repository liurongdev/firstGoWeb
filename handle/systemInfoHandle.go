package handle

import (
	"awesomeProject/app"
	"awesomeProject/middleware/redis"
	"awesomeProject/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type User struct {
	Id    string `json:"id" binding:"required"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func QuerySystemInfo(c *gin.Context) {
	var user model.SystemUserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "绑定参数错误"})
		return
	}
	userStr, _ := redis.GetKey(strconv.Itoa(user.Id))
	if userStr != "" {
		fmt.Println("get user from redis:" + userStr)
		json.Unmarshal([]byte(userStr), &user)
		app.OK(c, user, "")
		return
	}
	user = model.FindById(user.Id)
	if user.Id != 0 {
		userStr, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
		}

		res := string(userStr)
		fmt.Println(res)
		redis.SetKey(strconv.Itoa(user.Id), res, 50*time.Second)
	}
	app.OK(c, user, "")

}

func InsertSystemUserInfo(c *gin.Context) {
	var user model.SystemUserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		app.ERROR(c, nil, "参数绑定错误", 400)
		return
	}
	id, insertE := model.Insert(user)
	if insertE != nil {
		app.ERROR(c, nil, insertE.Error(), 400)
		return
	}
	app.OK(c, id, "")
}

func UpdateSystemUserInfo(c *gin.Context) {
	var user model.SystemUserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		app.ERROR(c, nil, "参数绑定错误", 400)
		return
	}
	dbUser := model.FindById(user.Id)
	dbUser.UserName = user.UserName
	model.Update(dbUser)
	app.OK(c, dbUser, "success")
}

func DeleteSystemUserById(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		app.ERROR(c, nil, "参数绑定错误", 400)
		return
	}
	idInt, _ := strconv.Atoi(id)
	dbUser := model.FindById(idInt)
	if dbUser.UserName == "" {
		app.ERROR(c, nil, "db用户不存在", 400)
		return
	}
	model.Delete(dbUser)
	app.OK(c, idInt, "")
}
