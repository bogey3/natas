import requests
import re
from requests.auth import HTTPBasicAuth
def solve0(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    resp = session.get("http://natas0.natas.labs.overthewire.org/")
    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw

