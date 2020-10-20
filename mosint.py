#https://github.com/alpkeskin
import re
import requests
from bs4 import BeautifulSoup
import json
import os
from socialscan.util import Platforms, sync_execute_queries
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

print(f'''{bcolors.BOLD}
___  ________ _____ _____ _   _ _____ 
|  \/  |  _  /  ___|_   _| \ | |_   _|
| .  . | | | \ `--.  | | |  \| | | |  
| |\/| | | | |`--. \ | | | . ` | | |  
| |  | \ \_/ /\__/ /_| |_| |\  | | |  
\_|  |_/\___/\____/ \___/\_| \_/ \_/ 
{bcolors.ENDC}
v{bcolors.BOLD}1.2{bcolors.ENDC}
github.com/{bcolors.BOLD}alpkeskin{bcolors.ENDC}
''')

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


def pwnconnect(url='https://scylla.sh/' , timeout=20):
    try:
        req = requests.get(url, timeout=timeout)
        req.raise_for_status()
        print(f"{bcolors.OKGREEN}[+] You're connected to leaked database.{bcolors.ENDC}")
        return True
    except requests.HTTPError as e:
        print("[-] Checking internet connection failed, status code {0}.".format(e.response.status_code))
    except requests.ConnectionError:
        print(f"{bcolors.FAIL}[-] No leaked database available.{bcolors.ENDC}")
    return False

configfile = open('config.json', "r")
conf = json.loads(configfile.read())
for i in conf:
    checkconnect = (i['Check Connections'])
    setapi = (i['API Key'])
    socialscan = (i['Social Scan'])
    leakeddb = (i['Leaked DB'])
    pastebindumps = (i['Pastebin Dumps'])
    googlesearch = (i['Google Search'])
    dns = (i['DNS Lookup'])
if (checkconnect == "True" or checkconnect == "T" or checkconnect == "true"):
    connection()
    verifyconnect()
    pwnconnect()

verifyurl = "https://app.verify-email.org/api/v1/"+setapi+"/verify/"
creditsurl = "https://app.verify-email.org/api/v1/"+setapi+"/credits"
leakedpassurl = "https://scylla.sh/search?q=email:"
psbdmpurl = "https://psbdmp.ws/api/search/"
searchurlP="https://s.sudonull.com/?q=site%3Apastebin.com+intext%3A%22"
searchurlT="https://s.sudonull.com/?q=site%3Athrowbin.io+intext%3A%22"
if (setapi != ""):
    print('API Key : '+'\x1b[6;30;42m' + ' OK! ' + '\x1b[0m')
    response = requests.get(creditsurl)
    html = response.content
    soup=BeautifulSoup(html,"html.parser")
    strsoup = str(soup)
    data = json.loads(strsoup)
    print(f"{bcolors.UNDERLINE}Credit:{bcolors.ENDC}" + str(data['credits']))
    print("")
while True:
    mail=input(f"{bcolors.OKBLUE}MAIL > {bcolors.ENDC}")
    if (mail == "q"):
        print("Thank you for using "+f"{bcolors.BOLD}MOSINT{bcolors.ENDC}.")
        break
    elif (setapi != ""): 
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
        if (pastebindumps == "True" or pastebindumps == "T" or pastebindumps == "true"):
            print(f"{bcolors.WARNING} -- Scanning Pastebin Dumps...{bcolors.ENDC}")
            print("")
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
        print("Pastebin Dump DB Connection Error!")
    try:
        if (googlesearch == "True" or googlesearch == "T" or googlesearch == "true"):
            print(f"{bcolors.WARNING} -- Google Searching... [Pastebin & Throwbin]{bcolors.ENDC}")
            print("")    
            x = mail.replace("@", "%40")
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
