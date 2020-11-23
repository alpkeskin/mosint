from bs4 import BeautifulSoup
import json, requests
from insides.bcolors import bcolors

def Psbdmp(mail,_verbose=None):
	if _verbose != None:
		try:
			headers = { 
			"Accept": "application/json" 
			}
			print(f"{bcolors.WARNING} -- Scanning Pastebin Dumps...{bcolors.ENDC}")
			print("")
			psbdmpurl = "https://psbdmp.ws/api/search/"
			u = (psbdmpurl+mail)
			response = requests.get(u,headers=headers)
			html = response.content
			lp = json.loads(html)
			for i in lp['data']:
				print(f"{bcolors.OKGREEN}|-- {bcolors.ENDC}"+"https://pastebin.com/"+i['id'])
		except:
			print(f"{bcolors.FAIL}Dump not found!{bcolors.ENDC}")
