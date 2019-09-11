package bloodtype

import "github.com/nocai/gocomm/enum"

// 血型:A,B,O,AB
// 说明:
// 血型分型标准有很多，但目前我国在临床实际用到的只有两种
// 一是ABO分型，另一种是RH分型
// ABO分型将人的血型分为A、B、O、AB四个血型
// RH分型将血型分为RH阳性和RH阴性
// 由于汉族人基本上都是RH阳性，RH阴性的只占不到0.5%，所以过去通常不考虑，只做ABO血型测定
type BloodType int

func (bt BloodType) Invalid() bool {
	return enum.DefaultInvalid(int(bt), int(invalid))
}

func (bt BloodType) String() string {
	return names[bt]
}

const (
	// A
	BloodType_A BloodType = iota
	// B
	BloodType_B
	// O
	BloodType_O
	// AB
	BloodType_AB

	invalid
)

var names = map[BloodType]string{
	BloodType_A:  "A",
	BloodType_B:  "B",
	BloodType_AB: "AB",
	BloodType_O:  "O",
}
