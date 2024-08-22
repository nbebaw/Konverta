package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
	"unicode"

	"github.com/nbebaw/konverta/lib"
)

func main() {
	// Define flags as a beginning
	lib.DefineFlags()

	switch {
	// --help || -h
	case lib.Help || lib.HelpAlias:
		if len(os.Args) > 2 {
			log.Println("Error: Illegal option or argument.")
			flag.Usage()
			return
		}
		lib.GetHelp()
		return
	// --symbols || -s
	case lib.Symbols || lib.SymbolsAlias:

		dataMap := lib.CallAPI("symbols")

		if query, ok := dataMap["symbols"].(map[string]interface{}); ok {
			for i, s := range query {
				fmt.Println(i, s)
			}
		}
		return
	// --convert || -c
	case lib.Convert || lib.ConvertAlias:
		if lib.CurrencyFromTo {
			if flag.NArg() > 0 {
				currencyFrom := flag.Arg(0)
				currencyFromCheck := false
				currencyTo := flag.Arg(1)
				currencyToCheck := false
				amount := flag.Arg(2)
				// Check if the user added the right currency symbols
				dataMap := lib.CallAPI("symbols")
				if query, ok := dataMap["symbols"].(map[string]interface{}); ok {
					for i := range query {
						if currencyFrom == i {
							currencyFromCheck = true
						} else if currencyTo == i {
							currencyToCheck = true
						}
					}
				}
				isDigitOnly := true
				// Check if the user amount input are only digits
				for _, char := range amount {
					if !unicode.IsDigit(char) {
						isDigitOnly = false
						break
					}
				}
				if isDigitOnly && currencyFromCheck && currencyToCheck {
					url := fmt.Sprintf("convert?from=%v&to=%v&amount=%v", currencyFrom, currencyTo, amount)
					dataMap := lib.CallAPI(url)
					if query, ok := dataMap["query"].(map[string]interface{}); ok {
						fmt.Println("From Currency:", query["from"])
						fmt.Println("To Currency:", query["to"])
						fmt.Println("Amount:", query["amount"], currencyFrom)
					}
					if info, ok := dataMap["info"].(map[string]interface{}); ok {
						fmt.Println("Conversion Rate:", info["rate"], currencyTo)
						if timestamp, ok := info["timestamp"].(float64); ok { // Ensure the timestamp is extracted correctly as float64
							// Convert timestamp to human-readable time
							humanReadableTime := time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
							fmt.Println("Timestamp:", humanReadableTime)
						}
					}
					fmt.Println("Date of Rate:", dataMap["date"])
					fmt.Println("Converted Amount:", dataMap["result"], currencyTo)
				}

				if !isDigitOnly {
					fmt.Println("Your amount input is: ", amount)
					fmt.Println("Amount should only digits.")
				} else if !currencyFromCheck {
					fmt.Println("Your currency from input is: ", currencyFrom)
					fmt.Println("This currency does not exists.")
					fmt.Println("Check the available currency: konverta --symbols")
				} else if !currencyToCheck {
					fmt.Println("Your currency from input is: ", currencyTo)
					fmt.Println("This currency does not exists.")
					fmt.Println("Check the available currency: konverta --symbols")
				}
			}
		} else {
			fmt.Println("Please provide an argument with " + os.Args[1])
			flag.Usage()
			os.Exit(1)
		}
		return
	}

	// Show help when user give only the main options
	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	// Show error and help when user write an illegal arguments/options
	if flag.NArg() > 0 {
		log.Println("Error: Illegal option or argument.")
		flag.Usage()
		return
	}
}
