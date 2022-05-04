# mosint v2.1
# Author: Alp Keskin
# Github: github.com/alpkeskin
# Website: https://imalp.co
from socialscan.util import Platforms, sync_execute_queries
import argparse


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument("-e", "--email", type=str, required=True, help="Email")
    return parser.parse_args()


def main():
    try:
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
        f = open("socialscantempresult.txt", "w")
        for result in results:
            if not result.available:
                f.write(f"{result.platform}\n")
        f.close()
    except Exception as e:
        print(e)


main()
