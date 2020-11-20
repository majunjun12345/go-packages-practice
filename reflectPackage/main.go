package main

import (
	"fmt"
	"reflect"
)

/*
	reflect 有两个数据类型: Type(数据类型) 和 Value(值)

	反射是一种检查存储在 接口变量 中的 <类型 值> 的机制
	通过 TypeOf 可以访问到 Type, 通过 ValueOf 可以访问到 Value, 从而可以进一步得到这个接口的结构类型和对其值进行操作;
*/

type User struct {
	Name    string `gogo:"name" json:"name" id:"100"`
	Age     int    `gogo:"age"`
	Married bool   `married:"married"`
}

func (u *User) Hello() string {
	return fmt.Sprintf("hello,my name is:%s\n", u.Name)
}

func main() {

	// u := &User{"mengliam", 21, false}
	// TStructType(u)

	// TStructValue(u)

	TStructCall()
}

// TStructType 通过反射获取类型信息
func TStructType(s interface{}) {
	t := reflect.TypeOf(s) // main.User
	k := t.Kind()          // 如果传入的是指针: ptr，如果传入的是结构体: struct
	fmt.Println("typeof valueof kind", t, k)

	// 获取结构体字段
	switch t.Kind() {
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			/*
				Name string mengliam
				Age int 21
				Married bool false
			*/
			fmt.Println(t.Field(i).Name, t.Field(i).Type)

			// reflect.StructField Name name  0 [0] false
			fmt.Println("reflect.StructField", t.Field(i).Name, t.Field(i).Tag.Get("gogo"), t.Field(i).PkgPath, t.Field(i).Offset, t.Field(i).Index, t.Field(i).Anonymous)

			switch t.Field(i).Type.Kind() {
			case reflect.String:
				fmt.Println(t.Field(i).Name)
			case reflect.Int:
				fmt.Println(t.Field(i).Name)
			case reflect.Bool:
				fmt.Println(t.Field(i).Name)
			default:
				fmt.Println("Unsupported type", t.Field(i).Name)
			}
		}

		// 获取方法， 方法必须是可导出的
		for i := 0; i < t.NumMethod(); i++ {
			m := t.Method(i)
			fmt.Println("=======", m.Name)
		}

		if m, ok := t.MethodByName("Hello"); ok {
			/*
				在Go的指针知识中，有一条规则：一个指针类型拥有它以及它的基底类型为接收者类型的所有方法，而它的基底类型却只能拥有以它本身为接收者类型的方法。
			*/
			in := []reflect.Value{reflect.ValueOf(&User{Name: "mamama"})}
			rs := m.Func.Call(in)
			fmt.Println("result:", rs[0].Interface().(string))
		}
	case reflect.Ptr: // 如果是指针，则要剥离指针
		fmt.Println("is ptr")

		// 方法
		{
			// 获取方法， 方法必须是可导出的
			for i := 0; i < t.NumMethod(); i++ {
				m := t.Method(i)
				fmt.Println("=======", m.Name)
			}

			if m, ok := t.MethodByName("Hello"); ok {
				/*
					User类型是*User的基底类型
					在Go的指针知识中，有一条规则：一个指针类型拥有它以及它的基底类型为接收者类型的所有方法，而它的基底类型却只能拥有以它本身为接收者类型的方法。
					如果被剥离指针, 这里将 panic
				*/
				in := []reflect.Value{reflect.ValueOf(&User{Name: "mamama"})}
				rs := m.Func.Call(in)
				fmt.Println("result:", rs[0].Interface().(string))
			}
		}

		// 剥离指针
		for t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		// 具体字段
		for i := 0; i < t.NumField(); i++ {

			// Name string
			fmt.Println("name, type", t.Field(i).Name, t.Field(i).Type.String())

			// reflect.StructField Name name  0 [0] false
			fmt.Println("reflect.StructField", t.Field(i).Name, t.Field(i).Tag.Get("gogo"), t.Field(i).PkgPath, t.Field(i).Offset, t.Field(i).Index, t.Field(i).Anonymous)

			switch t.Field(i).Type.Kind() {
			case reflect.String:
				fmt.Println(t.Field(i).Name)
			case reflect.Int:
				fmt.Println(t.Field(i).Name)
			case reflect.Bool:
				fmt.Println(t.Field(i).Name)
			default:
				fmt.Println("Unsupported type", t.Field(i).Name)
			}
		}

		// 通过字段名, 找到字段类型信息
		if nameType, ok := t.FieldByName("Name"); ok {
			// 从tag中取出需要的tag
			fmt.Println(nameType.Tag.Get("json"), nameType.Tag.Get("id"))
		}
	}
}

// TStructValue 通过反射获取值信息
func TStructValue(s interface{}) {
	// 判断反射值得空和有效性
	{

		var a *int
		fmt.Println("var a *int:", reflect.ValueOf(a).IsNil()) //*int的空指针
		fmt.Println("nil:", reflect.ValueOf(nil).IsValid())    //nil值

		//实例化一个结构体
		s := struct{}{}
		fmt.Println("不存在的结构体成员:", reflect.ValueOf(s).FieldByName("").IsValid()) //尝试从结构体中查找一个不存在的字段
		fmt.Println("不存在的方法:", reflect.ValueOf(s).MethodByName("").IsValid())   //尝试从结构体中查找一个不存在的方法

		//实例化一个map
		m := map[int]int{}
		fmt.Println("不存在的键:", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid()) //尝试从map中查找一个不存在的键
	}

	v := reflect.ValueOf(s) // &{mengliam 21 false}

	i := v.Interface() // 转换成 interface 类型，可以用来类型断言
	if vv, ok := i.(*User); ok {
		fmt.Println("======", vv)
	}

	// 通过反射访问结构体成员的值
	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.String:
				fmt.Println(v.Field(i).String())
			case reflect.Int:
				fmt.Println(v.Field(i).Int())
			case reflect.Bool:
				fmt.Println(v.Field(i).Bool())
			default:
				fmt.Println("Unsupported type", v.Field(i).Interface())
			}
		}
	case reflect.Ptr: // 如果是指针，则要剥离指针
		fmt.Println("is ptr")
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		rv := v.FieldByName("Name")
		fmt.Println("=======", rv.Type().Name())

		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.String:
				fmt.Println(v.Field(i).String())
			case reflect.Int:
				fmt.Println(v.Field(i).Int())
			case reflect.Bool:
				fmt.Println(v.Field(i).Bool())
			default:
				fmt.Println("Unsupported type", v.Field(i).Interface())
			}
		}
	}

	// Int 和 String 可以帮助我们分别取出 reflect.Value 作为 int64 和 string
	{
		a := 56
		x := reflect.ValueOf(a).Int()
		fmt.Printf("type:%T value:%v\n", x, x) // type:int64 value:56
		b := "Naveen"
		y := reflect.ValueOf(b).String()
		fmt.Printf("type:%T value:%v\n", y, y) // type:string value:Naveen
	}

	// 获取方法， 方法必须是可导出的
	{
		m := v.MethodByName("Hello")
		/*
			这里会 panic，因为上面将传入的 v 被 剥离了指针
			User类型是*User的基底类型
			在Go的指针知识中，有一条规则：一个指针类型拥有它以及它的基底类型为接收者类型的所有方法，而它的基底类型却只能拥有以它本身为接收者类型的方法。
		*/
		rs := m.Call(nil)
		fmt.Println("result:", rs[0].Interface().(string))
	}
}

// TStructModifyValue 通过反射修改值
func TStructModifyValue() {

	var a int = 1024
	rValue := reflect.ValueOf(&a) // 获取变量a的反射值对象
	rValue = rValue.Elem()        // 取出a地址的元素(a的值)
	rValue.SetInt(1)              // 尝试将a修改为1
	fmt.Println(rValue.Int())

	type dog struct {
		legCount int
	}

	valueOfDog := reflect.ValueOf(&dog{}) // 获取dog实例的反射值对象
	valueOfDog = valueOfDog.Elem()
	vLegCount := valueOfDog.FieldByName("legCount") // 获取legCount字段的值
	vLegCount.SetInt(4)                             //尝试设置legCount的值(这里会发生崩溃), 原因是 legCount 不能被导出
}

// TStructCall 通过反射调用结构体的方法
func TStructCall() {
	//通过反射的方式调用结构体类型的方法
	var setNameStr string = "SetName"
	var addAgeStr string = "AddAge"
	user := UserInfo{
		Id:   1,
		Name: "env107",
		Age:  18,
	}

	//1.获取到结构体类型变量的反射类型
	refUser := reflect.ValueOf(&user) // 需要传入指针，后面再解析
	fmt.Println(refUser)

	//2.获取确切的方法名
	//带参数调用方式
	setNameMethod := refUser.MethodByName(setNameStr)
	args := []reflect.Value{reflect.ValueOf("Mike")} //构造一个类型为reflect.Value的切片
	setNameMethod.Call(args)                         //返回Value类型

	//不带参数调用方式
	addAgeMethod := refUser.MethodByName(addAgeStr)
	addAgeMethod.Call(make([]reflect.Value, 0))

	fmt.Println("User.Name = ", user.Name)
	fmt.Println("User.Age = ", user.Age)

	{
		var myMath = MyMath{Pi: 3.14159}
		rValue := reflect.ValueOf(myMath)                                      // 获取myMath的值对象
		paramList := []reflect.Value{reflect.ValueOf(30), reflect.ValueOf(20)} // 构造函数参数，传入两个整形值

		//调用结构体的第一个方法Method(0)
		//注意:在反射值对象中方法索引的顺序并不是结构体方法定义的先后顺序
		//而是根据方法的ASCII码值来从小到大排序，所以Dec排在第一个，也就是Method(0)
		result := rValue.Method(0).Call(paramList)
		fmt.Println(result[0].Int())
	}
}

type UserInfo struct {
	Id   int
	Name string
	Age  int
}

//ToString方法
func (u UserInfo) String() string {
	return "User[ Id " + string(u.Id) + "]"
}

//设置Name方法
func (u *UserInfo) SetName(name string) string {
	oldName := u.Name
	u.Name = name
	return oldName
}

//年龄数+1
func (u *UserInfo) AddAge() bool {
	u.Age++
	return true
}

//测试方法
func (u UserInfo) TestUser() {
	fmt.Println("我只是输出某些内容而已....")
}

type MyMath struct {
	Pi float64
}

//普通函数
func (myMath MyMath) Sum(a, b int) int {
	return a + b
}

func (myMath MyMath) Dec(a, b int) int {
	return a - b
}
