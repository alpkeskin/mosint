import requests
from prettytable import PrettyTable
from insides.bcolors import bcolors

def DNS(mail,_verbose=None):
	if _verbose != None:
		try:
			print(f"{bcolors.WARNING} -- DNS Records [ Cloudflare ]{bcolors.ENDC}")
			domain = mail.split("@")[1]
			dnstable = PrettyTable([f"{bcolors.WARNING}Record Type{bcolors.ENDC}","Answer"])
			recordTypes = ["NS", "A", "AAAA", "TXT", "MX"]
			for recordType in recordTypes:
				params = (
					('name', domain),
					('type', recordType),
					('cd', 'false'),
				)

				r = requests.get('https://cloudflare-dns.com/dns-query', headers={ 'accept': 'application/dns-json' }, params=params).json().get("Answer") or []

				if recordType == "NS" and r == []:
					print(f"{bcolors.FAIL}NS can not found! [ DNS Lookup ]{bcolors.ENDC}")
					return

				for record in r:
					value = record.get('data')
					dnstable.add_row([recordType, value])
			print(dnstable)
		except:
			print(f"{bcolors.FAIL}Cloudflare DoH error!{bcolors.ENDC}")