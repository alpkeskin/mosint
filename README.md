# MOSINT

<p align="center">
  <img src="https://raw.githubusercontent.com/alpkeskin/mosint/master/banner2-2.png" width="500" title="mosint">
</p>

## What is the MOSINT

MOSINT is a fastest OSINT Tool for emails. It helps you gather information about the target email.

#### Features: :eyes:

* Email validation
* Check social accounts
* Check data breaches and password leaks
* Find related emails and domains
* Scan Pastebin Dumps
* Google Search
* DNS Lookup


## Services (APIs):

\[not required to run the program\]

| Service | Function | Status |
| :--- | :--- | :--- |
| [ipapi.co](https://ipapi.co/) - Public | More Information About Domain | :construction: |
| [hunter.io](https://hunter.io/) - Public | Related Emails | :white\_check\_mark: :key: |
| [emailrep.io](https://emailrep.io/) - Public | Breached Sites Names | :white\_check\_mark: :key: |
| [scylla.so](https://scylla.so/) - Public | Database Leaks | :construction: |
| [psbdmp.ws](https://psbdmp.ws/) - Public | Pastebin Dumps | :white\_check\_mark: :key: |
| [Intelligence X](https://intelx.io/)| Password Leaks | :white\_check\_mark: :key: |

:key: API key required

#### If you want to use mosint with full features, set your API keys:

 ```
  mosint set hunter <hunter.io API key>
  mosint set emailrep <emailrep.io API key>
  mosint set intelx <intelx.io API key>
  mosint set psbdmp <psbdmp.ws API key>
  mosint set breachdirectory <breachdirectory.org API key>
  ```

## Installation:

`go install -v github.com/alpkeskin/mosint@latest`

## Usage:

`mosint example@email.com`

## Screen :

[![mosint](https://asciinema.org/a/529726.svg)](https://asciinema.org/a/529726)


