package main

import "fmt"


// func <Function Name> (<Variable Value> <Type Input Value>) <Type of Variable output> {
// ...
// <Code Working>
// ...
// return <Result>
// }

func plus(num1 int, num2 int) int {
	return num1 + num2
}

func prism(length, width, height int) int  {
	return length * width * height
}

func intit() {
	fmt.Println("Init ... ")
}

func main() {

	fmt.Println("DAY1 Start ... ")
	fmt.Println(" Hello World! Golang ")
	fmt.Println("-----------------------")



	fmt.Println("DAY2 Start ... ")

	// Variables and Integer
	// String and Const
	// message

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

	fmt.Println(" \n ")

	fmt.Println("- Start My Training Day#2 -")

	//  if <condition <Valuse>> <Decide> {
	//   < Code Worker >
	//   } else if {
	//   < Code Worker >
	//   else { 
	//   < Code Worker >
	//  }

	if 10 - 5 == 5 {
		fmt.Println(" Condition is true. ")
	} else {
		fmt.Println(" Condition is false. ")
	}

	if applause := 10; applause <= 3 {
		fmt.Println(" Cool!! ")
	} else {
		fmt.Println(" Awesome!! ")
	}
	fmt.Println("----------------------- \n ")

	// Switch < Condition < Values >>; < Variables >
	// Case < Variables Values > 
	// < Code Worker >
	// < Fallthough >
	// Default:
	// < Code Worker >


	// method 1
	n := 2
	fmt.Println(" Write ",n," as ")
	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// method 2 fallthrough will not exit switch.It will continue case
	k := 3
	switch k {
	case 3:
		fmt.Println("three")
		fallthrough
	case 2:
		fmt.Println("two")
		fallthrough
	case 1:
		fmt.Println("one")
		fallthrough
	default:
		fmt.Println(" Go!!! ")
	}

	fmt.Println("----------------------- \n ")

// method 3
	app := 10
	switch {
	case app <= 3:
		fmt.Println(" Good! ")
	case app <= 5:
		fmt.Println(" Cool! ")
	default:
		fmt.Println(" Awesome!! ")
	}

	fmt.Println("----------------------- \n ")
	fmt.Println("  Arrays Training \n ")

	// var < Array Name > [<Array Scaling>] < Type of Array >
	//     < Array Name > := [<Array Scaling>] <Type of Array> {<mention Values>}
	//
	//     < Array Name > [<Array Scaling> <Array Scaling>]<Type of Array>

	var a[3]int
	b := [5]int{5, 5, 5, 5}

	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(" Length of an array a: ",len(a))
	fmt.Println(" Values of an array a: ", a[0], a[1], a[2])
	fmt.Println(" Values of an array b: ",b)

	var twoD [2][3]int
	for g :=0; g < 2; g++ {
		for o := 0; o < 3; o++ {
			twoD[g][o] = g + o
		}
	}
	fmt.Println(" 2nd arrays: ", twoD)

	fmt.Println("  ----------------------- \n ")
	fmt.Println("  Slice Training \n ")

	// <Slice Name> := make([] <type of Slice>, <default Values>)
	// <Slice Name> := []string{<default Values>, <default Values>}

	c := make([]string, 3)

	c[0] = "a"
	c[1] = "b"
	c[2] = "c"
	fmt.Println(" Length of slice c: ", len(c))
	fmt.Println(" Values of slice c: ", c[0], c[1], c[2])
	fmt.Println("  ----------------------- \n ")

	// more info : https://blog.golang.org/go-slices-usage-and-internals
	
	// <Slice Name>[<default position>:<end position>]
	
	e := []string{"a", "b", "c"}
	f := make([]string, len(e))
	copy(f, e)
	f = append(f, "d", "e")

	fmt.Println(" Value of slice e: ", e)
	fmt.Println(" Value of slice f: ", f)

	h := f[0:len(f)]
	fmt.Println(" Value of slice h: ", h)
	fmt.Println(" ---- Mutate Value of C ---- ")
	h[0] = "x"

	fmt.Println(" Value of slice h: ", h)
	fmt.Println(" Value of slice f: ", f)

	fmt.Println("  ----------------------- \n ")

	q := []string{"a", "b", "c", "d", "e"}
	w := q[0:5]
	v := q[:5]
	r := q[0:]
	t := q[:]

	fmt.Println(" Value of slice q: ",q)
	fmt.Println(" Value of slice w: ",w)
	fmt.Println(" Value of slice v: ",v)
	fmt.Println(" Value of slice r: ",r)
	fmt.Println(" Value of slice t: ",t)

	fmt.Println("  ----------------------- \n ")
	p := q[2:4]
	fmt.Println(" Value of slice p: ", p)

	fmt.Println("  ----------------------- \n ")

	// <maps's name> := make[(map<type of key>]<type of Value>)
	// <maps's name> := map[<type of key>]<type of Values>{<key>: <key>,...}

	Amap := make(map[string]int)
	Amap["math"] = 85
	Amap["sci"] = 92
	Amap["com"] = 28
	fmt.Println("Length of maps Amap: ",len(Amap))
	fmt.Println(" Values of Maps Amap: ", Amap)

	fmt.Println("  -----------End Mapping Mathod#1 ------------ \n ")
	fmt.Println("  -----------Start Mapping Mathod#2 ------------ \n ")

	Bmap := map[string]int{"math": 90,"sci": 90,"com": 100}
	fmt.Println(" Value of Maps Bmap: ",Bmap)
	delete(Bmap,"math")
	fmt.Println(" Delete math from map Bmap. ")
	fmt.Println(" Value of Map Bmap: ", Bmap)

	values, prs := Bmap["math"]
	fmt.Println(" State of math: ",prs)
	fmt.Println(" Value if math: ",values)

	fmt.Println("  ----------------------- \n ")
	if values, prs := Bmap["math"]; prs {
		fmt.Println(" Value of math: ", values)
	}
	if values, prs := Bmap["com"]; prs {
		fmt.Println(" Value of com: ", values)
	}

	fmt.Println("  ----------------------- \n ")
	fmt.Println("  -----------Start Mapping Mathod#3 ------------ \n ")

	a = [3]int{1, 2, 3}
	s := []int{1, 2, 3}
	m := map[string]int{"a": 1,"b": 2,"c": 3 }
	fmt.Println("  Value of l,s,u: ", a ,s, m)

	fmt.Println("  ----------------------- \n ")
	for i,num := range a {
		fmt.Println(i,num)
	}
	
	fmt.Println("  ----------------------- \n ")
	for i,num := range s {
		fmt.Println(i,num)
	}

	fmt.Println("  ----------------------- \n ")
	for i,num := range m {
		fmt.Println(i,num)
	}

	fmt.Println("  ----------- End Map Method ------------ \n ")
	fmt.Println("  ----------- End Day#3 Training ------------ \n ")

	fmt.Println("  ----------- Start Day#4 Function ------------ \n ")

	num1, num2 := 1, 2
	res := plus(num1, num2)
	fmt.Println(" 1 + 2 =", res)

	res = prism(4, 2, 2)
	fmt.Println(" Rectangular prism =", res)


	

	fmt.Println("  ----------------------- \n ")
	fmt.Println("  ----------- End Day#4 Function ------------ \n ")

}


