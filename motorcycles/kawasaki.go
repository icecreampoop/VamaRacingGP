package motorcycles

import (
	"fmt"
)

type Kawasaki_ZX10RR struct {
	Cost       float32
	Power      float32
	Handling   float32
	Durability float32
}

func (b *Kawasaki_ZX10RR) String() string {
	return fmt.Sprintf("Kawasaki ZX10RR Specs\nPrice: $%v\nPower: %v\nHandling: %v\nDurability: %v", b.Cost, b.Power, b.Handling, b.Durability)
}

type Kawasaki_ZX6R struct {
	Cost       float32
	Power      float32
	Handling   float32
	Durability float32
}

func (b *Kawasaki_ZX6R) String() string {
	return fmt.Sprintf("Kawasaki ZX6R Specs\nPrice: $%v\nPower: %v\nHandling: %v\nDurability: %v", b.Cost, b.Power, b.Handling, b.Durability)
}
