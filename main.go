package main

import "fmt"

func intit() {
	fmt.Println("Init ... ")
}

func main() {

	fmt.Println("DAY1 Start ... ")
	fmt.Println(" Hello World! Golang ")
	fmt.Println("-----------------------")



	fmt.Println("DAY2 Start ... ")

	var variableTypeBoolean = true
	var variableTypeInteger = 6
	variableTypeFloat := 1.25
	variableTypeString := "Hello DAY#2"
	variableTypeFunction := func ()  {
		fmt.Println("In VariableTypeFunction")
	}

	fmt.Println(variableTypeBoolean)
	fmt.Println(variableTypeInteger)
	fmt.Println(variableTypeFloat)
	fmt.Println(variableTypeString)
	variableTypeFunction()

	const x int = 10
	const y = 5
	const z = x + y
	fmt.Println(z)

	const Question = "A : You can't change me."
	const Answer = "Are you sure?"

	fmt.Println(Question , Answer)
	fmt.Println("-----------------------")
	fmt.Println(" FOR Loop Training ")

	// FOR <Start Values>; <Condition> <End Values> {
	//  <Code Worker>
	//}
	for i := 0; i <= 100; i++ {
		if i == 5 {
			continue
		} else if i == 10 {
			break
		}
		fmt.Println(i)
	}

	// while in go
	// for <condition> {
	//  <Code Worker>
	// }
	// for <condition2> {
	// <Code Worker2>
	// }
	fmt.Println("-----------------------")
	fmt.Println(" While Loop Training ")
	fmt.Println("-----------------------")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("-----------------------")
	for {
		fmt.Println("loop")
		break
	}
	
	fmt.Println("-----------------------")
	fmt.Println(" Do While Training ")
	fmt.Println("-----------------------")
	l := 0
	for ok := true; ok; ok = l > 0 {
		fmt.Println(l)
		l--
	}
	fmt.Println("-----------------------")
	j := 0
	for {
		fmt.Println(j)
		if !(j > 0) {
			break
		}
	}

	fmt.Println("- End My Training Day#1 -")
}
