# Konverta

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