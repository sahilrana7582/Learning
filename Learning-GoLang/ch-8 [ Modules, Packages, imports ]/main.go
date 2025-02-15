package main

import (
	"github.com/Learning-GoLang/math"
)

/*
	Modules, Packages, Imports Chapter:

	Modules are the collection of packages.
	Packages are the collection of files.
	Files are the collection of functions, variables, and constants.

	The Module is the central respositry for a project.

	go.mod: You can see we have a go.mod file in our root directory.

	it make our project into the modules.

	go.mod file contains the following information:
	module name
	go version
	replace
	dependencies

	how to make our project into the module: By running the command "go mod init" in the root directory.

*/

/*
	Import / Export:

	Import: Importing something from another file into the current file.
	Export: Exporting something from the current file to another file.

	How Importing Works:
	Go compiler will search for the exported constants from other modules, files, in order:
	1. Current Directory
	2. All Directories in the GOPATH
	3. All Directories in the PATH environment variable

	Constants = variables, functions, structs, interfaces, etc

	How Exporting Works in GoLang:
	Exported Constants:
	1. Start with a capital letter
	2. Defined outside the function

	Exported Functions:

	We can use the export functionality of the function by using the capital letter in the constant name.

*/

func main() {
	math.Add(5, 5)

}
