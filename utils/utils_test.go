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
	fmt.Println(Base36Encode(1549631778698827714))
}
