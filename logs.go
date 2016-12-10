// logs project logs.go
package logs

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cihub/seelog"
)

var Logger seelog.LoggerInterface

// DisableLog disable all library log output
func DisableLog() {
	Logger = seelog.Disabled
}

func init() {
	DisableLog()
	loadAppConfig()
}

// UseLogger uses a specified seelog.LoggerInterface to output library log.
// Use this func if you are using Seelog logging system in your app.
func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}

// 加载配置文件
func loadAppConfig() {
	logger, err := seelog.LoggerFromConfigAsFile(getCurrentPath() + "appConfig.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}
