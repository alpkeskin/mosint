[![forthebadge made-with-python](http://ForTheBadge.com/images/badges/made-with-python.svg)](https://www.python.org/)

# MOSINT
:up: Help me improve the tool

## :question: What is the MOSINT ?
MOSINT is an OSINT Tool for emails.
It helps you gather information about the target email.
#### :briefcase: Features:
  - [need API] Verification Service { Check if email exist }
  - Check social accounts with Socialscan
  - Check data breach
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

## :key: API Key:
**[not required to run the program]**

https://verify-email.org/ --> API Integration
You can create an API key for the verification feature.
Also "credit" data is get from the verification service.

https://hunter.io/ --> API Integration
You can show the emails related to the target email
#### For Use:
Save your API key in the `config.json`


## :package: Cloning:
`git clone https://github.com/alpkeskin/mosint.git`

## :shipit: Usage:
`cd mosint`

`pip3 install -r requirements.txt`

- You can edit the `Config.json` file

`python3 mosint.py`

- Set Target Email

Also, you can **exit** by pressing the `q` key.

## :computer: Screen:
```python
___  ________ _____ _____ _   _ _____ 
|  \/  |  _  /  ___|_   _| \ | |_   _|
| .  . | | | \ `--.  | | |  \| | | |  
| |\/| | | | |`--. \ | | | . ` | | |  
| |  | \ \_/ /\__/ /_| |_| |\  | | |  
\_|  |_/\___/\____/ \___/\_| \_/ \_/ 

v1.3
github.com/alpkeskin

Config File
├── Verify API
│   ├── True
│   └── Credits
│       └── 98
├── Social Scan
│   └── True
├── Leaked DB
│   └── True
├── Hunter API
│   └── True
├── Related Phone Numbers
│   └── True
├── Related Domains
│   └── True
├── Pastebin Dumps
│   └── True
├── Google Search
│   └── True
└── DNS Lookup
    └── True

MAIL > 

```

### :money_with_wings: My Bitcoin Wallet:
`3NFfd1QXUVFsZzfbwGJiAJdehtPB9D88tK`

#### :white_check_mark: Tested on:
- Kali Linux
- MacOS (without SocialScan module)
