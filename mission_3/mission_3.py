import httplib2
import json
import struct
import time
import hmac
import base64


num_digits = 10
timestep = 30
T0 = 0

# host = "http://127.0.0.1:8000/"
host = "https://hdechallenge-solve.appspot.com/challenge/003/endpoint"

class TOTP_yuqi:
    def __init__(self):
        pass

    def int2byte(self, i):
        return struct.pack('>Q', int(i))

    def truncated_value(self, h):
        # function borrowed from OATH library
        v = h[-1]
        if not isinstance(v, int): v = ord(v)  # Python 2.x
        offset = v & 0xF
        (value,) = struct.unpack('>I', h[offset:offset + 4])
        return value & 0x7FFFFFFF

    def calculate(self, key=None, num_digits=10, timestep=30, T0=0):

        Counter = int(time.time() - T0) // timestep

        key_bytes = str.encode(key)

        Counter_bytes = self.int2byte(Counter)

        val_before_trunc = hmac.new(key=key_bytes, msg=Counter_bytes, digestmod="sha512").digest()

        passwd_before_zeropadding = self.truncated_value(val_before_trunc)

        passwd = str(passwd_before_zeropadding).zfill(num_digits)

        return passwd


with open('personal_info.json', 'r') as f:
    data = json.load(f)

yuqi_email = data["contact_email"]

http_body_data = json.dumps(data)

yuqi_email = str(yuqi_email)

key = yuqi_email + "HDECHALLENGE003"

passwd = TOTP_yuqi.calculate(self=TOTP_yuqi(), key=key, num_digits=10, timestep=timestep, T0=0)
passwd = str(passwd)
print("password: " + passwd)

content_length = len(http_body_data)

# auth_raw_str = yuqi_email + ":" + passwd
# auth_raw = str.encode(auth_raw_str)
# auth = base64.encodebytes(auth_raw)
# auth = base64.decodebytes(auth)

h = httplib2.Http()
h.add_credentials(name=yuqi_email, password=passwd)
resp, content, request_headers, request_body = h.request(host, method="POST",
                                                         headers={'Content-Type': 'application/json', 'Accept': '*/*', 'content-length': str(content_length),
                                                                  },
                                                         body=http_body_data)

print("request headers: ")
print(request_headers)
print("request body: ")
print(request_body)
print("response:")
print(resp)
print("content:")
print(content)

#
# global_headers = {"Authorization": "Basic XXXXXXX"}
# client = Client(host='base_url', request_headers=global_headers)
# query_params={"hello":0, "world":1}
# request_headers={"X-Test": "test"}
# data={"some": 1, "awesome": 2, "data": 3}
# response = client.your.api._(param).call.post(request_body=data,
#                                               query_params=query_params,
#                                               request_headers=request_headers)
# print response.status_code
# print response.headers
# print response.body
