from bs4 import BeautifulSoup
import json
import requests
from insides.bcolors import bcolors
from insides.commonMails import commonMails


def Hunter(mail, hunterAPIKey, _verbose=None):
	if _verbose != None:
		try:
			domain = mail.split("@")[1]
			if (domain in commonMails):
				print(f"{bcolors.FAIL}Unacceptable domain :{bcolors.ENDC} {domain}")
			else:
				res = requests.get(f"https://api.hunter.io/v2/domain-search?domain={domain}&api_key={hunterAPIKey}").json()	
				if len(res['data']['emails']):
					print(f"{bcolors.BOLD}Related emails:{bcolors.ENDC}")
					for i in res['data']['emails'][:100]:
						print(i["value"])
				else:
					print(f"{bcolors.FAIL}No related mails found!{bcolors.ENDC}")
		except:
			print(f"{bcolors.FAIL}Hunter.io error!{bcolors.ENDC}")
