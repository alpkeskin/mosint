from bs4 import BeautifulSoup
import json, requests
from insides.bcolors import bcolors
from socialscan.util import Platforms, sync_execute_queries

def SocialScan(mail,_verbose=None):
	if _verbose != None:
		try:
			queries = [mail]
			platforms = [Platforms.GITHUB, Platforms.TWITTER, Platforms.INSTAGRAM, Platforms.PINTEREST, Platforms.SPOTIFY]
			results = sync_execute_queries(queries, platforms)
			for result in results:
				print(f"{bcolors.BOLD}{result.platform}:{bcolors.ENDC}{bcolors.WARNING} {result.message} (Success: {result.success}, Available: {result.available}){bcolors.ENDC}")
		except:
			print(f"{bcolors.FAIL}SocialScan Error!{bcolors.ENDC}")
