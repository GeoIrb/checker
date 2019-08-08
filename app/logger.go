package app

import (
	"fmt"
	"log"
	"os"
)

type logger struct {
	proc *log.Logger
	err  *log.Logger
}

func openLog() logger {
	logger := logger{
		proc: log.New(os.Stdout, "", log.Ldate|log.Ltime),
		err:  log.New(os.Stdout, "", 0),
	}

	logFile := fmt.Sprintf("%s.log", GetNameApp())
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		if _, err := os.Create(logFile); err != nil {
			log.Println(err)
			return logger
		}
	}

	errFile := fmt.Sprintf("%s_err.log", GetNameApp())
	if _, err := os.Stat(errFile); os.IsNotExist(err) {
		if _, err := os.Create(errFile); err != nil {
			log.Println(err)
			return logger
		}
	}

	logFD, _ := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0600)
	errFD, _ := os.OpenFile(errFile, os.O_APPEND|os.O_WRONLY, 0600)
	logger.proc.SetOutput(logFD)
	logger.err.SetOutput(errFD)

	return logger
}

//In пишет в файл лог файл
func (app Data) Log(mess string, arg ...interface{}) {
	mess = fmt.Sprintf(mess, arg...)

	fmt.Println(mess)
	app.l.proc.Println(mess)
}

//Err логирование ошибок
func (app Data) Err(mess string, arg ...interface{}) {
	mess = fmt.Sprintf(mess, arg...)

	fmt.Println(mess)
	app.l.err.Println(mess)
}

//Completion write info about ending function
func (app Data) Completion(mess string, arg ...interface{}) {
	if r := recover(); r != nil {
		app.Err(mess, r.(string))
	} else {
		app.Log("End %s", fmt.Sprintf(mess, arg...))
	}
}
