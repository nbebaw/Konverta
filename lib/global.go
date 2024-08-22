package lib

// Rapidapi_host is a constant that holds the RapidAPI host URL for the currency conversion and exchange rates API.
const Rapidapi_host string = "currency-conversion-and-exchange-rates.p.rapidapi.com"

// ApiKey is a variable that stores the API key retrieved from the environment variables.
var ApiKey string = getApiKey()

// Declaration of boolean flags used throughout the application.
// These flags are used to determine which operation the user wants to perform.
var (
	Help             bool // Help indicates if the help flag is set.
	HelpAlias        bool // HelpAlias is the alias for the help flag (-h).
	ShowVersion      bool // ShowVersion indicates if the version flag is set.
	ShowVersionAlias bool // ShowVersionAlias is the alias for the version flag (-v).
	Convert          bool // Convert indicates if the convert currency flag is set.
	ConvertAlias     bool // ConvertAlias is the alias for the convert flag (-c).
	Symbols          bool // Symbols indicates if the list symbols flag is set.
	SymbolsAlias     bool // SymbolsAlias is the alias for the symbols flag (-s).
	CurrencyFromTo   bool // CurrencyFromTo indicates if the currency conversion direction flag is set.
)
