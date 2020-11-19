package main

import (
	// register driver
	"database/sql"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int64     `json:"id"`         // 自增主键
	Age       int64     `json:"age"`        // 年龄
	FirstName string    `json:"first_name"` // 姓
	LastName  string    `json:"last_name"`  // 名
	Email     string    `json:"email"`      // 邮箱地址
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}

// Connect db by dsn e.g. "user:password@tcp(127.0.0.1:3306)/dbname"
func Connect(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//设置连接池
	conn.SetMaxOpenConns(100)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(10 * time.Minute)
	return conn, conn.Ping()
}

//Query will build a sql
type Query struct {
	db    *sql.DB
	table string
}

//Table bind db and table
func Table(db *sql.DB, tableName string) func() *Query {
	return func() *Query {
		return &Query{
			db:    db,
			table: tableName,
		}
	}
}

// Insert add a line
func (q *Query) Insert(in interface{}) (int64, error) {
	var keys, values []string
	v := reflect.ValueOf(in)
	// 剥离指针
	for v.Kind() == reflect.Ptr {
		v = v.Elem() // &User{} 变成 User{}
	}

	switch v.Kind() {
	case reflect.Struct:
	case reflect.Map:
	case reflect.Slice:
	}
	return 0, nil
}

// sKV struct
func sKV(v reflect.Value) {
	var keys, values []string
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		tf := t.Field(i)
		vf := v.Field(i)

		// 忽略非导出字段
		if tf.Anonymous {
			continue
		}

		// 忽略无效、零值字段
		if !vf.IsValid() || reflect.DeepEqual(vf.Interface(), reflect.Zero(vf.Type())) {
			continue
		}

		if vf.Kind() == reflect.Struct && tf.Type.Name() == "Time" {

		}
	}
}
