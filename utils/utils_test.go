package utils

import (
	"fmt"
	"testing"
)

func TestBase62Encode(t *testing.T) {
	fmt.Println(Base62Encode(3830248033757650944))
}

func TestBase36Encode(t *testing.T) {
	//fmt.Println(Base36Encode(3830248033757650944))
	//fmt.Println(Base36Encode(1550487228994551818))
	//fmt.Println(Base36Encode(1550640147949682702))
	//fmt.Println(Base36Encode(1549631778698827714))
	//fmt.Println(Base36Encode(85419477293924872))
	//fmt.Println(Base36Encode(3830248033757650944))
	//fmt.Println(Base36Encode(91212108980552970))
	//fmt.Println(Base36Encode(93682487934322469))
	//fmt.Println(Base36Encode(91212108980552970))
	//fmt.Println(Base36Encode(91220970706044181))
	fmt.Println(Base36Encode(91339144667596113))
}

func TestRealPhone(t *testing.T) {
	ret, err := GetRealPhone("6309307853678", 63)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ret)

	// 纯本地号码(不带国家码),靠 countryCode 作为 region 提示
	ret, err = GetRealPhone("9307853671", 63)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ret)

	// 无效国家码
	if _, err = GetRealPhone("09307853678", 999); err == nil {
		t.Error("expected error for invalid country code")
	}
}

func TestBase36Decode(t *testing.T) {
	//fmt.Println(Base36Decode("t3m5lz9lp8g0"))
	//fmt.Println(Base36Decode("bs2punt1f8y2"))
	//fmt.Println(Base36Decode("oy74vf2pl3p"))
	fmt.Println(Base36Decode("pmfo4tlxp9h"))
}
