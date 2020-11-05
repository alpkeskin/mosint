from bs4 import BeautifulSoup
import json, requests
from modules.bcolors import bcolors

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
			print("")   
			print("------------------------")  
			print("")
		except:
			print("Dump not found!")