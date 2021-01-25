package main

import (
	"database/sql"
	"fmt"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*
	[深入剖析在事务中并发执行查询异常](https://medium.com/impopper-engineering/go-concurrency-query-in-mysql-transactions-f6018c7b16b2)
	执行 query 语句获取 rows，数据还在 buffer 里面没有读出来，用 for rows.Next(){rows1.Scan()} 可以将数据读出来（next取出buffer内数据，scan绑定数据），清空 buffer，close 也可以；
	一个事物维持一个连接，如果在读出来之前执行另一个 query 操作，会导致 busy buffer

	每个 sql 语句都对应一个 buffer，执行结果存储在里面，一个事物离只会有一个 连接实例，
	也就是在一个事物里如果有多个 query 语句，每执行一个 query 语句后必须清空 buffer

	query 和 exec 语句的不同：
		在一个事物中可以并发的执行更新或插入操作，因为更新或插入调用的是 exec，db 返回的 result 直接从 buffer 里面取出来了；
*/

/*
	go 和 mysql 中的 时间类型和bool 类型可以相互转换，设置 parseTime=true

	DB.Exec(): 执行不返回 row 的命令, 比如: delete update insert 等; 可以单独执行, 也可以 prepare 预处理后再执行
	DB.Query() DB.QueryRow: 用于查询一行或多行, prepare 后的 stmt 也可以执行 query

	每个执行操作 (query queryRow等) 在并发时都是独立的连接,但是包裹在事物里面的是一个连接,里面所有的操作都是串行执行;

*/

/*
	程序连接数据库会有连接泄漏的情况，需要及时释放连接

	Go sql包中的Query和QueryRow两个方法的连接不会自动释放连接，只有在遍历完结果或者调用close方法才会关闭连接

	Go sql中的Ping和Exec方法在调用结束以后就会自动释放连接, scan 之后也会释放连接

	忽略了函数的某个返回值不代表这个值就不存在了，如果该返回值需要close才会释放资源，直接忽略了就会导致资源的泄漏。

	有close方法的变量，在使用后要及时调用该方法，释放资源
*/

/*
	DROP TABLE IF EXISTS `user`;
CREATE TABLE `user_info` (
	`id` INT(10) NOT NULL AUTO_INCREMENT COMMENT '自增id',
	`username` VARCHAR(64) NULL DEFAULT "majun" COMMENT '用户名',
	`created` DATETIME NULL DEFAULT "1990-09-28 06:15:15" COMMENT '创建时间',
	`married` BOOLEAN NOT NULL DEFAULT TRUE COMMENT 'haha',
	PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT=' 用户管理';
*/

type user_info struct {
	Id          int       `db:"id"`
	Username    string    `db:"username"`
	CreatedTime time.Time `db:"created"`
	Married     bool      `db:"married"`
}

var DB *sql.DB

func init() {
	var err error
	if DB, err = sql.Open("mysql", "root:Mj900928@tcp(127.0.0.1:3306)/orm_db?charset=utf8&parseTime=true"); err != nil {
		panic(err)
	}
	if err = DB.Ping(); err != nil { // open 并不会建立一个连接,只有当你使用的时候才会建立连接,所以这里需要提前 ping 一下,确保连接正常
		panic(err)
	}

	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(10)                   //设置最大连接数
	DB.SetMaxIdleConns(5)                    //设置闲置连接数
}

func main() {
	// insert()

	// query()

	tx()

	// queryNoRows()
}

// --------------------------insert exec 主要是执行插入和更新操作
// 两种插入方式
func insert() {
	stmt, err := DB.Prepare("INSERT user_info set username=?, created=?, married=?") // 最好先使用预处理(这种方式)
	CheckErr(err)
	_, err = stmt.Exec("memglima", GetTime(), false)
	CheckErr(err)

	DB.Exec("INSERT user_info set username=?, created=?, married=?", "memglima", GetTime(), false)
}

// --------------------------query 主要是执行 查询 操作
func query() {
	/*
		返回的 row 必须调用 Close 将连接返还连接池, 会导致超出最大连接数限制，并导致死锁或者高延迟；
		[为何sql.Rows使用结束后一定要Close](https://blog.csdn.net/xz_studying/article/details/109588435)

		**当你使用它执行数据库 **查询** 等任务时，一个连接就会被标记成 in-use；任务执行完成后该连接又会被标记成 idle。**

		query 语句会建立一次连接，使用 next scan 后清空的是 buffer 并不是关闭连接   https://www.jianshu.com/p/06f26f879d61
		rows.close 才是关闭的连接
		* 也可以替换为具体的查询字段
	*/
	rows, err := DB.Query("select id , username, created, married from user_info") // QueryRow 查询一行
	defer rows.Close()                                                             // 这里才是关闭 query 的连接

	if err != nil {
		panic(err)
	}
	var t string
	var b bool
	for rows.Next() {
		user := user_info{}
		// 参数绑定需要进行错误处理
		if err := rows.Scan(&user.Id, &user.Username, &t, &b); err != nil { // 和 query 一一对应,最好不要用 *
			fmt.Println(err)
			continue // 如果有错误，则忽略这条记录
		}
		fmt.Println(user)
	}
	if err := rows.Err(); err != nil { // 再次进行判断，遍历过程中是否有错误产生
		fmt.Println(err)
	}

	// 查询单条记录
	var username string
	row := DB.QueryRow("select username from user_info where id=?", 2)
	err = row.Scan(&username)
	switch err {
	case sql.ErrNoRows: // 查询单条数据可能出现没有该条记录的错误，需要单独处理
		fmt.Println("no such data")
	case nil:
		if _, file, line, ok := runtime.Caller(3); ok { // // 使用该方式可以打印出运行时的错误信息, 该种错误是编译时无法确定的
			fmt.Println(err, file, line)
		}
	}

	// 关于 null
	// 所有查询出来的字段都不允许有NULL, 避免该方式最好的办法就是建表字段时, 不要设置类似DEFAULT NULL属性
	// 下面这种语句可以避免 null，不过 having 一般和 group by 搭配使用，相当于 where 后面的条件语句；
	var id int32
	err = DB.QueryRow(`
        SELECT
            SUM(id) id
        FROM user_info
        WHERE id = ?
        HAVING id <> NULL
    `, 10).Scan(&id)
	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		fmt.Println(err)
	}
	fmt.Println(id)
}

func queryNoRows() {
	// 查询 row 为空，scan 返回对象的零值，err 为 sql.ErrNoRows
	user := user_info{}
	row := DB.QueryRow("select id , username, created, married from user_info where id=?", 1)
	err := row.Scan(&user.Id, &user.Username, &user.CreatedTime, &user.Married)
	fmt.Println(user)
	fmt.Printf("err0:%v\n", err)
	fmt.Println(err == sql.ErrNoRows)

	// 为空，不会报错，
	stmt, err := DB.Prepare("select id , username, created, married from user_info where id=?")
	if err != nil {
		fmt.Printf("err1:%v\n", err)
	}
	res, err := stmt.Exec(4)
	if err != nil {
		fmt.Printf("err2:%v\n", err)
	}

	fmt.Println(res.RowsAffected()) // 0
}

// --------------------------事物 tx
func tx() { //这里没有清空 buffer，但是也没有报错，就有点不懂了
	tx, _ := DB.Begin()
	_, err1 := DB.Query("select id , username, created, married from user_info")

	stmt, err3 := DB.Prepare("INSERT user_info set username=?, created=?, married=?")
	_, err4 := stmt.Exec("memglima", GetTime(), false)

	if err1 == nil && err3 == nil && err4 == nil {

		//只有两条更新同时成功，Begin与Commit配对，才会提交
		tx.Commit()
		fmt.Println("Success")
	} else {

		//否则回滚到Begin，提高了安全性
		tx.Rollback()
		fmt.Println(err1)
		fmt.Println("Fail")
	}

	// ret, _ := tx.Exec("update user set username = hahaha where id = ?", 1)
	// ret1, _ := tx.Exec("update user set username = masanqi where id = ?", 2)

	// upd_num1, _ := ret.RowsAffected()
	// upd_num2, _ := ret1.RowsAffected()

	// if upd_num1 > 0 && upd_num2 > 0 {

	// 	//只有两条更新同时成功，Begin与Commit配对，才会提交
	// 	tx.Commit()
	// 	fmt.Println("Success")
	// } else {

	// 	//否则回滚到Begin，提高了安全性
	// 	tx.Rollback()
	// 	fmt.Println("Fail")
	// }
}

func GetTime() string {
	const shortForm = "2006-01-02 15:04:05"
	t := time.Now()
	return t.Format(shortForm)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
