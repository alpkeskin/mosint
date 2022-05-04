# mosint v2.1
# Author: Alp Keskin
# Github: github.com/alpkeskin
# Website: https://imalp.co
from genericpath import exists
import trio
import httpx
import argparse
from holehe.modules.social_media.wattpad import wattpad
from holehe.modules.social_media.discord import discord
from holehe.modules.social_media.plurk import plurk
from holehe.modules.social_media.imgur import imgur


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument("-e", "--email", type=str, required=True, help="Email")
    return parser.parse_args()


async def main():
    try:
        args = parse_args()
        mail = args.email
        out = []
        client = httpx.AsyncClient()

        await wattpad(mail, client, out)
        await discord(mail, client, out)
        await plurk(mail, client, out)
        await imgur(mail, client, out)
        file = open("holehetempresult.txt", "w")
        for i in out:
            if i["exists"]:
                file.write(i["name"])
        await client.aclose()
        file.close()
    except Exception as e:
        print(e)


trio.run(main)
