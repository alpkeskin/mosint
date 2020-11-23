from insides.bcolors import bcolors

def Header(title):
	print("")
	print("-------------------------------")
	print(f"{bcolors.OKGREEN}>{bcolors.ENDC}"+title)
	print("-------------------------------")