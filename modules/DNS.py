from bs4 import BeautifulSoup
import requests
from prettytable import PrettyTable
from modules.bcolors import bcolors

def DNS(mail,_verbose=None):
	if _verbose != None:
		try:
			at = "@"
			domain = (mail[mail.index(at) + len(at):])
			dnsurl = ("https://api.hackertarget.com/dnslookup/?q="+domain)
			dnstable = PrettyTable([f"{bcolors.WARNING}DNS LOOKUP{bcolors.ENDC}"])
			response = requests.get(dnsurl)
			html = response.content
			soup=BeautifulSoup(html,"html.parser")
			dnstable.add_row([soup])
			print(dnstable)
		except:
			print("Service Error! { DNS Lookup}")