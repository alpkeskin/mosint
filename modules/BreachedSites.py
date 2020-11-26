import requests
from insides.bcolors import bcolors

def BreachedSites(mail,breachedsites,_verbose=None):
	if _verbose != None:
		try:
			url = "https://leak-lookup.com/api/search"
			payload = {"key": breachedsites, "type": "email_address", "query": mail}
			res = requests.post(url, data=payload, timeout=30).json()
			if res['error'] == 'false' and isinstance(res['message'], dict):
				for i in res['message'].keys():
					print(f"[{bcolors.WARNING}!{bcolors.ENDC}] {i}")
			else:
				print(f"{bcolors.FAIL}Leak-lookup.com API error:{bcolors.ENDC} {res['message']}")
		except:
			print(f"{bcolors.FAIL}Leak-lookup.com error!{bcolors.ENDC}")