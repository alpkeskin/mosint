# MOSINT

<p align="center">
  <img src="https://raw.githubusercontent.com/alpkeskin/mosint/master/banner.png" width="500" title="mosint">
</p>

## What is the MOSINT :question:

MOSINT is an OSINT Tool for emails. It helps you gather information about the target email.

#### Features: :eyes:

* Email validation
* Check social accounts with Socialscan
* Check data breaches
* Find related emails
* Find related domains
* Scan Pastebin Dumps
* Google Search
* DNS Lookup
* IP Lookup
* Find subdomains of domain 


## Services (APIs): :key:

\[not required to run the program\]

| Service | Function | Status |
| :--- | :--- | :--- |
| [ipapi.co](https://ipapi.co/) - Public | More Information About Domain | :white\_check\_mark: |
| [hunter.io](https://hunter.io/) - Public | Related Emails | :white\_check\_mark: :key: |
| [emailrep.io](https://emailrep.io/) - Public | Breached Sites Names | :white\_check\_mark: :key: |
| [scylla.so](https://scylla.so/) - Public | Database Leaks | :construction: |
| [breachdirectory.org](https://breachdirectory.org/) - Public | Database Leaks | :white\_check\_mark: :key: |

_- API key required_

#### For Use:

- Save your API key in the `keys.json`
- Install Go and Python on your system

## Installation:

`git clone https://github.com/alpkeskin/mosint.git`

`cd mosint`

`pip3 install -r requirements.txt`

## Usage:

you can type `-h` for help menu.

```text
+-------+--------------------------------+------------+
| FLAGS |          DESCRIPTION           | ISREQUIRED |
+-------+--------------------------------+------------+
| -e    | Set target email               | Yes        |
| -v    | Verify the target email        | No         |
| -ss   | Social scan for target email   | No         |
| -re   | Find related emails with       | No         |
|       | target email                   |            |
| -rd   | Find related domains with      | No         |
|       | target email                   |            |
| -l    | Find password leaks for target | No         |
|       | email                          |            |
| -pd   | Search pastebin dumps for      | No         |
|       | target email                   |            |
| -er   | EmailRep.io API                | No         |
| -d    | More information about target  | No         |
|       | email's domain                 |            |
| -all  | All features!                  | No         |
+-------+--------------------------------+------------+
```

### Example:

`go run main.go -e example@domain.com -all`


## Screen :

[![mosint](https://asciinema.org/a/444753.svg)](https://asciinema.org/a/444753)

### My Bitcoin Wallet: :money_with_wings:
[BTC]
`19N6A1yAGcfLpaFGQtWaVf316ETWweRbUo`

#### Tested on:

- [x] Linux
- [x] macOS

#### To-Do list: :memo:

- PDF Scanner for Related Emails
- Output file (.txt)
- Related phone number sources
- Useful API's

