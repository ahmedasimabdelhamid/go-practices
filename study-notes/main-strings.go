package main 
import (
	"fmt"
	"strings"
)

func main () {

	s1 := "Hi Go!"  // what's inside the quoates called string literal
	fmt.Println(s1)

	fmt.Println("He Says \"Hello!\"")  // backsalsh as escaping char
	fmt.Println(`He Says "Hello! "`)  // backtics can be used also, a string literral in backticks called raw string
	fmt.Println(`He Says \n Hello! `)  // raw string, \n will be printed
	fmt.Printf("%s\n", s1)  // %s for string
	fmt.Printf("%q\n", s1)  // %q for quoated string

    p := fmt.Println   // println as alias alias 
	result := strings.Contains("I love Go!", "love")  // check if string contains specific string
	p(result)
	

	result = strings.ContainsAny("success", "xys")  // check if string has x or y
	p(result)

	result = strings.ContainsRune("golang", 'g')   // check if rune g is within golang, rune has to be inside ''
    p(result)

	n := strings.Count("cheese", "e") // how many e in cheese
    p(n)
    
    // if the search was empty the count will add 1 to the total rune of the string

	n = strings.Count("Five", "")
    p(n)


	p(strings.ToLower("Go python Java"))
	p(strings.ToUpper("Go python Java"))

    p("go" == "go")  // compare two strings, true
	p("Go" == "go") // false
    p(strings.ToLower("GO") == strings.ToLower("go"))  // convert them to lower for better comparision

    p(strings.EqualFold("Go", "go"))  // best solution if we don't care about case sensitivty
    p(strings.EqualFold("Go", "gO"))  // best solution if we don't care about case sensitivty
	myStr := strings.Repeat("##", 10)  // rebeat string 10 times
	p(myStr)
    
	myStr = strings.Replace("192.168.0.1", ".", ":", 4) // if -1 it will replace all values
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", -1) // if -1 it will replace all values
	p(myStr)

	myStr = strings.ReplaceAll("192.168.0.1", ".", ":") // replace all will replace all . with : no need to use int
	p(myStr)
    
	s := strings.Split("a,b,c", ",")   // convert the sting into slice of strings separator is ,
    fmt.Printf("%T\n", s)    // %T print the type
	fmt.Printf("%#v\n", s)   // %v print the value of slice
 
	s = strings.Split("I love Go programming", "")   // convert the sting into slice of strings, separator is space
	fmt.Printf("%#v\n", s)   // %v print the value of slice
    
	s = []string{"I", "learn", "Golang"}
	myStr = strings.Join(s, "-")  // output I-learn-Golang
	p(myStr)


	myStr = "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr)  // will rerun slice of strings

	fmt.Printf("%T\n", fields)    // %T print the type
	fmt.Printf("%#v\n", fields)   // %v print the value 

	// Trimspce returns slice of strings with leading and trailing space removed
	myStr = strings.TrimSpace("\t Goodby Windows, Welcome Linux")  // \t for tab
	
	fmt.Printf("%T\n", myStr)    // %T print the type
	fmt.Printf("%#v\n", myStr)   // %v print the value 
	fmt.Printf("%q\n", myStr )   // %q print the quoated string


	myStr = strings.Trim("....Hello Gophers!!!!?", ".!?") // remove any .!? from the string
	fmt.Printf("%q\n", myStr )   // %q print the quoated string

}