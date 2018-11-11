# ING CSV to YNAB

This CLI tool aims to convert ING bank statements in CSV to a format that YNAB is able to understand.

## Disclaimer
This is being done in my spare time and I have completely no associations with ING or YNAB.
Also, quality was not the initial focus so obviously this code can use improvements - I'm happy to receive PRs!

## Prerequisites
Go installed and environment properly configured (think of `$GOPATH` and such). 

## Building it
Just run `go get github.com/viniciushana/ing-csv-ynab` to fetch the sources, then navigate to the folder where this tool was cloned to and run `go build` to build the binary.

## Running it
Calling `ing-csv-ynab` will display all the available options, and they're currently one: `convert`.

Calling `ing-csv-ynab convert` will display the details needed to run the command, and they're basically two:
* An input CSV file, which is the bank statement you got from ING
* An output CSV file, which will be written with the conversion results


