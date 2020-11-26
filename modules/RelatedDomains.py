from bs4 import BeautifulSoup
import json, requests
from insides.bcolors import bcolors

def RelatedDomains(mail,_verbose=None):
	if _verbose != None:
		try:
			res = requests.get(f"https://www.threatcrowd.org/searchApi/v2/email/report/?email={mail}").json()
			domains = res.get('domains') or []
			if len(domains):
				print(f"{bcolors.BOLD}Related Domains:{bcolors.ENDC}")
				for domain in domains:
					print(domain)
			else:
				print(f"{bcolors.FAIL}No related domains found!{bcolors.ENDC}")
		except:
			print(f"{bcolors.FAIL}Threatcrowd Error!{bcolors.ENDC}")
