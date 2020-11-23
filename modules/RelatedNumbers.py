from bs4 import BeautifulSoup
import requests, re
from insides.bcolors import bcolors

def RelatedNumbers(mail,_verbose=None):
	if _verbose != None:
		try:
			dnsurl = ("https://domainbigdata.com/email/"+mail)
			response = requests.get(dnsurl)
			html = response.content
			soup=BeautifulSoup(html,"html.parser")
			rgx = str(soup)
			urls = re.findall('<td colspan="2">\+.*.</td>', rgx)
			for x in range(len(urls)): 
				n = urls[x].replace("<td colspan="+'"'+'2'+'"'+'>', "")
				print(n)
		except:
			print(f"{bcolors.FAIL}Phone Number Data Error!{bcolors.ENDC}")
