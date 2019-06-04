package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// https://www.jianshu.com/p/bc8120bec94e

/*
	prepare 的设计初衷是多次执行:
		可以实现自定义参数
		比手动拼接 sql 语句更高效
		可以防止 SQL 注入攻击

	事务:
	tx, err := db.Begin()

	tx.Prepare
	stmt.Exec
		defer stmt.Close()  中间这三个步骤有点小问题，因为 commit 或 rollback 后，连接就释放了，不需要再进行 close，因此在事物中尽量避免 prepare

		tx.Commit()
	defer tx.RollBack()

	一旦创建了 tx 对象，事物的处理都依赖于 tx 对象，该对象会从连接池中取出一个空闲的连接，
	接下来的所有操作都基于该连接，直到 commit 或 rollback 后才会将连接释放至连接池；
	事物过程中只有一个连接；
	事物内的操作都是顺序执行；
	普通的 db 执行一个操作就是一个连接；


	tx 的并发问题：
		对于sql.Tx对象，因为事务过程只有一个连接，事务内的操作都是顺序执行的，在开始下一个数据库交互之前，必须先完成上一个数据库交互。例如下面的例子：
		rows, _ := db.Query("SELECT id FROM user")
		for rows.Next() {
			var mid, did int
			rows.Scan(&mid)
			db.QueryRow("SELECT id FROM detail_user WHERE master = ?", mid).Scan(&did)

		}

		调用了Query方法之后，在Next方法中取结果的时候，rows是维护了一个连接，再次调用QueryRow的时候，db会再从连接池取出一个新的连接。rows和db的连接两者可以并存，并且相互不影响。
		可是，这样逻辑在事务处理中将会失效：
		rows, _ := tx.Query("SELECT id FROM user")
		for rows.Next() {
		var mid, did int
		rows.Scan(&mid)
		tx.QueryRow("SELECT id FROM detail_user WHERE master = ?", mid).Scan(&did)
		}

		tx执行了Query方法后，连接转移到rows上，在Next方法中，tx.QueryRow将尝试获取该连接进行数据库操作。因为还没有调用rows.Close，因此底层的连接属于busy状态，tx是无法再进行查询的。上面的例子看起来有点傻，毕竟涉及这样的操作，使用query的join语句就能规避这个问题。例子只是为了说明tx的使用问题。
*/

type BCCode struct {
	b_code     string
	c_code     string
	code_type  int
	is_integer int
}

func main() {
	db, err := sql.Open("sqlite3", "file::memory:?mode=memory&cache=shared&loc=auto")
	defer db.Close()
	CheckErr1(err)

	fmt.Println("SQLite start")

	//创建表//delete from BC;，SQLite字段类型比较少，bool型可以用INTEGER，字符串用TEXT
	sqlStmt := `create table BC (
		b_code text not null primary key, 
		c_code text not null, 
		code_type INTEGER, 
		is_new INTEGER
		);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Println("create table error->%q: %s\n", err, sqlStmt)
		return
	}

	//创建索引，有索引和没索引性能差别巨大，根本就不是一个量级，有兴趣的可以去掉试试
	_, err = db.Exec("CREATE INDEX inx_c_code ON BC(c_code);")
	if err != nil {
		fmt.Println("create index error->%q: %s\n", err, sqlStmt)
		return
	}

	// 插入 10 万条数据
	start := time.Now().Unix()
	tx, err := db.Begin()
	CheckErr1(err)
	stmt, err := tx.Prepare("insert into BC(b_code, c_code, code_type, is_new) values(?,?,?,?)")
	CheckErr1(err)
	defer stmt.Close()
	var n int = 1000 * 1000
	for i := 0; i < n; i++ {
		_, err := stmt.Exec(fmt.Sprintf("B%024d", i), fmt.Sprintf("B%024d", i), 0, 1)
		CheckErr1(err)
	}
	end := time.Now().Unix()

	/*
		用来实现提交或回滚，这里的 recover 只针对事物里面的 err(示例代码)
		为了实现 err 不被覆盖，还可以进行检测，当 err 不为 nil 的时候，return，然后调用 defer 就行，也不需要 recover 了；
	*/
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			tx.Rollback()
		} else {
			tx.Commit()
			fmt.Println("ok")
		}
	}()

	// 随机检索 10 万次
	var count int = 0
	stmt, err = db.Prepare("select b_code, c_code, code_type, is_new from BC where c_code=?")
	defer stmt.Close()
	CheckErr1(err)
	bc := new(BCCode)
	for i := 0; i < n; i++ {
		err := stmt.QueryRow(fmt.Sprintf("B%024d", i)).Scan(&bc.b_code, &bc.c_code, &bc.code_type, &bc.is_integer)
		CheckErr1(err)
		count++
	}

	queryEnd := time.Now().Unix()
	fmt.Println("insert into 10万次,cost:", float64(end)-float64(start))
	fmt.Println("query 10万次,cost:", float64(queryEnd)-float64(end))
	fmt.Println(count)
}

func CheckErr1(err error) {
	if err != nil {
		panic(err)
	}
}
