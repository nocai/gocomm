package constellation

import "github.com/nocai/gocomm/enum"

// 十二星座:
// Aquarius 水瓶座（1月21日～2月19日）
// Pisces 双鱼座（2月20日～3月20日）
// Aries 白羊座（3月21日～4月20日）
// Taurus 金牛座（4月21～5月21日）
// Gemini 双子座（5月22日～6月21日）
// Cancer 巨蟹座（6月22日～7月22日）
// Leo 狮子座（7月23日～8月23日）
// Virgo 处女座（8月24日～9月23日）
// Libra 天秤座（9月24日～10月23日）
// Scorpio 天蝎座（10月24日～11月22日）
// Sagittarius 射手座（11月23日～12月21日）
// Capricorn 摩羯座（12月22日～1月20日）
type Constellation int

func (c Constellation) Invalid() bool {
	return enum.DefaultInvalid(int(c), int(invalid))
}

func (c Constellation) String() string {
	return names[c]
}

const (
	// Aquarius 水瓶座（1月21日～2月19日）
	Constellation_Aquarius Constellation = iota
	// Pisces 双鱼座（2月20日～3月20日）
	Constellation_Pisces
	// Aries 白羊座（3月21日～4月20日）
	Constellation_Aries
	// Taurus 金牛座（4月21～5月21日）
	Constellation_Taurus
	// Gemini 双子座（5月22日～6月21日）
	Constellation_Gemini
	// Cancer 巨蟹座（6月22日～7月22日）
	Constellation_Cancer
	// Leo 狮子座（7月23日～8月23日）
	Constellation_Leo
	// Virgo 处女座（8月24日～9月23日）
	Constellation_Virgo
	// Libra 天秤座（9月24日～10月23日）
	Constellation_Libra
	// Scorpio 天蝎座（10月24日～11月22日）
	Constellation_Scorpio
	// Sagittarius 射手座（11月23日～12月21日）
	Constellation_Sagittarius
	// Capricorn 摩羯座（12月22日～1月20日）
	Constellation_Capricorn // 11

	invalid
)

var names = map[Constellation]string{
	Constellation_Aquarius:    "Aquarius",
	Constellation_Pisces:      "Pisces",
	Constellation_Aries:       "Aries",
	Constellation_Taurus:      "Taurus",
	Constellation_Gemini:      "Gemini",
	Constellation_Cancer:      "Cancer",
	Constellation_Leo:         "Leo",
	Constellation_Virgo:       "Virgo",
	Constellation_Libra:       "Libra",
	Constellation_Scorpio:     "Scorpio",
	Constellation_Sagittarius: "Sagittarius",
	Constellation_Capricorn:   "Capricorn",
}
