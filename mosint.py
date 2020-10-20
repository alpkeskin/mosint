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

connection()

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

verifyconnect()

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

pwnconnect()

def remo():
    if os.path.exists("html.txt"):
        os.remove("html.txt")
print()
api = open("api.txt", "r")
setapi = api.read()
verifyurl = "https://app.verify-email.org/api/v1/"+setapi+"/verify/"
creditsurl = "https://app.verify-email.org/api/v1/"+setapi+"/credits"
leakedpassurl = "https://scylla.sh/search?q=email:"
psbdmpurl = "https://psbdmp.ws/api/search/"
pwnedurl = "https://dehashed.com/search?query="
searchurlP="https://s.sudonull.com/?q=site%3Apastebin.com+intext%3A%22"
searchurlT="https://s.sudonull.com/?q=site%3Athrowbin.io+intext%3A%22"
if (setapi != ""):
    print('API Key : '+'\x1b[6;30;42m' + 'OK!' + '\x1b[0m')
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

    queries = [mail]
    platforms = [Platforms.GITHUB, Platforms.TWITTER, Platforms.INSTAGRAM, Platforms.PINTEREST, Platforms.SPOTIFY]
    results = sync_execute_queries(queries, platforms)
    for result in results:
        print(f"{bcolors.BOLD}{result.platform}:{bcolors.ENDC}{bcolors.WARNING} {result.message} (Success: {result.success}, Available: {result.available}){bcolors.ENDC}")
    print("")    
    print("------------------------")  
    print("")
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
    print(f"{bcolors.WARNING} -- Scanning Pastebin Dumps...{bcolors.ENDC}")
    print("")
    u = (psbdmpurl+mail)
    response4 = requests.get(u,headers=headers)
    html4 = response4.content
    lp2 = json.loads(html4)
    for i in lp2['data']:
        print(f"{bcolors.OKGREEN}|-- {bcolors.ENDC}"+"https://pastebin.com/"+i['id'])
    print("")    
    print("------------------------")  
    print("")
    print(f"{bcolors.WARNING} -- Google Searching... [Pastebin & Throwbin]{bcolors.ENDC}")
    print("")    
    x = mail.replace("@", "%40")
    u = (searchurlP+x+"%22")
    response = requests.get(u)
    html = response.content
    soup=BeautifulSoup(html,"html.parser")
    rgx = str(soup)
    urls = re.findall('http[s]?://pastebin.com(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+<', rgx)
    try:
        for x in range(len(urls)): 
            p = urls[x].replace("<", "")
            print(f"{bcolors.OKGREEN}|-- {bcolors.ENDC}"+p)

    except:
        print("Pastebin search error!")
    x = mail.replace("@", "%40")
    u = (searchurlT+x+"%22")
    response = requests.get(u)
    html = response.content
    soup=BeautifulSoup(html,"html.parser")
    rgx = str(soup)
    urls = re.findall('http[s]?://throwbin.io(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+<', rgx)
    try:
        for x in range(len(urls)): 
            t = urls[x].replace("<", "")
            print(f"{bcolors.OKGREEN}|-- {bcolors.ENDC}"+t)

    except:
        print("Throwbin search error!")
    print("")    
    print("------------------------")  
    print("")
    response = requests.get(pwnedurl+str(mail))
    html = response.content
    soup=BeautifulSoup(html,"html.parser")
    with open("html.txt","w") as file :
        file.write(str(soup))
    logfile = open("html.txt", "r") 
    c=0
    find="#ffffff;"+'"'+'>'+"Sourced"
    for line in logfile:
        if find in line.split():
            c += 1
    if c == 1:
        print(f"{bcolors.HEADER}[#]{bcolors.ENDC}" + " Pwned on "+str(c)+" breached site!")
        remo()
        continue
    elif c == 0:
        print(f"{bcolors.FAIL}[#]{bcolors.ENDC}" + " No pwnage found!")
        remo()
        continue
    else:
        print(f"{bcolors.HEADER}[#]{bcolors.ENDC}" + " Pwned on "+str(c)+" breached sites!")
        remo()
        continue
