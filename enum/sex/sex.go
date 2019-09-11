package sex

import "github.com/nocai/gocomm/enum"

// 性别
type Sex int

func (s Sex) Invalid() bool {
	return enum.DefaultInvalid(int(s), int(invalid))
}

func (s Sex) String() string {
	return names[s]
}

const (
	Sex_Male Sex = iota
	Sex_Female

	invalid
)

var names = map[Sex]string{
	Sex_Male:   "Male",
	Sex_Female: "Female",
}
