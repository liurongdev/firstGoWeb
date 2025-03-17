package tool

import (
	"fmt"
	"github.com/liurongdev/firstGoWeb/handle"
)

func HelloModule(user *handle.User) string {
	fmt.Println(user.Name)
	return user.Id + user.Name + user.Email + "local version test"
}
