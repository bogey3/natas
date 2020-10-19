import requests
import re
import base64
from requests.auth import HTTPBasicAuth
def solve11(username, password):
    cookieData = b'{"showpassword":"yes","bgcolor":"#ffffff"}'
    xorKey = b"qw8J"
    encryptedData = b""
    for i in range(len(cookieData)):
        encryptedData = encryptedData + bytes([cookieData[i]^xorKey[i%len(xorKey)]])
    newCookie = base64.b64encode(encryptedData)
    session = requests.session()
    session.auth = HTTPBasicAuth(username, password)
    session.cookies["data"] = newCookie.decode("utf-8")
    resp = session.get("http://natas11.natas.labs.overthewire.org/")
    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw