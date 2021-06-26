# MOSINT

## What is the MOSINT ?

MOSINT is an OSINT Tool for emails. It helps you gather information about the target email.

#### Features:

* Verification Service { Check if email exist }
* Check social accounts with Socialscan
* Check data breaches
* \[need API\] Find related emails
* Find related phone numbers
* Find related domains
* Scan Pastebin Dumps
* Google Search
* DNS Lookup



You can turn features on off from the `config.json` 

```javascript
[
{
  "verify-email.org API Key": "set API KEY here",
  "hunter.io API Key": "set API KEY here",
  "Breached Sites[leak-lookup.com API Key]": "set API KEY here",
  "Social Scan": "True",
  "Leaked DB": "True",
  "Related Phone Numbers" : "True",
  "Related Domains" : "True",
  "Pastebin Dumps": "True",
  "Google Search": "True",
  "DNS Lookup": "True"
}
]
```

## APIs:

\[not required to run the program\]

| Service | Function | Status |
| :--- | :--- | :--- |
| [verify-email](https://verify-email.org/) | Email Verification | :white\_check\_mark: :key: |
| [hunter.io](https://hunter.io/) - Public | Related Emails | :white\_check\_mark: :key: |
| [leak-lookup](https://leak-lookup.com/) | Breached Sites Names | :white\_check\_mark: :key: |
| [scylla.sh](https://scylla.sh/) | Database Leaks | :white\_check\_mark: |
| [hackertarget](https://hackertarget.com/) | DNS Lookup | :white\_check\_mark: |
| [psbdmp](https://psbdmp.ws/) | Pastebin Dumps | :white\_check\_mark: |

_- API key required_

#### For Use:

Save your API key in the `config.json`

## Cloning:

`git clone https://github.com/alpkeskin/mosint.git`

## Usage:

`cd mosint`

`pip3 install -r requirements.txt`

* You can edit the `config.json` file

`python3 mosint.py`

* Set Target Email

Also, you can **exit** by pressing the `q` key.

## Screen:

[![mosint](https://asciinema.org/a/2vXl00ACUTpPULeQsYcDiFsXy.svg)](https://asciinema.org/a/2vXl00ACUTpPULeQsYcDiFsXy)

### My Bitcoin Wallet:

`3NFfd1QXUVFsZzfbwGJiAJdehtPB9D88tK`

#### Tested on:

* Kali Linux
* Parrot OS
* MacOS \(without SocialScan module\)

