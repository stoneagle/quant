package main

import (
	"evolution/backend/common/config"
	"evolution/backend/system/models"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	configPath := os.Getenv("ConfigPath")
	if configPath == "" {
		configPath = "../../config/.config.yaml"
	}
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	conf := &config.Conf{}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		panic(err)
	}
	exec(conf.System.Database)
}

func exec(dbConfig config.DBConf) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Target)

	engine, err := xorm.NewEngine(dbConfig.Type, source)
	if err != nil {
		panic(err)
	}
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	engine.TZLocation = location
	engine.StoreEngine("InnoDB")
	engine.Charset("utf8")
	if dbConfig.Reset {
		fmt.Printf("reset table\r\n")
		err = engine.DropTables(new(models.User))
		if err != nil {
			panic(err)
		}
	}
	err = engine.Sync2(new(models.User))
	if err != nil {
		panic(err)
	}
}
