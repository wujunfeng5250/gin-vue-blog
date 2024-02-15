package main

import (
	"encoding/json"
	"fmt"
	"gin-vue-blog/server/models/res"
	"github.com/sirupsen/logrus"
	"os"
)

const file = "server/models/res/err_code.json"

type ErrMap map[res.ErrorCode]string

func main() {
	byteData, err := os.ReadFile(file)
	if err != nil {
		logrus.Error(err)
		return
	}
	var errMap = ErrMap{}
	err = json.Unmarshal(byteData, &errMap)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(errMap)
	fmt.Println(errMap[res.SettingsError])
}
