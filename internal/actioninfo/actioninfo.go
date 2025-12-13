package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) (err error)
	ActionInfo() (datastring string, err error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Print("invalid parse(data)", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Print("output error", err)
		}
		fmt.Println(info)
	}
}
