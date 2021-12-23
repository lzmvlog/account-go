package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Config *AppConfig
)

// AppConfig 系统配置
type AppConfig struct {
	Server struct {
		Port string `yaml:"port"`
	}

	// 数据库配置
	Datasource struct {
		DriverName string `yaml:"driverName"`
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		Database   string `yaml:"database"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		Charset    string `yaml:"charset"`
	}
}

func init() {
	// 所有配置文件实体
	v := AppConfig{}
	// 下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	// JsonParse.Load("../../config.json", &v)
	// 首先获取当前激活是激活那个配置文件
	context, err := ioutil.ReadFile("./application.yml")
	if err != nil {
		log.Printf("读取配置文件错误:%s", err.Error())
		panic(err)
	}
	if err = yaml.Unmarshal(context, &v); err != nil {
		log.Printf("解析配置文件错误:%s", err.Error())
		// panic(err)
	}
	Config = &v
	log.Printf("配置文件信息:%+v", v)

	initLog()
}

// initLog 配置日志
func initLog() {
	// 配置 log
	logfile, err := os.Create("./gin_http.log")
	if err != nil {
		fmt.Println("Could not create log file")
	}
	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = io.MultiWriter(logfile)
}
