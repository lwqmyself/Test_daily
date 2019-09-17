package utils

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func InitLogs() (err error) {
	path := "./socket.log"
	config := make(map[string]interface{})
	config["filename"] = path
	config["level"] = logs.LevelDebug
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed ,err:%v", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}
