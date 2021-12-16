package cmd

import (
	"fiber-root/config"
	"fiber-root/db"
	"fiber-root/util"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// gen生成model
func genModel() *cobra.Command {
	return &cobra.Command{
		Use:   "gen",
		Short: "gen生成model",
		Long:  `gen生成model`,
		Run: func(cmd *cobra.Command, args []string) {
			config.ConfigSetup(envFile)
			util.ExecCmd("go get github.com/sliveryou/grom")

			if table == "all" {
				genAllModel()
			} else {
				_genModel(table)
			}
		},
	}
}

func _genModel(table string) {
	cmdstr := "go run github.com/sliveryou/grom convert -H %s -P %s -u %s -p %s -d %s -t %s -e INITIALISM,FIELD_COMMENT,JSON_TAG,GORM_V2_TAG -o app/model/%s.go"
	cmdstr = fmt.Sprintf(cmdstr,
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USERNAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_DATABASE"), table, table)
	util.ExecCmd(cmdstr)
}

func genAllModel() {
	rows, _ := db.GetDB().Raw("show tables").Rows()
	table := ""
	for rows.Next() {
		rows.Scan(&table)
		_genModel(table)
	}
}
