# https://github.com/alpkeskin
import json, re
from insides.bcolors import bcolors
from insides.Header import Header

with open('config.json', "r") as configFile:
    conf = json.loads(configFile.read())

for i in conf:
    verifyApi = (i['verify-email.org API Key'])
    socialscan = (i['Social Scan'])
    leakeddb = (i['Leaked DB'])
    breachedsites = (i['Breached Sites[leak-lookup.com API Key]'])
    hunterApi = (i['hunter.io API Key'])
    dbdata = (i['Related Phone Numbers'])
    tcrwd = (i['Related Domains'])
    pastebindumps = (i['Pastebin Dumps'])
    googlesearch = (i['Google Search'])
    dns = (i['DNS Lookup'])

from insides.Banner import Banner
Banner()

from modules.ConfigTree import ConfigTree
ConfigTree(verifyApi,socialscan,leakeddb,breachedsites,hunterApi,dbdata,tcrwd,pastebindumps,googlesearch,dns,_verbose=True)

print("")

# TODO: Clean boolean in string.

EMAIL_REGEX = r'(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)'

while True:
    mail=input(f"{bcolors.OKBLUE}MAIL > {bcolors.ENDC}")

    if (mail == "q" or mail == "Q" or mail == "exit"):
        print("Thank you for using "+f"{bcolors.BOLD}MOSINT{bcolors.ENDC}.")
        break
    elif not re.match(EMAIL_REGEX, mail):
        print(f"{bcolors.FAIL}Email format is wrong!{bcolors.ENDC}")
        continue

    if (verifyApi != ""): 
        from modules.VerifyMail import VerifyMail
        title = "VERIFICATION SERVICE"
        Header(title)
        VerifyMail(verifyApi,mail,_verbose=True)

    if (socialscan == "True" or socialscan == "T" or socialscan == "true"):
        from modules.SocialScan import SocialScan
        title = "SOCIAL SCAN"
        Header(title)
        SocialScan(mail,_verbose=True)

    if (leakeddb == "True" or leakeddb == "T" or leakeddb == "true"):
        from modules.LeakedDB import LeakedDB
        title = "LEAKED DB [Password dumps]"
        Header(title)
        LeakedDB(mail,_verbose=True)

    if (breachedsites != ""):
        from modules.BreachedSites import BreachedSites
        title = "BREACHED SITES"
        Header(title)
        BreachedSites(mail,breachedsites,_verbose=True)

    if (hunterApi != ""):
        from modules.Hunter import Hunter
        title = "RELATED EMAILS"
        Header(title)
        Hunter(mail,hunterApi,_verbose=True)

    if (dbdata == "True" or dbdata == "T" or dbdata == "true"):
        from modules.RelatedNumbers import RelatedNumbers
        title = "RELATED PHONE NUMBERS"
        Header(title)
        RelatedNumbers(mail,_verbose=True)

    if (tcrwd == "True" or tcrwd == "T" or tcrwd == "true"):
        from modules.RelatedDomains import RelatedDomains
        title = "RELATED DOMAINS"
        Header(title)
        RelatedDomains(mail,_verbose=True)

    if (pastebindumps == "True" or pastebindumps == "T" or pastebindumps == "true"):
        from modules.Psbdmp import Psbdmp
        title = "PASTEBIN DUMPS"
        Header(title)
        Psbdmp(mail,_verbose=True)

    if (googlesearch == "True" or googlesearch == "T" or googlesearch == "true"):
        from modules.Googling import Googling
        title = "GOOGLING"
        Header(title)
        Googling(mail,_verbose=True)

    if (dns == "True" or dns == "T" or dns == "true"):
        from modules.DNS import DNS
        title = "DNS LOOKUP"
        Header(title)
        DNS(mail,_verbose=True)
      
