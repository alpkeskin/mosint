from bs4 import BeautifulSoup
import json,requests
from insides.bcolors import bcolors


def VerifyMail(verifyApi,mail,_verbose=None):
        if _verbose != None:
                
        	verifyurl = "https://app.verify-email.org/api/v1/"+verifyApi+"/verify/"
        	response1 = requests.get(verifyurl+str(mail))
        	html1 = response1.content
        	soup1=BeautifulSoup(html1,"html.parser")
        	strsoup1 = str(soup1)
        	data1 = json.loads(strsoup1)
        	print(f"[{bcolors.HEADER}#{bcolors.ENDC}]" + " Verification result : "+str(data1['status_description']))
