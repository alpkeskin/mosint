from bs4 import BeautifulSoup
import json, requests
from insides.bcolors import bcolors

def Hunter(mail,hunterApi,_verbose=None):
	if _verbose != None:
		try:

			dmnlist = ["gmail.com","outlook.com","hotmail.com","yahoo.com","hotmail.co.uk","icloud.com"]
			at = "@"
			domain = (mail[mail.index(at) + len(at):])
			if (domain in dmnlist):
				print(f"{bcolors.FAIL}Unacceptable domain : {bcolors.ENDC}"+domain)
			else:
				print(f"{bcolors.BOLD}Related emails:{bcolors.ENDC}")
				u = "https://api.hunter.io/v2/domain-search?domain="+domain+"&api_key="+hunterApi
				response = requests.get(u)
				html = response.content
				lp = json.loads(html)
				for i in range(0,99):
					print(lp['data']['emails'][i]['value'])
		except:
			pass
