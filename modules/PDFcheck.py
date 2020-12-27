from googlesearch import search
import requests,re,PyPDF2,tabula,os

def PDFcheck(mail,_verbose=None):
	domain = mail.split("@")[1]
	term = "site:"+domain+" filetype:PDF intext:"+'"'+"email"+'"'
	try:
		data = search(term, num_results=5)
		for i in data:
			r = requests.get(i, stream=True)
			with open('data.pdf', 'wb') as f:
				f.write(r.content)
			pdfFileObj = open('data.pdf', 'rb')
			for pageNumber in range(1,3):
				tabula.convert_into("data.pdf","out.txt",pages=pageNumber,silent=True)
				file = open("out.txt","r",encoding="utf-8")
				read = file.read()
				findPDFs= re.findall('[\w\.-]+@[a-z0-9\.-]+', read)
				try:
					if(findPDFs[0] is not None):
						for pdfs in findPDFs:
							print(pdfs)
				except:
					pass
			pdfFileObj.close()
			file.close()
			if os.path.exists("data.pdf"):
				os.remove("data.pdf")
			if os.path.exists("out.txt"):
				os.remove("out.txt")
	except:
		print("PDF Search error!")