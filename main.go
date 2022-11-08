package main

import (
	"fmt"
)

// Метод
func (c Car) PrintArray() {
	for _, value := range CarArray {
		fmt.Println(value)
	}
}
func PrintWhiteCars() {
	for key, value := range CarArray {
		if carColor[value] == "White" {
			fmt.Printf("White Cars : %v \n", CarArray[key])
		}
	}
}

// Создаем структуру Car
type Car struct {
	brandName string
	carName   string
	price     int
	year      int
}

// Объявляем массив с нашимим машинами
var CarArray = []Car{
	{
		brandName: "Toyota",
		carName:   "Mark2",
		price:     1000000,
		year:      1996,
	},
	{
		brandName: "Nissan",
		carName:   "350z",
		price:     200000,
		year:      2002,
	},
	{
		brandName: "Lada",
		carName:   "VAZ 2114",
		price:     150000,
		year:      2010,
	},
	{
		brandName: "Toyota",
		carName:   "Camry",
		price:     330000,
		year:      2014,
	},
}
var carColor = map[Car]string{
	CarArray[0]: "White",
	CarArray[1]: "Black",
	CarArray[2]: "Blue",
}

func main() {
	Car.PrintArray(Car{})
	PrintWhiteCars()

}
