# MOSINT

<p align="center">
  <img src="https://raw.githubusercontent.com/alpkeskin/mosint/master/banner.png" width="500" title="mosint">
</p>

## What is the MOSINT

MOSINT is an OSINT Tool for emails. It helps you gather information about the target email.

#### Features: :eyes:

* Email validation
* Check social accounts with Socialscan and Holehe
* Check data breaches and password leaks
* Find related emails and domains
* Scan Pastebin and Throwbin Dumps
* Google Search
* DNS Lookup
* IP Lookup
* Output to text file


## Services (APIs):

\[not required to run the program\]

| Service | Function | Status |
| :--- | :--- | :--- |
| [ipapi.co](https://ipapi.co/) - Public | More Information About Domain | :white\_check\_mark: |
| [hunter.io](https://hunter.io/) - Public | Related Emails | :white\_check\_mark: :key: |
| [emailrep.io](https://emailrep.io/) - Public | Breached Sites Names | :white\_check\_mark: :key: |
| [scylla.so](https://scylla.so/) - Public | Database Leaks | :construction: |
| [breachdirectory.org](https://breachdirectory.org/) - Public | Password Leaks | :white\_check\_mark: :key: |
| [Intelligence X](https://intelx.io/)| Password Leaks | :white\_check\_mark: :key: |

:key: API key required

#### For Use:

- Save your API key in the `keys.json`
- Install Go and Python on your system

## Installation:

`git clone https://github.com/alpkeskin/mosint.git`

`cd mosint`

`pip3 install -r requirements.txt`

## Usage:

you can type `-h` for help menu.

| FLAGS     | DESCRIPTION                                       | ISREQUIRED |
|-----------|---------------------------------------------------|------------|
| -e        | Set target email                                  | Yes        |
| -verify   | Verify target email                               | No         |
| -social   | Social scan for target email                      | No         |
| -relateds | Find related emails and domains with target email | No         |
| -leaks    | Find password leaks for target email              | No         |
| -dumps    | Search pastebin dumps for target email            | No         |
| -domain   | More information about target email's domain      | No         |
| -o        | Output to text file                               | No         |
| -v        | Version of mosint                                 | No         |
| -h        | Help Menu                                         | No         |
| -all      | All features!                                     | No         |

### Example:

`go run main.go -e example@domain.com -all`

Just type `-o` for output file (.txt)


## Screen :

[![mosint](https://asciinema.org/a/479072.svg)](https://asciinema.org/a/479072)

### Buy me a coffee: :money_with_wings:

https://www.buymeacoffee.com/alpkeskin

#### Tested on:

- [x] Linux
- [x] macOS
