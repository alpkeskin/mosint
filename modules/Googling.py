from googlesearch import search
from insides.bcolors import bcolors


def Googling(mail,_verbose=None):
	if _verbose != None:
		try:
			print(f"{bcolors.WARNING} -- Google Searching... [Pastebin & Throwbin]{bcolors.ENDC}")
			print(f"{bcolors.FAIL}!{bcolors.ENDC}"+"Google Search may not work properly.\n")
			urls = search(f'site:throwbin.io intext:"{mail}"') + search(f'site:pastebin.com intext:"{mail}"')
			if len(urls):
				for x in urls: 
					print(f"{bcolors.OKGREEN}|-- {bcolors.ENDC}"+x)
			else:
				print(f"{bcolors.FAIL}No Google Search result found!{bcolors.ENDC}")
		except:
			print(f"{bcolors.FAIL}Google Search error!{bcolors.ENDC}")
