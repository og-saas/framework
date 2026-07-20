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

func TestBase36Decode(t *testing.T) {
	//fmt.Println(Base36Decode("t3m5lz9lp8g0"))
	//fmt.Println(Base36Decode("bs2punt1f8y2"))
	//fmt.Println(Base36Decode("oy74vf2pl3p"))
	fmt.Println(Base36Decode("pmfo4tlxp9h"))
}
