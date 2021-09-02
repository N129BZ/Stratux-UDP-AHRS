import socket
from kaitaistruct import KaitaiStream, BytesIO
from ahrs_record import AhrsRecord
import json
sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM, socket.IPPROTO_UDP)

udp_host = "0.0.0.0" 
udp_port = 4000	
	                
sock.bind((udp_host,udp_port))

while True:
    data,addr = sock.recvfrom(128)
    if data[1] == 76:
        print(AhrsRecord(KaitaiStream(BytesIO(data))))
    
