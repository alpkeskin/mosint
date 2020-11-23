import requests
import json
from insides.bcolors import bcolors

def BreachedSites(mail,breachedsites,_verbose=None):
	if _verbose != None:
		try:
			url = "https://leak-lookup.com/api/search"
			payload = {"key": breachedsites, "type": "email_address", "query": mail}
			req = requests.post(url, data=payload, timeout=30)
			email_response = req.json()
			data = str(email_response['message'])
			data_parse = data.split(',')
			for leaks in data_parse:
				ll = str(leaks)
				print(f"[{bcolors.WARNING}!{bcolors.ENDC}] "+ll[2:-5])
		except:
			print(f"{bcolors.FAIL}Leak-lookup.com error!{bcolors.ENDC}")