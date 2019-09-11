package zodiac

import "github.com/nocai/gocomm/enum"

// 十二生肖
// 鼠（Rat）、牛（Ox）、虎（Tiger）、兔（Hare）、龙（Dragon）、
// 蛇（Snake）、马（Horse）、羊（Sheep）、猴（Monkey）、鸡（Rooster）、狗（Dog）、猪（Boar）
type Zodiac int

func (z Zodiac) Invalid() bool {
	return enum.DefaultInvalid(int(z), int(invalid))
}

func (z Zodiac) String() string {
	return names[z]
}

const (
	Zodiac_Rat Zodiac = iota
	Zodiac_Ox
	Zodiac_Tiger
	Zodiac_Hare
	Zodiac_Dragon
	Zodiac_Snake
	Zodiac_Horse
	Zodiac_Sheep
	Zodiac_Monkey
	Zodiac_Rooster
	Zodiac_Dog
	Zodiac_Boar

	invalid
)

var names = map[Zodiac]string{
	Zodiac_Rat:     "Rat",
	Zodiac_Ox:      "Ox",
	Zodiac_Tiger:   "Tiger",
	Zodiac_Hare:    "Hare",
	Zodiac_Dragon:  "Dragon",
	Zodiac_Snake:   "Snake",
	Zodiac_Horse:   "Horse",
	Zodiac_Sheep:   "Sheep",
	Zodiac_Monkey:  "Monkey",
	Zodiac_Rooster: "Rooster",
	Zodiac_Dog:     "Dog",
	Zodiac_Boar:    "Boar",
}
