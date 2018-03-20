package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Users_20180320_103928 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20180320_103928{}
	m.Created = "20180320_103928"

	migration.Register("Users_20180320_103928", m)
}

// Run the migrations
func (m *Users_20180320_103928) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE users(`id` int(11) NOT NULL AUTO_INCREMENT,`name` int(11) DEFAULT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Users_20180320_103928) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `users`")
}
