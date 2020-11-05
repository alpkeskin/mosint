# https://github.com/alpkeskin
import json
from bs4 import BeautifulSoup
from modules.bcolors import bcolors

configfile = open('config.json', "r")
conf = json.loads(configfile.read())
for i in conf:
    verifyApi = (i['verify-email.org API Key'])
    socialscan = (i['Social Scan'])
    leakeddb = (i['Leaked DB'])
    hunterApi = (i['hunter.io API Key'])
    dbdata = (i['Related Phone Numbers'])
    tcrwd = (i['Related Domains'])
    pastebindumps = (i['Pastebin Dumps'])
    googlesearch = (i['Google Search'])
    dns = (i['DNS Lookup'])

from modules.Banner import Banner
Banner()

from modules.ConfigTree import ConfigTree
ConfigTree(verifyApi,socialscan,leakeddb,hunterApi,dbdata,tcrwd,pastebindumps,googlesearch,dns,_verbose=True)

print("")
while True:
    mail=input(f"{bcolors.OKBLUE}MAIL > {bcolors.ENDC}")
    if (mail == "q"):
        print("Thank you for using "+f"{bcolors.BOLD}MOSINT{bcolors.ENDC}.")
        break
    elif (mail.find("@") == -1 and mail.find(".") == -1):
        print(f"{bcolors.FAIL}Email format is wrong!{bcolors.ENDC}")
        break

    if (verifyApi != ""): 
        from modules.VerifyMail import VerifyMail
        VerifyMail(verifyApi,mail,_verbose=True)

    if (socialscan == "True" or socialscan == "T" or socialscan == "true"):
        from modules.SocialScan import SocialScan
        SocialScan(mail,_verbose=True)

    if (leakeddb == "True" or leakeddb == "T" or leakeddb == "true"):
        from modules.LeakedDB import LeakedDB
        LeakedDB(mail,_verbose=True)

    if (hunterApi != ""):
        from modules.Hunter import Hunter
        Hunter(mail,hunterApi,_verbose=True)

    if (dbdata == "True" or dbdata == "T" or dbdata == "true"):
        from modules.RelatedNumbers import RelatedNumbers
        RelatedNumbers(mail,_verbose=True)

    if (tcrwd == "True" or tcrwd == "T" or tcrwd == "true"):
        from modules.RelatedDomains import RelatedDomains
        RelatedDomains(mail,_verbose=True)

    if (pastebindumps == "True" or pastebindumps == "T" or pastebindumps == "true"):
        from modules.Psbdmp import Psbdmp
        Psbdmp(mail,_verbose=True)

    if (googlesearch == "True" or googlesearch == "T" or googlesearch == "true"):
        from modules.Googling import Googling
        Googling(mail,_verbose=True)

    if (dns == "True" or dns == "T" or dns == "true"):
        from modules.DNS import DNS
        DNS(mail,_verbose=True)
      
