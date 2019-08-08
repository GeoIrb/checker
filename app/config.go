package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	yaml "gopkg.in/yaml.v2"
)

//Config настройки приложения
//Name - имя приложения, совпадает с названием кофнигурациооного файла
type Data struct {
	Name string
	db   *sqlx.DB
	l    logger
}

func GetNameApp() string {
	if strings.Index(os.Args[0], "debug") > -1 {
		return "debug"
	}
	name := strings.Replace(os.Args[0], ".go", "", -1)
	name = strings.Replace(name, ".exe", "", -1)

	return name
}

//Load загрузка настроек из конфигурационного файла
//Имя конфигурационного файла совпадает с именем запускаемого файла
//Расширение конфигурационого файла .conf
func Load(field string) map[interface{}]interface{} {
	fileConfig := fmt.Sprintf("%s.conf", GetNameApp())

	if _, err := os.Stat(fileConfig); os.IsNotExist(err) {
		log.Fatalln("Config file is not exist")
	}

	file, err := ioutil.ReadFile(fileConfig)
	if err != nil {
		log.Fatalln(err)
	}

	var cfg map[string]interface{}
	if err = yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalln(err)
	}

	return cfg[field].(map[interface{}]interface{})
}

//Init начальная настройка приложения
func Init(args ...string) (app Data) {
	app.Name = Load("application")["name"].(string)
	app.db = connectDB(args...)
	app.l = openLog()

	app.Log("App init")
	return
}

//Stop завершает приложение
func (app Data) Pause() {
	app.Log("Pause app")
	app.db.Close()
}

//Stop завершает приложение
func (app Data) Stop() {
	app.Log("Stop app")
	app.db.Close()
}
