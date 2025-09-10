package motorcycles

import (
	"fmt"
)

type Yamaha_YZF_R1 struct {
	Cost       float32
	Power      float32
	Handling   float32
	Durability float32
}

func (y *Yamaha_YZF_R1) String() string {
	return fmt.Sprintf("Yamaha YZF R1 Specs\nPrice: $%v\nPower: %v\nHandling: %v\nDurability: %v", y.Cost, y.Power, y.Handling, y.Durability)
}

type Yamaha_YZF_R7 struct {
	Cost       float32
	Power      float32
	Handling   float32
	Durability float32
}

func (y *Yamaha_YZF_R7) String() string {
	return fmt.Sprintf("Yamaha YZF R7 Specs\nPrice: $%v\nPower: %v\nHandling: %v\nDurability: %v", y.Cost, y.Power, y.Handling, y.Durability)
}
