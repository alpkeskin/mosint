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
    parser.add_argument("-e", "--email", type=str, required=True, help="Email")
    return parser.parse_args()


def main():
    args = parse_args()
    mail = args.email
    with open("./keys.json", "r") as configFile:
        conf = json.loads(configFile.read())
    if conf[0]["BreachDirectory.org API Key"] == "":
        print(
            f"{bcolors.FAIL}[-] Enter the API key in the keys.json file to use this feature!{bcolors.ENDC}"
        )
        exit()
    url = "https://breachdirectory.p.rapidapi.com/"
    querystring = {"func": "auto", "term": mail}
    headers = {
        "x-rapidapi-host": "breachdirectory.p.rapidapi.com",
        "x-rapidapi-key": conf[0]["BreachDirectory.org API Key"],
    }
    response = requests.request("GET", url, headers=headers, params=querystring)
    data = response.json()
    try:
        if data["success"]:
            count = data["found"]
            x = 0
            passwords = []
            sha1 = []
            breached_sites = []
            while x < count:
                if data["result"][x]["has_password"]:
                    passwords.append(data["result"][x]["password"])
                    sha1.append(data["result"][x]["sha1"])
                else:
                    passwords.append("No data!")
                    sha1.append("No data!")
                breached_sites.append(data["result"][x]["sources"])
                x = x + 1

            x = 0
            rng = len(passwords)
            while x < rng:
                print(
                    f"Source: {bcolors.BOLD}" + breached_sites[x][0] + f"{bcolors.ENDC}"
                )
                print(
                    f"|-- Password: {bcolors.OKGREEN}"
                    + passwords[x]
                    + f"{bcolors.ENDC}"
                )
                print(f"|-- SHA1: {bcolors.OKGREEN}" + sha1[x] + f"{bcolors.ENDC}")
                x = x + 1
        else:
            print(f"{bcolors.FAIL}[-] No data!{bcolors.ENDC}")
    except:
        print(f"{bcolors.FAIL}[-] Error!{bcolors.ENDC}")


main()
