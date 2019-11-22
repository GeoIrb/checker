package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/jmoiron/sqlx"
	yaml "gopkg.in/yaml.v2"
)

//Data данные приложения
type Data struct {
	db      *sqlx.DB
	logging *log.Logger
	Cancel  chan os.Signal
}

//GetPath получить путь до запускаемого файла
func GetPath() string {
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
	fileConfig := fmt.Sprintf("%s.conf", GetPath())

	if _, err := os.Stat(fileConfig); os.IsNotExist(err) {
		log.Println("Config file is not exist")
		return nil
	}

	file, err := ioutil.ReadFile(fileConfig)
	if err != nil {
		log.Fatalln(err)
	}

	var cfg map[string]interface{}
	if err = yaml.Unmarshal(file, &cfg); err != nil || cfg[field] == nil {
		log.Println(err)
		return nil
	}

	return cfg[field].(map[interface{}]interface{})
}

//Init установка начальных соединений
func Init(args ...string) (conn Data) {
	conn = Data{
		db:      connectDB(args...),
		logging: openLog(),
		Cancel:  make(chan os.Signal),
	}

	signal.Notify(conn.Cancel, os.Kill)
	conn.Log("Start\n")
	if conn.db != nil {
		conn.Log("Start\nConnect with DB")
	}

	return
}

//Start ждет прерывания работы
func (conn Data) Start() {
	<-conn.Cancel

	conn.Stop()
}

//Stop закрывает все соединения
func (conn *Data) Stop() {
	if conn.db != nil {
		conn.Log("Disconnect DB\n")

		conn.db.Close()
		conn.db = nil
	}
	conn.Log("Stop\n")
}

//Connect восстанавливает соединение
func (conn *Data) Connect() {
	conn.Log("Connection to DB")
	conn.db = connectDB()
}
