package db

import (
	"strconv"
	"testGoScript/0INACTION/routerRedisRestful/common"
	"testGoScript/0INACTION/routerRedisRestful/models"

	"github.com/garyburd/redigo/redis"
)

var Conn redis.Conn

func init() {
	// 第三个参数可以指定连接选项
	conn, err := redis.Dial("tcp", ":6379", redis.DialDatabase(1))
	common.CheckErr(err)
	Conn = conn
}

func Insert(u *models.User) error {
	UserMaxID, err := redis.Int(Conn.Do("GET", "UserMaxID")) // redis 可以自动转换为相应的类型
	if err != redis.ErrNil && err != nil {
		return err
	}

	u.Id = UserMaxID + 1
	i := strconv.Itoa(u.Id)

	_, err = Conn.Do("HMSET", redis.Args{}.Add(i).AddFlat(u)...) // HMSET 可以将 struct 直接映射成 map, 标签注明redis
	if err != nil {
		return err
	}
	_, err = Conn.Do("INCR", "UserMaxID")
	if err != nil {
		return err
	}

	// v, err := redis.Values(Conn.Do("HGETALL", strconv.Itoa(u.Id))) // 获取所有值
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(reflect.TypeOf(v), v)

	// uu := new(models.User)        // 如果 var uu *models.User, 那么 uu 将是 nil
	// err = redis.ScanStruct(v, uu) // 将上面的值,转换为 go 的 struct
	// fmt.Println("err :", err)
	// fmt.Println(uu)
	// fmt.Printf("uu:%+v", &uu)

	return err
}

func FindOne() {

}

func FindAdd() {

}

func Delete() {

}

func Update() {

}
