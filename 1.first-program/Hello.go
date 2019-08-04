//package main is special package
//it allows Go to create an executable file
package main

/*
This is multi-line comment.

import keyword makes another package available
  for this .go "file".

import "fmt" lets you acces fmt package's functionality
  here in this .go "file".
*/

import "fmt"

//"func main" is special.
//
// Go has to know where to start
//
// func main creates a starting point for Go
//
//After compiling a starting pint for Go
//Go runtime will first run this function
func main() {
	// after: import "fmt"
	// Println function of "fmt" package become available

	// Look at what it looks like by typing in the console:
	// go doc -src fmt Println

	//Println is just an exported function form
	// "fmt" package

	// Exported = first Letter is uppercase
	fmt.Println("Hello Gopher !")

	//Go cannot call Println function by itself.
	//That's why you need to call it here.
	//It only calls 'func main' automatically

	//---

	// Go supports Unicode characters in string litersals
	// ANd also in source-cede: KOSTEBEK!
	//
	// Because: Literal ~= Source Code

}
