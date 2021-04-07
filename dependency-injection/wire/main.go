package main

import (
	"errors"
	"fmt"
	"time"
)

/*
	[Go 每日一库之 wire](https://zhuanlan.zhihu.com/p/110453784)

	[使用 google/wire 对 Go 项目进行依赖注入](https://segmentfault.com/a/1190000039060354)
*/

type PlayerParam string
type MonsterParam string

type Monster struct {
	Name string
}

func NewMonster(name MonsterParam) *Monster {
	return &Monster{Name: string(name)}
}

type Player struct {
	Name string
}

func NewPlayer(name PlayerParam) (*Player, error) {
	if time.Now().Unix()%2 == 0 {
		return nil, errors.New("player dead")
	}
	return &Player{Name: string(name)}, nil
}

type Mission struct {
	Player  *Player
	Monster *Monster
}

func NewMission(p *Player, m *Monster) *Mission {
	return &Mission{p, m}
}

func (m *Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player.Name, m.Monster.Name)
}

func main() {
	// monster := NewMonster()
	// player := NewPlayer("dj")
	// mission := NewMission(player, monster)
	// mission.Start()

	mission, err := InitMission(PlayerParam("dj"), MonsterParam("lili"))
	if err != nil {
		panic(err)
	}
	mission.Start()
}
