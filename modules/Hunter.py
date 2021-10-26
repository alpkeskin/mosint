import argparse
import json
import requests


class bcolors:
    OKGREEN = "\033[92m"
    FAIL = "\033[91m"
    BOLD = "\033[1m"
    ENDC = "\033[0m"


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument("-d", "--domain", type=str, required=True, help="Domain")
    return parser.parse_args()


def main():
    args = parse_args()
    domain = args.domain
    with open("./keys.json", "r") as configFile:
        conf = json.loads(configFile.read())
    key = conf[0]["hunter.io API Key"]
    print("From hunter.io:")
    if conf[0]["hunter.io API Key"] == "":
        print(
            f"{bcolors.FAIL}[-] Enter the API key in the keys.json file to use this feature!{bcolors.ENDC}"
        )
        exit()
    res = requests.get(
        f"https://api.hunter.io/v2/domain-search?domain={domain}&api_key={key}"
    ).json()
    x = 0
    try:
        if len(res["data"]["emails"]) == 0:
            print(f"{bcolors.FAIL}[-] No data!{bcolors.ENDC}")
        while x < len(res["data"]["emails"]):
            print(
                f"{bcolors.OKGREEN}[+] "
                + res["data"]["emails"][x]["value"]
                + f"{bcolors.ENDC}"
            )
            x = x + 1
    except:
        print(f"{bcolors.FAIL}[-] Error!{bcolors.ENDC}")


main()
