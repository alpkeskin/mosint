import requests
from verify_email import verify_email
from insides.bcolors import bcolors


def VerifyMail(verifyAPIKey, mail, _verbose=None):
	if _verbose != None:
		res = requests.get(f"https://app.verify-email.org/api/v1/{verifyAPIKey}/verify/{mail}").json()
		print(f"[{bcolors.HEADER}#{bcolors.ENDC}] Verify-Email.org result : {res.get('status_description')}")
		if verify_email(mail):
			print(f"[{bcolors.HEADER}#{bcolors.ENDC}] Built-in Mail Verify result : {bcolors.OKGREEN}OK{bcolors.ENDC}")
		else:
			print(f"[{bcolors.HEADER}#{bcolors.ENDC}] Built-in Mail Verify result : {bcolors.FAIL}FAILED{bcolors.ENDC}")