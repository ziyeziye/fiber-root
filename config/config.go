package config

import (
	"fiber-root/pkg/logger"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

//载入配置文件
func ConfigSetup(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Read config file[%s] fail: %s", path, err.Error()))
	}

	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		logger.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}

	// 日志配置
	logger.Init()
}

func SetConfig(envFile string) {
	viper.SetConfigFile(envFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.AutomaticEnv()

	loc, _ := time.LoadLocation(viper.GetString("APP_LOCATION"))
	// handle err
	time.Local = loc // -> this is setting the global timezone
}
