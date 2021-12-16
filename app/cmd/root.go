package cmd

import (
	"fiber-root/config"
	"fiber-root/db"
	"fiber-root/pkg/logger"
	"fiber-root/router"
	"fiber-root/util"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"time"
)

var (
	envFile string
	table   string
	rootCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start API fiber-root",
		Example: "fiber-root server config/.env",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&envFile, "env", "e", "config/.env", "Start server with provided configuration file")
	rootCmd.PersistentFlags().StringVarP(&table, "table", "t", "all", "gen [-t tableName],default all tables")

	//go run .\main.go gen -t tableName
	rootCmd.AddCommand(genModel())
}

func run() error {
	//注册路由
	app := router.InitRouter()
	//设置配置
	config.SetConfig(envFile)

	go func() {
		if err := app.Listen(":" + viper.GetString("APP_PORT")); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	fmt.Printf("%s Server Run http://%s:%s/ \r\n", util.GetCurrentTimeStr(),
		viper.GetString("APP_HOST"), viper.GetString("APP_PORT"))

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", util.GetCurrentTimeStr())
	if err := app.Shutdown(); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}
	logger.Info("Server exiting")
	return nil
}

func usage() {
	usageStr := `starting api server`
	log.Printf("%s\n", usageStr)
}

func setup() {
	// 1. 读取配置
	config.ConfigSetup(envFile)
	// 2. 数据库连接
	SetupDB()
}

func Execute() {
	fmt.Println(time.Now())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func SetupDB() {
	//初始化数据库
	db.InitDB()
}
