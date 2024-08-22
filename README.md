# Konverta
The Konverta app is a command-line utility designed to help users easily convert currencies and manage currency-related information.<br>

Available Options:
```
Usage:
 Konverta [options]

Options:
 -v, --version                                                                           : Version of Konverta
 -h, --help                                                                              : Getting Help
 -c, --convert --fromto [currency_from e.g. USD] [currency_to e.g EUR] [amount e.g 200]  : Convert amount from a currency to another
 -s, --symbols                                                                           : List all available currency symbols

Konverta created and developed by nbebaw: https://github.com/nbebaw
```

Dependency:
- Rapidapi Key
- Subscribe https://rapidapi.com/principalapis/api/currency-conversion-and-exchange-rates

### Setting Up Environment Variables
#### For Windows
```
setx API_KEY "your_api_key_here"
```
```
Using PowerShell:
[Environment]::SetEnvironmentVariable("API_KEY", "your_api_key_here", "User")
```
#### For macOS and Linux
##### Temporary (current session only):
```
export API_KEY="your_api_key_here"
```

##### Permanent:
Add the export command to your shell's profile script (like ~/.bashrc, ~/.bash_profile, ~/.zshrc, etc.):
```
echo 'export API_KEY="your_api_key_here"' >> ~/.bashrc
source ~/.bashrc
```