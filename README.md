<div align="center">
<pre>
                                                                            _            
                                                                           | |           
   ___ _   _ _ __ _ __ ___ _ __   ___ _   _    ___ ___  _ ____   _____ _ __| |_ ___ _ __ 
  / __| | | | '__| '__/ _ \ '_ \ / __| | | |  / __/ _ \| '_ \ \ / / _ \ '__| __/ _ \ '__|
 | (__| |_| | |  | | |  __/ | | | (__| |_| | | (_| (_) | | | \ V /  __/ |  | ||  __/ |   
  \___|\__,_|_|  |_|  \___|_| |_|\___|\__, |  \___\___/|_| |_|\_/ \___|_|   \__\___|_|   
                                       __/ |                                             
                                      |___/                                              

-----------------------------------------------------------------------------------------
Currency Converter CLI program written in Go with natural-language interpretation.

</pre>

![GitHub top language](https://img.shields.io/github/languages/top/ItsThompson/currency-converter)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/ItsThompson/currency-converter/main)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/ItsThompson/currency-converter)

</div>

## Introduction
This project is a command-line tool that uses the [Open Exchange Rates API](https://openexchangerates.org/) to provide currency conversion. With a focus on user-friendly interactions, this program uses natural language interpretation for input. To optimize performance and prevent excessive API calls, the CLI features a caching mechanism with a 5-minute expiry time. Additionally, there are unit tests for the natural language interpretation ensuring the reliability of the feature using the `testing` package. My main objectives of this project was to get used to the Go programming language and its web capabilities.

## Showcase

https://github.com/ItsThompson/currency-converter/assets/33135895/6c4f1185-3384-4006-8f0d-93c01abf8deb

## Usage
### Getting Started 
1. Sign Up on [Open Exchange Rate](https://openexchangerates.org/)
2. Set your app id as an environment variable by creating a .env file and assign `APP_ID` to your app id.
```
# .env

APP_ID=<APP_ID>
```
3. In addition to the `APP_ID`, add the following constants
    - `FILE_NAME`: Cache file name.
    - `CACHE_EXPIRY_IN_SECONDS`: Expiry time for cache in seconds.
    - `PRECISION`: Level of precision in output's digits.

```
# Example .env

APP_ID=<APP_ID>
FILE_NAME=cache.json
CACHE_EXPIRY_IN_SECONDS=300
PRECISION=2
```
4. Run the CLI program using `go run .`
### Running Executable
1. Build the executable 
```
go build
```
2. Once you compiled the code into a executable, you can directly run the executable (`./currencyconverter` on Unix systems, `currencyconverter.exe` on Windows).
