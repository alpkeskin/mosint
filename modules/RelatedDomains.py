from bs4 import BeautifulSoup
import json, requests
from insides.bcolors import bcolors

def RelatedDomains(mail,_verbose=None):
	if _verbose != None:
		try:
			u = "https://www.threatcrowd.org/searchApi/v2/email/report/?email="+mail
			response = requests.get(u)
			html = response.content
			lp = json.loads(html)
			print(f"{bcolors.BOLD}Related Domains:{bcolors.ENDC}")
			for x in lp['domains']:
				print(x)
		except:
			print(f"{bcolors.FAIL}Domain not found!{bcolors.ENDC}")
