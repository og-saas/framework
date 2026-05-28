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
	fmt.Println(Base36Encode(3830248033757650944))
}

func TestBase36Decode(t *testing.T) {
	//fmt.Println(Base36Decode("t3m5lz9lp8g0"))
	fmt.Println(Base36Decode("bs2punt1f8y2"))
}
