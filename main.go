package main

import (
	"fmt"
	bike "main/motorcycles"
	t "main/interfaces"
)

func main() {
	test := &bike.Yamaha_YZF_R1{
		Cost: 10000,
	}

	fmt.Println(test)
	var test2 t.Motorcycle = &bike.Yamaha_YZF_R1{Cost:33,}

	fmt.Println(test2)

}