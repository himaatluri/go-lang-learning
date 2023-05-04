package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
	var area_code []string
	var first_three []string
	var last_four []string

	for i := 0; i < 3; i++ {
		d := strconv.Itoa(numbers[i])
		area_code = append(area_code, d)
	}

	for i := 3; i < 6; i++ {
		d := strconv.Itoa(numbers[i])
		first_three = append(first_three, d)
	}

	for i := 6; i < 10; i++ {
		d := strconv.Itoa(numbers[i])
		last_four = append(last_four, d)
	}

	ar_c := strings.Join(area_code, "")
	ftt := strings.Join(first_three, "")
	lft := strings.Join(last_four, "")
	str := []string{"(", ar_c, ")", " ", ftt, "-", lft}
	fmt.Println(strings.Join(str, ""))
}
