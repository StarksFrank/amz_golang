package main

import "fmt"

type vehicles interface{}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle //嵌入式结构体，也就是继承自 vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle //嵌入式结构体，也就是继承自 vehicle
	Jet bool
}

type boat struct {
	vehicle  //嵌入式结构体，也就是继承自 vehicle，
	Length int
}

func main() {
	prius := car{}
	tacoma := car{}
	bmw528 := car{}
	boeing747 := plane{}
	boeing757 := plane{}
	boeing767 := plane{}
	sanger := boat{}
	nautique := boat{}
	malibu := boat{}
	rides := []vehicles{prius, tacoma, bmw528, boeing747, boeing757, boeing767, sanger, nautique, malibu}

	for key, value := range rides {
		fmt.Println(key, " - ", value)
	}
}
