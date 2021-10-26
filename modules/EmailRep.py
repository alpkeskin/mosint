from emailrep import EmailRep
import argparse
import json


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
		if conf[0]["EmailRep.io API Key"] == "":
			print(f"{bcolors.FAIL}[-] Enter the API key in the keys.json file to use this feature!{bcolors.ENDC}")
			exit()
	emailrep = EmailRep(conf[0]["EmailRep.io API Key"])
	result = emailrep.query(mail)
	print(f"|-- blacklisted: {bcolors.BOLD}" + str(result["details"]["blacklisted"]) + f"{bcolors.ENDC}")
	print(f"|-- malicious_activity: {bcolors.BOLD}" + str(result["details"]["malicious_activity"]) + f"{bcolors.ENDC}")
	if result["details"]["credentials_leaked"]:
		print(f"|-- credentials_leaked: {bcolors.OKGREEN}" + str(result["details"]["credentials_leaked"]) + f"{bcolors.ENDC}")
	else:
		print(f"|-- credentials_leaked: {bcolors.BOLD}" + str(result["details"]["credentials_leaked"]) + f"{bcolors.ENDC}")
	if result["details"]["data_breach"]:
		print(f"|-- data_breach: {bcolors.OKGREEN}" + str(result["details"]["data_breach"]) + f"{bcolors.ENDC}")
	else:
		print(f"|-- data_breach: {bcolors.BOLD}" + str(result["details"]["data_breach"]) + f"{bcolors.ENDC}")
	print(f"|-- domain_exists: {bcolors.BOLD}" + str(result["details"]["domain_exists"]) + f"{bcolors.ENDC}")
	print(f"|-- new_domain: {bcolors.BOLD}" + str(result["details"]["new_domain"]) + f"{bcolors.ENDC}")
	print(f"|-- free_provider: {bcolors.BOLD}" + str(result["details"]["free_provider"]) + f"{bcolors.ENDC}")
	print(f"|-- valid_mx: {bcolors.BOLD}" + str(result["details"]["valid_mx"]) + f"{bcolors.ENDC}")
	print(f"|-- spoofable: {bcolors.BOLD}" + str(result["details"]["spoofable"]) + f"{bcolors.ENDC}")
	print(f"|-- spam: {bcolors.BOLD}" + str(result["details"]["spam"]) + f"{bcolors.ENDC}")

main()