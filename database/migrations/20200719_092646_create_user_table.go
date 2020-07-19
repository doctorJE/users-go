package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateUserTable_20200719_092646 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUserTable_20200719_092646{}
	m.Created = "20200719_092646"

	migration.Register("CreateUserTable_20200719_092646", m)
}

// Run the migrations
func (m *CreateUserTable_20200719_092646) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `user`" + `
		(
		    id       int         UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '編號',
		    account  varchar(50)          NOT NULL                COMMENT '帳號',
		    password varchar(60)          NOT NULL                COMMENT '密碼',
		    PRIMARY KEY (id),
		    UNIQUE INDEX unique_account (account)
		) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='使用者';`)
}

// Reverse the migrations
func (m *CreateUserTable_20200719_092646) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `user`;")
}
