package my_number

import "fmt"

func ShowString() {

	var nameOne = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"

	fmt.Println(nameFour)
}

func ShowInteger() {

	var ageOne = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint16 = 256

	fmt.Println(numOne, numTwo, numThree)

}

func ShowFloat() {

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 8129312931289312.7
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
