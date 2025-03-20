package tool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/liurongdev/firstGoWeb/handle"
	"github.com/liurongdev/firstGoWeb/middleware/logger"
	"net/http"
)

func HelloModule(user *handle.User) string {
	fmt.Println(user.Name)
	return user.Id + user.Name + user.Email + "local version test"
}

func callRemote(url string, req interface{}) interface{} {
	b, _ := json.Marshal(req)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
	defer resp.Body.Close()
	if err != nil {
		logger.Error(err.Error())
	}
	return resp
}
