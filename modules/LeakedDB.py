from bs4 import BeautifulSoup
import json, requests
from modules.bcolors import bcolors
from prettytable import PrettyTable

def LeakedDB(mail,_verbose=None):
	if _verbose != None:
		try:
			headers = { 
			"Accept": "application/json" 
			}
			leakedpassurl = "https://scylla.sh/search?q=email:"
			u = (leakedpassurl+mail)
			response = requests.get(u,timeout=5,headers=headers)
			html = response.content
			lp = json.loads(html)
			table = PrettyTable(["Domain","Email",f"{bcolors.FAIL}Password{bcolors.ENDC}"])
			for s in range(len(lp)):
				table.add_row([lp[s]["fields"]["domain"],lp[s]["fields"]["email"],lp[s]["fields"]["password"]])
			print(table)
			print("")
			print("------------------------")  
			print("")
		except:
			print("Leaked DB Connection Error!")