from anytree import Node, RenderTree
from bs4 import BeautifulSoup
import json, requests
from insides.bcolors import bcolors

# TODO: Clean boolean in string.

def ConfigTree(verifyApi,socialscan,leakeddb,breachedsites,hunterApi,checkPDF,dbdata,tcrwd,pastebindumps,googlesearch,dns, _verbose=None):
	if _verbose != None:
		try:
			fileshow = Node(f"{bcolors.BOLD}Config File{bcolors.ENDC} [Modules]")
			vrfctnsrvc1 = Node("Verify API", parent=fileshow)
			if (verifyApi != ""):
				creditsurl = "https://app.verify-email.org/api/v1/"+verifyApi+"/credits"
				response = requests.get(creditsurl)
				html = response.content
				soup=BeautifulSoup(html,"html.parser")
				strsoup = str(soup)
				data = json.loads(strsoup)
				str(data['credits'])
				vrfctnsrvc11 = Node('\x1b[6;30;42m'+'True'+ '\x1b[0m', parent=vrfctnsrvc1)
				vrfctnsrvc12 = Node("Credits", parent=vrfctnsrvc1)
				vrfctnsrvc121 = Node('\x1b[6;30;42m'+str(data['credits'])+ '\x1b[0m', parent=vrfctnsrvc12)
			else:
				vrfctnsrvc11 = Node('\x1b[1;31;40m'+'False'+ '\x1b[0m', parent=vrfctnsrvc1)
			sclscn1 = Node("Social Scan", parent=fileshow)
			if (socialscan == "True" or socialscan == "true"):
				sclscn11 = Node('\x1b[6;30;42m'+socialscan+ '\x1b[0m', parent=sclscn1)
			else:
				sclscn11 = Node('\x1b[1;31;40m'+'False'+ '\x1b[0m', parent=sclscn1)
			lkeddb1 = Node("Leaked DB", parent=fileshow)
			if (leakeddb == "True" or leakeddb == "true"):
				lkeddb11 = Node('\x1b[6;30;42m'+leakeddb+ '\x1b[0m', parent=lkeddb1)
			else:
				lkeddb11 = Node('\x1b[1;31;40m'+'False'+ '\x1b[0m', parent=lkeddb1)
			breachedS = Node("Breached Sites", parent=fileshow)
			if (breachedsites != ""):
				bs = Node('\x1b[6;30;42m'+'True'+ '\x1b[0m', parent=breachedS)
			else:
				bs = Node('\x1b[1;31;40m'+'False'+ '\x1b[0m', parent=breachedS)
			hntr1 = Node("Hunter API", parent=fileshow)
			if (hunterApi != ""):
				hntr11 = Node('\x1b[6;30;42m'+'True'+ '\x1b[0m', parent=hntr1)
			else:
				hntr11 = Node('\x1b[1;31;40m'+'False'+ '\x1b[0m', parent=hntr1)
			pdf = Node("PDF Check", parent=fileshow)
			if (checkPDF == "True" or checkPDF == "true"):
				pdf1 = Node('\x1b[6;30;42m'+'True'+ '\x1b[0m', parent=pdf)
			else:
				pdf1 = Node('\x1b[1;31;40m'+'False'+ '\x1b[0m', parent=pdf)
			dbdt1 = Node("Related Phone Numbers", parent=fileshow)
			if (dbdata == "True" or dbdata == "true"):
				dbdt11 = Node('\x1b[6;30;42m'+dbdata+ '\x1b[0m', parent=dbdt1)
			else:
				dbdt11 = Node('\x1b[1;31;40m'+'False'+ '\x1b[0m', parent=dbdt1)
			thrcwd1 = Node("Related Domains", parent=fileshow)
			if (tcrwd == "True" or tcrwd == "true"):
				thrcwd11 = Node('\x1b[6;30;42m'+tcrwd+ '\x1b[0m', parent=thrcwd1)
			else:
				thrcwd11 = Node('\x1b[1;31;40m'+'False'+ '\x1b[0m', parent=thrcwd1)
			pstbn1 = Node("Pastebin Dumps", parent=fileshow)
			if (pastebindumps == "True" or pastebindumps == "true"):
				pstbn11 = Node('\x1b[6;30;42m'+pastebindumps+ '\x1b[0m', parent=pstbn1)
			else:
				pstbn11 = Node('\x1b[1;31;40m'+'False'+ '\x1b[0m', parent=pstbn1)
			goo1 = Node("Google Search", parent=fileshow)
			if (googlesearch == "True" or googlesearch == "true"):
				goo11 = Node('\x1b[6;30;42m'+googlesearch+ '\x1b[0m', parent=goo1)
			else:
				goo11 = Node('\x1b[1;31;40m'+'False'+ '\x1b[0m', parent=goo1)
			dns1 = Node("DNS Lookup", parent=fileshow)
			if (dns == "True" or dns == "true"):
				dns11 = Node('\x1b[6;30;42m'+dns+ '\x1b[0m', parent=dns1)
			else:
				dns11 = Node('\x1b[1;31;40m'+'False'+ '\x1b[0m', parent=dns1)

			for pre, fill, node in RenderTree(fileshow):
				print("%s%s" % (pre, node.name))
			print("")
		except:
			print("Tree Error!")
