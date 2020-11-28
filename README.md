# MOSINT


## :question: What is the MOSINT ?
MOSINT is an OSINT Tool for emails.
It helps you gather information about the target email.
#### :briefcase: Features:
  - Verification Service { Check if email exist }
  - Check social accounts with Socialscan
  - Check data breaches
  - [need API] Find related emails
  - Find related phone numbers
  - Find related domains
  - Scan Pastebin Dumps
  - Google Search
  - DNS Lookup
  
  :bangbang: You can turn features on off from the `config.json` :bangbang:
  
  ```json
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

## :key: APIs:
[not required to run the program]

| Service | Function | Status |
|-|-|-|
| [verify-email](https://verify-email.org/) | Email Verification | :white_check_mark: :key: |
| [hunter.io](https://hunter.io/) - Public | Related Emails | :white_check_mark: :key: |
| [leak-lookup](https://leak-lookup.com/) | Breached Sites Names | :white_check_mark: :key: |
| [scylla.sh](https://scylla.sh/) | Database Leaks | :white_check_mark: |
| [hackertarget](https://hackertarget.com/) | DNS Lookup | :white_check_mark: |
| [psbdmp](https://psbdmp.ws/) | Pastebin Dumps | :white_check_mark: |

*:key: - API key required* 

#### For Use:
Save your API key in the `config.json`


## :package: Cloning:
`git clone https://github.com/alpkeskin/mosint.git`

## :shipit: Usage:
`cd mosint`

`pip3 install -r requirements.txt`

- You can edit the `config.json` file

`python3 mosint.py`

- Set Target Email

Also, you can **exit** by pressing the `q` key.

## :computer: Screen:

[![mosint](https://asciinema.org/a/2vXl00ACUTpPULeQsYcDiFsXy.svg)](https://asciinema.org/a/2vXl00ACUTpPULeQsYcDiFsXy)

### :money_with_wings: My Bitcoin Wallet:
`3NFfd1QXUVFsZzfbwGJiAJdehtPB9D88tK`

#### :white_check_mark: Tested on:
- Kali Linux
- Parrot OS
- MacOS (without SocialScan module)
