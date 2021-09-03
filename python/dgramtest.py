import socket
from kaitaistruct import KaitaiStream, BytesIO
from ahrs_record import AhrsRecord

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM, socket.IPPROTO_UDP)

PORT = 4000
ADDRESS = "0.0.0.0"
ID_AHRS = 76
	                
sock.bind((ADDRESS,PORT))

while True:
    data,addr = sock.recvfrom(128)
    if data[1] == ID_AHRS:
        print(vars(AhrsRecord(KaitaiStream(BytesIO(data)))))