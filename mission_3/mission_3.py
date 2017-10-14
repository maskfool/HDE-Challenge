import httplib2
import json
from oath import _hotp
from oath._totp import totp
import hashlib

with open('personal_info.json', 'r') as f:
    data = json.load(f)

yuqi_email = data["contact_email"]

data = str(data)
http_body_data = str.encode(data)

key = yuqi_email + "HDECHALLENGE003"
key = str.encode(key)
key = _hotp.truncated_value(key)
key = str(key)
password = totp(key, period=30, hash=hashlib.sha512, format='dec10')
print("the password: " + str(password))

h = httplib2.Http()
h.add_credentials(name=yuqi_email, password=password)
resp, content = h.request("http://gaia.cs.umass.edu/wireshark-labs/HTTP-wireshark-file3.html", method="POST", headers={"Content-Type":"Application/json"},
                          body=http_body_data)
print(resp)
print(content)
