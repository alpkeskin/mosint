from bs4 import BeautifulSoup
import json, requests
from insides.bcolors import bcolors

def Psbdmp(mail,_verbose=None):
	if _verbose != None:
		try:
			print(f"{bcolors.WARNING} -- Scanning Pastebin Dumps...{bcolors.ENDC}\n")
			res = requests.get(f"https://psbdmp.ws/api/search/{mail}",headers={ "Accept": "application/json" }).json().get('data') or []
			if len(res):
				for i in res:
					print(f"{bcolors.OKGREEN}|-- {bcolors.ENDC}"+"https://pastebin.com/"+i['id'])
			else:
				print(f"{bcolors.FAIL}No psbdump records found!{bcolors.ENDC}")
		except:
			print(f"{bcolors.FAIL}Psbdump Error!{bcolors.ENDC}")
