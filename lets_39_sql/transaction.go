// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/17 - 3:22 下午

// go语言中的数据库事务
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 用户资产表结构体
type assets struct {
	id    int
	name  string
	money int
}

var (
	dbs *sql.DB // 数据库连接池
	url = "root:root@tcp(192.168.64.2:3306)/assets"
)

func init() {
	dbc, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Println("database connection fail.")
		return
	}
	err = dbc.Ping()
	if err != nil {
		fmt.Println("username or password error！")
		return
	}
	dbs = dbc // 返回给连接池
	dbs.SetMaxOpenConns(10)
	//fmt.Println(db)
}
func main() {
	// 为什么要有事务???
	// 通常事务必须满足4个条件（ACID）：原子性（Atomicity，或称不可分割性）、一致性（Consistency）、隔离性（Isolation，又称独立性）、持久性（Durability）。
	tx, err := dbs.Begin()
	if err != nil {
		fmt.Println("开始初始化事务失败!")
		return
	}
	// 执行事务操作  模拟银行转账
	subMoney := `update user_assets set money=money-10000 where id = 1`
	addMoney := `update user_assets set money=money+10000 where id = 2`
	// 开始执行事务
	_, err = tx.Exec(subMoney)
	if err != nil {
		tx.Rollback() // 如果发生错误就回滚到之前的源数据
		fmt.Println("sub money failed.")
		return
	}
	_, err = tx.Exec(addMoney)
	if err != nil {
		tx.Rollback()
		fmt.Println("add money failed.")
		return
	}
	// 如果上面都没有错误就执行commit
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("commit transaction failed.")
		return
	}
	fmt.Println("commit successful.")
}
