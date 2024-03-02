package constants

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"os"
	"sre-tool/internal/config"
	"strings"
)

const (
	ErrorsBasePath   string = "初始化项目根目录失败"
	ErrorInitConfig  string = "初始化配置实例失败"
	ErrorInitLogger  string = "初始化日志实例失败"
	ErrorInitDb      string = "初始化数据库实例失败"
	ErrorInitMongoDb string = "初始化MongoDb实例失败"
	ErrorInitMQ      string = "初始化消息队列实例失败"
	ErrorInitEvent   string = "事件注册执行失败"
	ErrorInitElastic string = "初始化Elastic执行失败"
)

var (
	BasePath string
	Log      *zap.Logger
	Config   *config.Config
	DB       *gorm.DB
)

func init() {
	if curPath, err := os.Getwd(); err == nil {
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(curPath, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = curPath
		}
	} else {
		log.Fatal(ErrorsBasePath)
	}
}
