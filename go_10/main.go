package main

import (
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql" //第三方提供的mysql驱动包
)

func main() {
	db := openMysql()

	//	id := testInsertData(db)

	//	testUpdateData(db, id)

	testQueryData(db)

	testDeleteData(db, 2)

	db.Close()

}

func openMysql() (db *sql.DB) {
	//"mysql" :mysql的drivername
	//root:root :连接的mysql数据库的用户名和密码
	//tcp(localhost:3306) :连接的mysql的所在服务器的ip和端口
	//GOLANGTEST?charset=utf8 :需要操作的数据库,以及读写编码
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/GOLANGTEST?charset=utf8")

	if err != nil {
		fmt.Println("Open:", err)
	}

	return db
}

//模拟插入数据
func testInsertData(db *sql.DB) (id int64) {
	//插入数据
	stmt, err := db.Prepare("INSERT USER SET name=?,age=?")

	if err != nil {
		fmt.Println("Prepare:", err)
	}

	//传入对应的name和，age值，进行
	res, err := stmt.Exec("xijingping", "62")
	if err != nil {
		fmt.Println("Exec:", err)
	}

	_id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId:", err)
	}

	fmt.Println("插入成功，插入id为:", _id)

	return _id
}

//模拟更新数据
func testUpdateData(db *sql.DB, id int64) {
	//插入数据
	stmt, err := db.Prepare("UPDATE USER SET name=?,age=? where id =?")

	if err != nil {
		fmt.Println("Prepare:", err)
	}

	//传入对应的name和，age值，进行
	res, err := stmt.Exec("xijingping", "65", id)
	if err != nil {
		fmt.Println("Exec:", err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected:", affect)
	}

	fmt.Println("更新成功，更新的条目数：", affect)
}

//模拟查询数据
func testQueryData(db *sql.DB) {
	//查询数据
	rows, err := db.Query("SELECT name,age FROM USER")
	if err != nil {
		fmt.Println("Query:", err)
	}

	for rows.Next() {
		var name string
		var age string
		err = rows.Scan(&name, &age)
		if err != nil {
			fmt.Println("Scan:", err)
		}
		fmt.Println(name)
		fmt.Println(age)
	}
}

//模拟删除数据
func testDeleteData(db *sql.DB, id int64) {
	stmt, err := db.Prepare("DELETE from USER where id=?")
	if err != nil {
		fmt.Println("Prepare:", err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		fmt.Println("Exec:", err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected:", err)
	}

	fmt.Println("更新成功，删除的条目数：", affect)
}
