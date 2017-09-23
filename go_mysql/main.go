package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

/*
	go访问mysql数据库
*/
type dict map[string]interface{}

var database = dict{
	"name": "django-simple",
	"host": "localhost",
	"port": 3306,
	"user": "root",
	"password": "bk@321",
}

/* 打印字典 */
func PrintDict(p dict){
	for k,v := range p {
		fmt.Printf("%s: %v\n", k, v)
	}
}

/* 错误检查 */
func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// auth_user
type auth_user struct{
	id int
	password string
	last_login string
	is_superuser int8
	username string
	first_name string
	last_name string
	email string
	is_staff int8
	is_active int8
	date_joined string
}

func main()  {
	PrintDict(database)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		database["user"],
		database["password"],
		database["host"],
		database["port"],
		database["name"])
	fmt.Println("dsn =>", dsn)

	db, err := sql.Open("mysql", dsn)

	// 关闭db连接
	defer db.Close()
	CheckErr(err)

	err = db.Ping(); CheckErr(err)
	fmt.Println("Db Ping Success.")

	// 查询数据
	rows, err := db.Query("SELECT * FROM auth_user WHERE username=? and is_superuser=?", "hongsonggao", 1); CheckErr(err)

	// Get column names
	columns, err := rows.Columns();CheckErr(err)
	fmt.Println(columns, len(columns))

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	// rows.Scan wants '[]interface{}' as an argument
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// 获取数据
	//var user auth_user
	for rows.Next(){
		err = rows.Scan(scanArgs...); CheckErr(err)
		fmt.Println(values)
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
	}

	// 插入数据
	stmt, err := db.Prepare("INSERT auth_user SET " +
		"username=?, " +
		"is_active=?, " +
		"is_staff=?, " +
		"is_superuser=?, " +
		"password=?, " +
		"first_name=?, " +
		"date_joined=?, " +
		"last_login=?, " +
		"last_name=?, " +
		"email=?"); CheckErr(err)
	res, err := stmt.Exec("miya", 1, 1, 1, "miya", "", "2015-01-02", "2015-01-02 02:02:02", "", ""); CheckErr(err)
	id, err := res.LastInsertId(); CheckErr(err)
	fmt.Println(id)


}
