package db

import (
	"fmt"
	"github.com/yufeifly/go-examples/xmysql/model"
	"time"
)

func CreatePerson(p model.Person) {
	Db.Table("person").Create(&p)
}

func BatchCreatePeople() {
	var users = []model.Person{{Name: "jinzhu1", Birthday: time.Now()},
		{Name: "jinzhu2", Birthday: time.Now()},
		{Name: "jinzhu3", Birthday: time.Now()}}
	Db.Table("person").Create(&users)

	for _, user := range users {
		fmt.Println(user.ID) // 1,2,3
	}
}
