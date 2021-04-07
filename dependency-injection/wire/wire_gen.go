// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

// Injectors from wire.go:

/*
	参数 name 会给 NewPlayer，因为它需要一个 string 参数
	如果 Build 里面的函数都需要 string 参数，那么 name 会作为他们共同的参数。

	构造器需要的参数(依赖)从参数中查找，或通过其他构造器生成。
	决定选择哪个参数或构造器完全根据类型。
	如果参数或构造器生成的对象有类型相同的情况，运行wire工具时会报错。
	如果我们想要定制创建行为，就需要为不同类型创建不同的参数结构：

	error: NewPlayer 的 error 作为 InitMission 的返回值
*/
func InitMission(player PlayerParam, monster MonsterParam) (*Mission, error) {
	mainPlayer, err := NewPlayer(player)
	if err != nil {
		return nil, err
	}
	mainMonster := NewMonster(monster)
	mission := NewMission(mainPlayer, mainMonster)
	return mission, nil
}
