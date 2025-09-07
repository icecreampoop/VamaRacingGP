package motorcycles

import (
	"fmt"
)

type DefaultMotorcycle struct {
	Cost       float32
	Power      float32
	Handling   float32
	Durability float32
}

func (y *DefaultMotorcycle) String() string {
	return fmt.Sprintf("Default Motorcycle Specs\nPrice: $%v\nPower: %v\nHandling: %v\nDurability: %v", y.Cost, y.Power, y.Handling, y.Durability)
}
