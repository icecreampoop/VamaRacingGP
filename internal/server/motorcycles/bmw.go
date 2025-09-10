package motorcycles

import (
	"fmt"
)

type BMW_S1000RR struct {
	Cost       float32
	Power      float32
	Handling   float32
	Durability float32
}

func (b *BMW_S1000RR) String() string {
	return fmt.Sprintf("BMW S1000RR Specs\nPrice: $%v\nPower: %v\nHandling: %v\nDurability: %v", b.Cost, b.Power, b.Handling, b.Durability)
}
