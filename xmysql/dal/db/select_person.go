package db

import (
	"fmt"
	"github.com/yufeifly/go-examples/xmysql/model"
)

func SelectPerson() {
	var p []model.Person
	result := Db.Table("person").Where("name<>?", "alan").Find(&p)
	fmt.Println(p)
	fmt.Printf("rows affected: %v, err: %v\n", result.RowsAffected, result.Error)
}

func ChildSelectPerson() {
	var p []model.Person
	result := Db.Table("person").Where("age > (?)", Db.Table("person").Select("MIN(age)")).Find(&p)
	//result := Db.Table("person").Find(&p)
	fmt.Println(p)
	fmt.Printf("rows affected: %v, err: %v\n", result.RowsAffected, result.Error)
}
