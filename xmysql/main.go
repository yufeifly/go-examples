package main

import (
	"github.com/yufeifly/go-examples/xmysql/dal/db"
)

func main() {
	db.Init()
	//p := model.Person{
	//	Name:     "alan",
	//	Age:      26,
	//	Birthday: time.Now(),
	//}
	//db.CreatePerson(p)
	//db.BatchCreatePeople()
	db.SelectPerson()
}
