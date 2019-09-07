package main

import "fmt"
import "time"

//Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

//Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

var (
	toBe        = false
	luckyNumber = 7
)

const Pi = 3.14 //example of constant. Constants cannot be declared using the := syntax.

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) { //function can return multiple results
	return y, x
}

func split(sum int) (x, y int) { //you can name the return value. Naked return statements should be used only in short functions,
	//as with the example shown here. They can harm readability in longer functions.
	x = sum * 4 / 9
	y = sum - x
	return
}

func forLoops() { //for loops . The init and post statements are optional.
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
}

//main function
//key points:
//			1. You can't have unused variables
//			2. blank identifier is like a black hole for variables

//-------------------------------------------------------------------------------------------------------
//functions

func addition(a int, b int) int {
	return a + b
}

//you can omit the type name for the like types parameters
func addition2(a, b, c int) int {
	return a + b + c
}

//multi returns
func vals() (int, int) {
	return 3, 7
}
func main() {
	s, t := swap("flying", "pigs")
	fmt.Println(add(42, 8))
	fmt.Println(s, t)
	fmt.Println(split(27))
	fmt.Println(Pi)
	fmt.Println(forLoops)

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	res := addition(21, 23)
	res2 := addition2(2, 4, 6)
	fmt.Println(res, res2)

	//printing to the screen
	fmt.Println("yello, world!")
	//ints addition
	fmt.Println("1+1=", 1+1)
	//ints division
	fmt.Println("7.0/3.0", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	//variables: var declares one or more variables. go will infer the type of initialized variables
	//variables declared without a corresponding initialization are zero valued.
	//the := syntax is shorthand for declaring and initializing a variable  var f string = "apple" in the below example

	//You can assign two different types of objects
	//In this case "go" understands the type without us saying it, just like in python
	var a, d = "initial", 42
	fmt.Println(a, d)

	//you can declare the type at the end of the initialization
	var b, c int = 1, 2
	fmt.Println(b, c)

	//if no initialization, go assigns "zero" values.
	var e int
	fmt.Println(e)

	//shorthand assign
	f := "apple"
	fmt.Println(f)

	//"const" declares a constant value. It can appear anywhere a var statement can
	const CONSTANT string = "constant"
	const n = 5000000
	fmt.Println(s, n)
	//-------------------------------------------------------------------------------------------------------

	//looping. "for" is the only looping construct. There are 3 basic types of for loops
	//1. with a single condition.
	//2. A classic initial/condition/after for loop
	//3. "for" without a condition will loop repeatedly until you break out of the loop or return from the enclosing
	//function

	//this is with one condition
	num := 1
	for num <= 3 {
		fmt.Println(num)
		num = num + 1
	}

	//classic for loop
	for i := 7; i < 9; i++ {
		fmt.Println(i)
	}

	fmt.Println("------------")

	//for loop without a condition. This will print odd numbers upto 5
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	//-------------------------------------------------------------------------------------------------------

	//Branching: If else in go is pretty straightforward. It's similar to Java
	//only difference is we don't need to put brackets for the condition

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	fmt.Println("------------")

	//here num:=9 initializes num and num<0 is the first condition
	//any variables declared in the first if statement are available in all branches
	if num := 9; num < 0 {
		fmt.Println(num, " is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	fmt.Println("------------")

	//Switch statements

	//First kind
	k := 2
	fmt.Print("Write ", k, " as ")
	switch k {
	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")
	}

	//second kind.  You can apply  two conditions
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}
	fmt.Println("------------")
	//-------------------------------------------------------------------------------------------------------

	//Array : Array in go is a numbered sequence of elements of a specific length

	//arrays that holds 5 ints. here, it will assign 0 to all 5 ints
	var ar [5]int
	fmt.Println("emp:", ar)

	ar[4] = 100
	fmt.Println("set:", ar)

	//you can get elements with index
	fmt.Println("get:", ar[4])

	//get the length of the array
	fmt.Println(len(ar))

	//declare and initialize an array
	br := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", br)

	//arrays are 1-dimensional but you can declare 2-dimensional like this
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
	fmt.Println("------------")
	//-------------------------------------------------------------------------------------------------------
	//slices: slices are key data type in Go, giving more powerful interface to sequences than arrays

	slice := make([]string, 3) //this creates a slice with non-zero length, using the builtin "make"
	fmt.Println("emp:", s)

	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"
	fmt.Println("set:", slice)
	fmt.Println("get:", slice[2])

	fmt.Println("len:", len(s))

	//"append" returns a slices containing one or more new values
	slice = append(slice, "d")
	slice = append(slice, "e", "f")
	fmt.Println("append:", slice)

	//you can also copy one slices to another
	copySlice := make([]string, len(slice)) //create a new sliceice
	copy(copySlice, slice)
	fmt.Println("cslice:", copySlice)

	//you can get slice of the element. Below, we 'll get slice[2], slice[3], slice[4]
	slicesOp := slice[2:5]
	fmt.Println("slicesOp:", slicesOp)

	//this is slices up to (but excluding) sllice[5]
	//also note we're not using :=, instead we are using "="" because we are rassigning new value to the already created variable
	slicesOp = slice[:5]
	fmt.Println("slicesOp2:", slicesOp)

	//this is slices upfrom (and including slice[2])
	slicesOp = slice[2:]
	fmt.Println("slicesOp2:", slicesOp)

	//declare and initialize slice in a single line
	newSlice := []string{"g", "h", "i"}
	fmt.Println("newSlice:", newSlice)

	//multi-dimensional data structures using slices
	twoDSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDSlice[i][j] = i + j
		}
	}
	fmt.Println("2D-Slice:", twoDSlice)
	//-------------------------------------------------------------------------------------------------------
	//Maps: Mars in go is a builtin associative data type. Similar to hashes or dicts in other languages

	//create an empty map: make(map[key-type] val-type)
	vMap := make(map[string]int)

	//setting up the value of map is just name[key] = val
	vMap["key1"] = 7
	vMap["key2"] = 42
	fmt.Println(vMap)

	//get a value for a key
	value1 := vMap["key1"]
	fmt.Println(value1)

	//builtin delete removes key/value from a map
	delete(vMap, "key1")
	fmt.Println(vMap)

	//todo:need to understand this
	_, prs := vMap["key2"]
	fmt.Println("prs:", prs)

	newMap := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", newMap)
	//-------------------------------------------------------------------------------------------------------
	//Range: it iterates over elements in a variety of data structures.

	//e.g. sum the numbers in a slice
	nums := []int{1, 2, 3}
	sum := 0

	//in slices, range returns two values, first the index, second is the a copy of the element at that index
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	//range in strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second
	// the rune itself

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
