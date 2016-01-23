package models

import (
	"os"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func Init() {
	orm.RegisterModel(new(Owner))
	orm.RegisterModel(new(Site))
	orm.RegisterModel(new(User))
	orm.RegisterDataBase("default", "postgres", os.Getenv("POSTGRES_URL"), 30)
}