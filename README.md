# MOSINT

<p align="center">
  <img src="https://raw.githubusercontent.com/alpkeskin/mosint/master/banner2-3.png" width="500" title="mosint">
</p>

<<<<<<< HEAD
<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#services">Services</a> •
  <a href="#usage">Usage</a> •
  <a href="#docker">Docker</a> •
  <a href="#configuration-file">Config</a> •
  <a href="#screen">Screen</a>
</p>
=======
## What is the MOSINT
>>>>>>> parent of 09e81f3 (Update README.md)

MOSINT is a fastest OSINT Tool for emails. It helps you gather information about the target email.

#### Features: :eyes:

* Email validation
* Check social accounts
* Check data breaches and password leaks
* Find related emails and domains
* Scan Pastebin Dumps
* Google Search
* DNS/IP Lookup


## Services (APIs):

| Service | Function | Status |
| :--- | :--- | :--- |
| [ipapi.co](https://ipapi.co/) - Public | More Information About Domain | :white\_check\_mark: |
| [hunter.io](https://hunter.io/) - Public | Related Emails | :white\_check\_mark: :key: |
| [emailrep.io](https://emailrep.io/) - Public | Breached Sites Names | :white\_check\_mark: :key: |
| [scylla.so](https://scylla.so/) - Public | Database Leaks | :construction: |
| [psbdmp.ws](https://psbdmp.ws/) - Public | Pastebin Dumps | :white\_check\_mark: :key: |
| [Intelligence X](https://intelx.io/)| Password Leaks | :white\_check\_mark: :key: |
| [BreachDirectory](https://breachdirectory.org/)| Password Leaks | :white\_check\_mark: :key: |

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
```
go install -v github.com/alpkeskin/mosint@latest
```

<<<<<<< HEAD
# Usage
```sh
mosint example@email.com
```
Call the help (`-h`) flag for more information on usage.

# Docker

Build a docker image
```sh
docker build -t mosint .  
```
Run the docker container using the image
```sh
docker run mosint --help
```

# Screen
=======
## Usage:
```
mosint example@email.com
```

## Screen :
>>>>>>> parent of 09e81f3 (Update README.md)

[![mosint](https://asciinema.org/a/529726.svg)](https://asciinema.org/a/529726)


