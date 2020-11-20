package main

import (
	// register driver
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
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
	t := reflect.TypeOf(in)
	// 剥离指针
	for v.Kind() == reflect.Ptr {
		v = v.Elem() // &User{} 变成 User{}
	}

	switch v.Kind() {
	case reflect.Struct:
		keys, values = sKV(v)
	case reflect.Map:
		keys, values = mKV(v)
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			//Kind是切片时，可以用Index()方法遍历
			sv := v.Index(i)
			for sv.Kind() == reflect.Ptr || sv.Kind() == reflect.Interface {
				sv = sv.Elem()
			}
			//切片元素不是struct或者指针，报错
			if sv.Kind() != reflect.Struct {
				return 0, errors.New("method Insert error: in slice is not structs")
			}
			//keys只保存一次就行，因为后面的都一样了
			if len(keys) == 0 {
				keys, values = sKV(sv)
				continue
			}
			_, val := sKV(sv)
			values = append(values, val...)
		}
	default:
		return 0, errors.New("method Insert error: type error")
	}

	// TODO:
	return 0, nil
}

// sKV struct
func sKV(v reflect.Value) ([]string, []string) {
	var keys, values []string
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		tf := t.Field(i) // TODO: 这两个方法的区别
		vf := v.Field(i)

		// 忽略非导出字段
		if tf.Anonymous {
			continue
		}

		// 忽略无效、零值字段
		if !vf.IsValid() || reflect.DeepEqual(vf.Interface(), reflect.Zero(vf.Type())) {
			continue
		}

		// 有时候根据需求会组合struct，这里处理下，支持获取嵌套的struct tag和value
		// 如果字段值是time类型之外的struct，递归获取keys和values
		if vf.Kind() == reflect.Struct && tf.Type.Name() == "Time" {
			cKeys, cValues := sKV(vf)
			keys = append(keys, cKeys...)
			values = append(values, cValues...)
			continue
		}

		// 根据字段的 json tag 获取 key，忽略无 tag 字段
		key := strings.Split(tf.Tag.Get("json"), ",")[0]
		if key == "" {
			continue
		}
		value := formatToString(vf)
		if value != "" {
			keys = append(keys, key)
			values = append(values, value)
		}
	}
	return keys, values
}

func formatToString(v reflect.Value) string {
	// 断言出 time 类型，直接转 unix 时间戳
	if t, ok := v.Interface().(time.Time); ok {
		return fmt.Sprintf("FROM_UNIXTIME(%d)", t.Unix())
	}
	switch v.Kind() {
	case reflect.String:
		return fmt.Sprintf(`'%s'`, v.String())
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return fmt.Sprintf(`%d`, v.Interface())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf(`%f`, v.Interface())
	// 如果是切片类型，遍历元素，递归格式化成"(, , , )"形式
	case reflect.Slice:
		var values []string
		for i := 0; i < v.Len(); i++ {
			values = append(values, formatToString(v.Index(i)))
		}
		return fmt.Sprintf(`(%s)`, strings.Join(values, ","))
	// 接口类型剥一层递归
	case reflect.Interface:
		return formatToString(v.Elem())
	}
	return ""
}

func mKV(v reflect.Value) ([]string, []string) {
	var keys, values []string
	mapKeys := v.MapKeys()
	for _, k := range mapKeys {
		value := formatToString(v.MapIndex(k))
		if value != "" {
			values = append(values, value)
			keys = append(keys, k.Interface().(string))
		}
	}
	return keys, values
}
