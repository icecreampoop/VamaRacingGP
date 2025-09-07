package motorcycles

import (
	"fmt"
)

type Ducati_Panigale_V4R struct {
	Cost       float32
	Power      float32
	Handling   float32
	Durability float32
}

func (b *Ducati_Panigale_V4R) String() string {
	return fmt.Sprintf("Ducati Panigale V4R Specs\nPrice: $%v\nPower: %v\nHandling: %v\nDurability: %v", b.Cost, b.Power, b.Handling, b.Durability)
}

type Ducati_Panigale_V2 struct {
	Cost       float32
	Power      float32
	Handling   float32
	Durability float32
}

func (b *Ducati_Panigale_V2) String() string {
	return fmt.Sprintf("Ducati Panigale V2 Specs\nPrice: $%v\nPower: %v\nHandling: %v\nDurability: %v", b.Cost, b.Power, b.Handling, b.Durability)
}
