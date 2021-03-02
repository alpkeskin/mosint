import requests
from insides.bcolors import bcolors
from prettytable import PrettyTable

def LeakedDB(mail,_verbose=None):
	if _verbose != None:
		try:
			res = requests.get(f"https://www.scylladb.com/search_gcse/?q=email:{mail}&size=50&start=0", headers={'Accept': 'application/json'}).json()
			table = PrettyTable(["Domain","Email",f"{bcolors.FAIL}Password{bcolors.ENDC}"])
			if len(res):
				for s in res:
					table.add_row([s["fields"]["domain"],s["fields"]["email"],s["fields"].get("password") or s["fields"].get("passhash") or "No Pass Data"])
				print(table)
			else:
				print(f"{bcolors.FAIL}No leaked accounts found!{bcolors.ENDC}")
		except Exception as e:
			print(e)
			print(f"{bcolors.FAIL}Leaked DB Connection Error!{bcolors.ENDC}")
