//+build wireinject

package main

import "github.com/google/wire"


/*
	参数 name 会给 NewPlayer，因为它需要一个 string 参数
	如果 Build 里面的函数都需要 string 参数，那么 name 会作为他们共同的参数。

	构造器需要的参数(依赖)从参数中查找，或通过其他构造器生成。
	决定选择哪个参数或构造器完全根据类型。
	如果参数或构造器生成的对象有类型相同的情况，运行wire工具时会报错。
	如果我们想要定制创建行为，就需要为不同类型创建不同的参数结构：

	error: NewPlayer 的 error 作为 InitMission 的返回值


	直接在当前目录下执行 wire 或者 新建 wire_gen.go //go:generate wire package main 后再执行 go generate

	go generate
	当运行go generate时，它将扫描与当前包相关的源代码文件，找出所有包含"//go:generate"的特殊注释，
	提取并执行该特殊注释后面的命令，命令为可执行程序，形同shell下面执行。
*/
func InitMission(player PlayerParam, monster MonsterParam) (*Mission, error) {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return &Mission{}, nil
}
