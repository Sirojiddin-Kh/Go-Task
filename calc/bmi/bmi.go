package bmi

import (
	"fmt"
)

func BodyMassIndex(mass, height float32)  {

	var bmi float32 = mass / (height * height)

	if (bmi => 18.4) and (bmi =< 25) {

		fmt.Println("Normal weight")

	} else if (bmi > 25) and (bmi =< 30) {

		fmt.Println("Overweight")

	} else if (bmi > 30) and (bmi =< 40) {

		fmt.Println("Obesity")

	} else if bmi > 40 {

		fmt.Println("Morbid	Obesity")

	} else {

		fmt.Println("Mass and height to big")
	}
}