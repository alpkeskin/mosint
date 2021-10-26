from socialscan.util import Platforms, sync_execute_queries
import argparse


class bcolors:
    OKGREEN = "\033[92m"
    FAIL = "\033[91m"
    ENDC = "\033[0m"


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument("-e", "--email", type=str, required=True, help="Email")
    return parser.parse_args()


def main():
    args = parse_args()
    mail = args.email
    q = [mail]
    platforms = [
        Platforms.GITHUB,
        Platforms.TWITTER,
        Platforms.INSTAGRAM,
        Platforms.PINTEREST,
        Platforms.SPOTIFY,
        Platforms.FIREFOX,
    ]
    results = sync_execute_queries(q, platforms)
    count = 0
    for result in results:
        if not result.available:
            count = count + 1
            print(f"{bcolors.OKGREEN}[+] {result.platform}{bcolors.ENDC}")
    if count == 0:
        print(f"{bcolors.FAIL}[-] Not Found!{bcolors.ENDC}")


main()
