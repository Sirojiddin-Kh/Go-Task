package bmi

import (
	"fmt"
)

func BodyMassIndex(mass, height float32) {

	var BMI float32 = mass / (height * height)

	if BMI < 18.5 {
		fmt.Println("You are underweight")
	} else if BMI >= 18.5 && BMI < 25.0 {
		fmt.Println("Your weight is normal")
	} else if BMI >= 25.0 && BMI < 30.0 {
		fmt.Println("You're overweight")
	} else {
		fmt.Println("You're obese")
	}
} 