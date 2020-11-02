# https://github.com/alpkeskin
import requests
import re
from bs4 import BeautifulSoup
import json
import os
from prettytable import PrettyTable

class bcolors:
    HEADER = '\033[95m'
    OKBLUE = '\033[94m'
    OKGREEN = '\033[92m'
    WARNING = '\033[93m'
    FAIL = '\033[91m'
    ENDC = '\033[0m'
    BOLD = '\033[1m'
    UNDERLINE = '\033[4m'

configfile = open('config.json', "r")
conf = json.loads(configfile.read())
for i in conf:
    checkconnect = (i['Check Connections'])
    verifyApi = (i['verify-email.org API Key'])
    socialscan = (i['Social Scan'])
    leakeddb = (i['Leaked DB'])
    hunterApi = (i['shodan.io API Key'])
    dbdata = (i['Related Phone Numbers'])
    tcrwd = (i['Related Domains'])
    pastebindumps = (i['Pastebin Dumps'])
    googlesearch = (i['Google Search'])
    dns = (i['DNS Lookup'])

print(f'''{bcolors.BOLD}
___  ________ _____ _____ _   _ _____ 
|  \/  |  _  /  ___|_   _| \ | |_   _|
| .  . | | | \ `--.  | | |  \| | | |  
| |\/| | | | |`--. \ | | | . ` | | |  
| |  | \ \_/ /\__/ /_| |_| |\  | | |  
\_|  |_/\___/\____/ \___/\_| \_/ \_/ 
{bcolors.ENDC}
v{bcolors.BOLD}1.3{bcolors.ENDC}
github.com/{bcolors.BOLD}alpkeskin{bcolors.ENDC}
''')
try:
    from anytree import Node, RenderTree
    fileshow = Node(f"{bcolors.BOLD}Config File{bcolors.ENDC}")
    chckcnntns1 = Node("Check Connections", parent=fileshow)
    if (checkconnect == "True" or checkconnect == "true"):
        chckcnntns11 = Node('\x1b[6;30;42m'+checkconnect+ '\x1b[0m', parent=chckcnntns1)
    else:
        chckcnntns11 = Node('\x1b[1;31;40m'+'False'+ '\x1b[0m', parent=chckcnntns1)
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
    hntr1 = Node("Hunter API", parent=fileshow)
    if (hunterApi != ""):
        hntr11 = Node('\x1b[6;30;42m'+'True'+ '\x1b[0m', parent=hntr1)
    else:
        hntr11 = Node('\x1b[1;31;40m'+'False'+ '\x1b[0m', parent=hntr1)
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
    print()

def connection(url='https://www.google.com/' , timeout=5):
    try:
        req = requests.get(url, timeout=timeout)
        req.raise_for_status()
        print(f"{bcolors.OKGREEN}[+] You're connected to internet.{bcolors.ENDC}")
        return True
    except requests.HTTPError as e:
        print("[-] Checking internet connection failed, status code {0}.".format(e.response.status_code))
    except requests.ConnectionError:
        print(f"{bcolors.FAIL}[-] No internet connection available.{bcolors.ENDC}")
    return False


def verifyconnect(url='https://verify-email.org/' , timeout=5):
    try:
        req = requests.get(url, timeout=timeout)
        req.raise_for_status()
        print(f"{bcolors.OKGREEN}[+] You're connected to verification service.{bcolors.ENDC}")
        return True
    except requests.HTTPError as e:
        print("[-] Checking internet connection failed, status code {0}.".format(e.response.status_code))
    except requests.ConnectionError:
        print(f"{bcolors.FAIL}[-] No verification service available.{bcolors.ENDC}")
    return False


def pwnconnect(url='https://psbdmp.ws/' , timeout=5):
    try:
        req = requests.get(url, timeout=timeout)
        req.raise_for_status()
        print(f"{bcolors.OKGREEN}[+] You're connected to leaked dumps.{bcolors.ENDC}")
        return True
    except requests.HTTPError as e:
        print("[-] Checking internet connection failed, status code {0}.".format(e.response.status_code))
    except requests.ConnectionError:
        print(f"{bcolors.FAIL}[-] No leaked dumps available.{bcolors.ENDC}")
    return False


if (checkconnect == "True" or checkconnect == "T" or checkconnect == "true"):
    connection()
    verifyconnect()
    pwnconnect()
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
        verifyurl = "https://app.verify-email.org/api/v1/"+verifyApi+"/verify/"
        response1 = requests.get(verifyurl+str(mail))
        html1 = response1.content
        soup1=BeautifulSoup(html1,"html.parser")
        strsoup1 = str(soup1)
        data1 = json.loads(strsoup1)
        print(f"{bcolors.HEADER}[#]{bcolors.ENDC}" + " Verification result : "+str(data1['status_description']))
        print("")
        print("------------------------")
        print("")
    try:
        if (socialscan == "True" or socialscan == "T" or socialscan == "true"):
            from socialscan.util import Platforms, sync_execute_queries
            queries = [mail]
            platforms = [Platforms.GITHUB, Platforms.TWITTER, Platforms.INSTAGRAM, Platforms.PINTEREST, Platforms.SPOTIFY]
            results = sync_execute_queries(queries, platforms)
            for result in results:
                print(f"{bcolors.BOLD}{result.platform}:{bcolors.ENDC}{bcolors.WARNING} {result.message} (Success: {result.success}, Available: {result.available}){bcolors.ENDC}")
            print("")    
            print("------------------------")  
            print("")
    except:
        print("Social Scan Error!")
    try:
        if (leakeddb == "True" or leakeddb == "T" or leakeddb == "true"):
            headers = { 
            "Accept": "application/json" 
            }
            leakedpassurl = "https://scylla.sh/search?q=email:"
            u = (leakedpassurl+mail)
            response = requests.get(u,headers=headers)
            html = response.content
            lp = json.loads(html)
            table = PrettyTable(["Domain","Email",f"{bcolors.FAIL}Password{bcolors.ENDC}"])
            for s in range(len(lp)):
                table.add_row([lp[s]["fields"]["domain"],lp[s]["fields"]["email"],lp[s]["fields"]["password"]])
            print(table)
            print("")    
            print("------------------------")  
            print("")
    except:
        print("DB Connection Error!")
        print("")

    try:
        dmnlist = ["gmail.com","outlook.com","hotmail.com","yahoo.com","hotmail.co.uk","icloud.com"]
        at = "@"
        domain = (mail[mail.index(at) + len(at):])
        if (hunterApi != "" and domain in dmnlist):
            print(f"{bcolors.FAIL}Unacceptable domain : {bcolors.ENDC}"+domain)
            print("")    
            print("------------------------")  
            print("")
        elif (hunterApi != ""):
            print(f"{bcolors.BOLD}Related emails:{bcolors.ENDC}")
            u = "https://api.hunter.io/v2/domain-search?domain="+domain+"&api_key="+hunterApi
            response = requests.get(u)
            html = response.content
            lp = json.loads(html)
            for i in range(0,99):
                print(lp['data']['emails'][i]['value'])
    except:
        print("")    
        print("------------------------")  
        print("")

    try:
        if (dbdata == "True" or dbdata == "T" or dbdata == "true"):
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

    try:
        if (tcrwd == "True" or tcrwd == "T" or tcrwd == "true"):
            u = "https://www.threatcrowd.org/searchApi/v2/email/report/?email="+mail
            response = requests.get(u)
            html = response.content
            lp = json.loads(html)
            print(f"{bcolors.BOLD}Related Domains:{bcolors.ENDC}")
            for x in lp['domains']:
                print(x)
            print("")
            print("------------------------")  
            print("")
    except:
        print("Domain not found!")
    try:
        if (pastebindumps == "True" or pastebindumps == "T" or pastebindumps == "true"):
            print(f"{bcolors.WARNING} -- Scanning Pastebin Dumps...{bcolors.ENDC}")
            print("")
            psbdmpurl = "https://psbdmp.ws/api/search/"
            u = (psbdmpurl+mail)
            response = requests.get(u,headers=headers)
            html = response.content
            lp = json.loads(html)
            for i in lp['data']:
                print(f"{bcolors.OKGREEN}|-- {bcolors.ENDC}"+"https://pastebin.com/"+i['id'])
            print("")    
            print("------------------------")  
            print("")
    except:
        print("Dump not found!")
    try:
        if (googlesearch == "True" or googlesearch == "T" or googlesearch == "true"):
            print(f"{bcolors.WARNING} -- Google Searching... [Pastebin & Throwbin]{bcolors.ENDC}")
            print(f"{bcolors.FAIL}!{bcolors.ENDC}"+"Google Search may not work properly.")
            print("")    
            x = mail.replace("@", "%40")
            searchurlP="https://s.sudonull.com/?q=site%3Apastebin.com+intext%3A%22"
            u = (searchurlP+x+"%22")
            response = requests.get(u)
            html = response.content
            soup=BeautifulSoup(html,"html.parser")
            rgx = str(soup)
            urls = re.findall('http[s]?://pastebin.com(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+<', rgx)
    
            for x in range(len(urls)): 
                p = urls[x].replace("<", "")
                print(f"{bcolors.OKGREEN}|-- {bcolors.ENDC}"+p)

            x = mail.replace("@", "%40")
            searchurlT="https://s.sudonull.com/?q=site%3Athrowbin.io+intext%3A%22"
            u = (searchurlT+x+"%22")
            response = requests.get(u)
            html = response.content
            soup=BeautifulSoup(html,"html.parser")
            rgx = str(soup)
            urls = re.findall('http[s]?://throwbin.io(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+<', rgx)
    
            for x in range(len(urls)): 
                t = urls[x].replace("<", "")
                print(f"{bcolors.OKGREEN}|-- {bcolors.ENDC}"+t)
            print("")    
            print("------------------------")  
            print("")
    except:
        print("Google Search error!")  
    try:
        if (dns == "True" or dns == "T" or dns == "true"):
            at = "@"
            domain = (mail[mail.index(at) + len(at):])
            dnsurl = ("https://api.hackertarget.com/dnslookup/?q="+domain)
            dnstable = PrettyTable([f"{bcolors.WARNING}DNS LOOKUP{bcolors.ENDC}"])
            response = requests.get(dnsurl)
            html = response.content
            soup=BeautifulSoup(html,"html.parser")
            dnstable.add_row([soup])
            print(dnstable)
    except:
        print("Service Error! { DNS Lookup}")
