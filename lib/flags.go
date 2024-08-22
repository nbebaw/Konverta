package lib

import "flag"

// DefineFlags sets up and defines command-line flags for the application.
func DefineFlags() {
	// Allowed main options
	// Define the -version or -v flag to show version information
	flag.BoolVar(&ShowVersion, "version", false, "Show version information") // Long form: --version
	flag.BoolVar(&ShowVersionAlias, "v", false, "Show version information")  // Short form: -v

	// Define the -help or -h flag to show help information
	flag.BoolVar(&Help, "help", false, "Show help")   // Long form: --help
	flag.BoolVar(&HelpAlias, "h", false, "Show help") // Short form: -h

	// Define the -convert or -c flag to convert currency
	flag.BoolVar(&Convert, "convert", false, "Convert currency") // Long form: --convert
	flag.BoolVar(&ConvertAlias, "c", false, "Convert currency")  // Short form: -c
	// Define the -fromto flag for specifying currency conversion direction
	flag.BoolVar(&CurrencyFromTo, "fromto", false, "Specify currency conversion direction")

	// Define the -symbols or -s flag to list all available currencies
	flag.BoolVar(&Symbols, "symbols", false, "List all available Currency") // Long form: --symbols
	flag.BoolVar(&SymbolsAlias, "s", false, "List all available Currency")  // Short form: -s

	// Set the custom usage function to display help information
	flag.Usage = GetHelp
	// Parse the command-line flags
	flag.Parse()
}
