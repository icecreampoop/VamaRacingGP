package main

import (
	"fmt"
	m "main/motorcycles"
)

func main() {
	bikeMenu := make(map[string]m.Motorcycle)
	bikeMenu["BMW_S1000RR"] = m.NewBMW_S1000RR()
	bikeMenu["DefaultMotorcycle"] = m.NewDefaultMotorcycle()
	bikeMenu["Ducati_Panigale_V2"] = m.NewDucati_Panigale_V2()
	bikeMenu["Ducati_Panigale_V4R"] = m.NewDucati_Panigale_V4R()
	bikeMenu["Kawasaki_ZX10RR"] = m.NewKawasaki_ZX10RR()
	bikeMenu["Kawasaki_ZX6R"] = m.NewKawasaki_ZX6R()
	bikeMenu["Yamaha_YZF_R1"] = m.NewYamaha_YZF_R1()
	bikeMenu["Yamaha_YZF_R7()"] = m.NewYamaha_YZF_R7()

	for _, v := range bikeMenu {
		fmt.Println(v)
		fmt.Println("-------------------------")
	}
}
