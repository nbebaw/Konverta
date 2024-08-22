package lib

import "fmt"

// GetHelp displays the help information and usage instructions for the application.
func GetHelp() {
	fmt.Println("Usage:")
	fmt.Println(" Konverta [options]")
	fmt.Println("\nOptions:")
	fmt.Printf(" -v, --version \t\t\t\t\t\t\t\t\t\t : Version of Konverta\n")
	fmt.Printf(" -h, --help \t\t\t\t\t\t\t\t\t\t : Getting Help\n")
	fmt.Printf(" -c, --convert --fromto [currency_from e.g. USD] [currency_to e.g EUR] [amount e.g 200]\t : Convert amount from a currency to another\n")
	fmt.Printf(" -s, --symbols \t\t\t\t\t\t\t\t\t\t : List all available currency symbols\n")
	fmt.Println("")
	fmt.Println("Konverta created and developed by nbebaw: https://github.com/nbebaw")
}
