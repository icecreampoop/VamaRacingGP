package motorcycles

var (
	Menu map[string]Motorcycle
)

func init() {
	Menu = make(map[string]Motorcycle)
	Menu["BMW_S1000RR"] = NewBMW_S1000RR()
	Menu["DefaultMotorcycle"] = NewDefaultMotorcycle()
	Menu["Ducati_Panigale_V2"] = NewDucati_Panigale_V2()
	Menu["Ducati_Panigale_V4R"] = NewDucati_Panigale_V4R()
	Menu["Kawasaki_ZX10RR"] = NewKawasaki_ZX10RR()
	Menu["Kawasaki_ZX6R"] = NewKawasaki_ZX6R()
	Menu["Yamaha_YZF_R1"] = NewYamaha_YZF_R1()
	Menu["Yamaha_YZF_R7()"] = NewYamaha_YZF_R7()
}

type Motorcycle interface {
	String() string
}

func NewDefaultMotorcycle() *DefaultMotorcycle {
	return &DefaultMotorcycle{
		Cost:       5000,
		Power:      30,
		Handling:   30,
		Durability: 30,
	}
}

func NewBMW_S1000RR() *BMW_S1000RR {
	return &BMW_S1000RR{
		Cost:       28500,
		Power:      88,
		Handling:   76,
		Durability: 72,
	}
}

func NewDucati_Panigale_V4R() *Ducati_Panigale_V4R {
	return &Ducati_Panigale_V4R{
		Cost:       30000,
		Power:      93,
		Handling:   70,
		Durability: 81,
	}
}

func NewDucati_Panigale_V2() *Ducati_Panigale_V2 {
	return &Ducati_Panigale_V2{
		Cost:       18000,
		Power:      65,
		Handling:   57,
		Durability: 53,
	}
}

func NewYamaha_YZF_R1() *Yamaha_YZF_R1 {
	return &Yamaha_YZF_R1{
		Cost:       20000,
		Power:      72,
		Handling:   85,
		Durability: 77,
	}
}

func NewYamaha_YZF_R7() *Yamaha_YZF_R7 {
	return &Yamaha_YZF_R7{
		Cost:       9000,
		Power:      50,
		Handling:   53,
		Durability: 65,
	}
}

func NewKawasaki_ZX10RR() *Kawasaki_ZX10RR {
	return &Kawasaki_ZX10RR{
		Cost:       22500,
		Power:      76,
		Handling:   65,
		Durability: 91,
	}
}

func NewKawasaki_ZX6R() *Kawasaki_ZX6R {
	return &Kawasaki_ZX6R{
		Cost:       13000,
		Power:      58,
		Handling:   59,
		Durability: 63,
	}
}
