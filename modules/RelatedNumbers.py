from bs4 import BeautifulSoup
import requests, re
from insides.bcolors import bcolors

def RelatedNumbers(mail, _verbose=None):
	if _verbose != None:
		try:
			html = requests.get(f"https://domainbigdata.com/email/{ mail }").content
			soup=BeautifulSoup(html,"html.parser")
			rgx = str(soup)
			phones = re.findall('<td colspan="2">\+.*.</td>', rgx)
			for phone in phones:
				number = phone.replace("<td colspan="+'"'+'2'+'"'+'>', "").replace("</td>","")
				print(number)
			if len(phones) == 0:
				print(f"{bcolors.FAIL}No phone numbers found!{bcolors.ENDC}")
		except:
			print(f"{bcolors.FAIL}Phone Number Data Error!{bcolors.ENDC}")
