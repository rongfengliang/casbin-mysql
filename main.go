package main

import (
	"log"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	a, _ := xormadapter.NewAdapter("mysql", "root:dalongrong@tcp(127.0.0.1:3306)/") // Your driver and data source.

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	e, _ := casbin.NewEnforcer("examples/rbac_model.conf", a)

	e.EnableAutoSave(true)
	e.AddPolicy("data2_admin", "demoapp", "read")
	added, err := e.AddGroupingPolicy("dalong", "data2_admin")
	if err != nil {
		log.Println("some wrong", err.Error())
	} else {
		log.Println("add result", added)
	}
	// Check the permission.
	ok, err := e.Enforce("dalong", "demoapp", "read")
	if err != nil {
		log.Println("some wrong", err.Error())
	} else {
		log.Println("has permission", ok)
	}
}
