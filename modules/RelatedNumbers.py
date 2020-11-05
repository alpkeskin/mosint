from bs4 import BeautifulSoup
import requests, re
from modules.bcolors import bcolors

def RelatedNumbers(mail,_verbose=None):
	if _verbose != None:
		try:
			dnsurl = ("https://domainbigdata.com/email/"+mail)
			response = requests.get(dnsurl)
			html = response.content
			soup=BeautifulSoup(html,"html.parser")
			rgx = str(soup)
			urls = re.findall('<td colspan="2">\+.*.</td>', rgx)
			print(f"{bcolors.BOLD}Related Phone Numbers:{bcolors.ENDC}")
			for x in range(len(urls)): 
				n = urls[x].replace("<td colspan="+'"'+'2'+'"'+'>', "")
				print(n)
			print("")    
			print("------------------------")  
			print("")
		except:
			print("Phone Number Data Error!")