from bs4 import BeautifulSoup
import re, requests
from insides.bcolors import bcolors


def Googling(mail,_verbose=None):
	if _verbose != None:
		try:
			print(f"{bcolors.WARNING} -- Google Searching... [Pastebin & Throwbin]{bcolors.ENDC}")
			print(f"{bcolors.FAIL}!{bcolors.ENDC}"+"Google Search may not work properly.")
			print("")    
			x = mail.replace("@", "%40")
			searchurlP="https://s.sudonull.com/?q=site%3Apastebin.com+intext%3A%22"
			u = (searchurlP+x+"%22")
			response = requests.get(u)
			html = response.content
			soup=BeautifulSoup(html,"html.parser")
			rgx = str(soup)
			urls = re.findall('http[s]?://pastebin.com(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+<', rgx)
			for x in range(len(urls)): 
				p = urls[x].replace("<", "")
				print(f"{bcolors.OKGREEN}|-- {bcolors.ENDC}"+p)
			x = mail.replace("@", "%40")
			searchurlT="https://s.sudonull.com/?q=site%3Athrowbin.io+intext%3A%22"
			u = (searchurlT+x+"%22")
			response = requests.get(u)
			html = response.content
			soup=BeautifulSoup(html,"html.parser")
			rgx = str(soup)
			urls = re.findall('http[s]?://throwbin.io(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+<', rgx)
			for x in range(len(urls)): 
				t = urls[x].replace("<", "")
				print(f"{bcolors.OKGREEN}|-- {bcolors.ENDC}"+t)
		except:
			print(f"{bcolors.FAIL}Google Search error!{bcolors.ENDC}")
