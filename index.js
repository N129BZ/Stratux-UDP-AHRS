const dgram =  require('dgram');
const AhrsRecord = require('./AhrsRecord');
const KaitaiStream = require('kaitai-struct/KaitaiStream');

var PORT = 4000;

var client = dgram.createSocket('udp4');

client.on('error', (err) => {
    console.log(err.message);
});

client.on('listening', () => {
    var address = client.address();
    console.log('UDP Client listening on ' + address.address + ':' + address.port);
});
                                             
client.on('message', (message, remote) => {
    var s = "\\n";
    if (message[1] == 76) {
        var record = new AhrsRecord(new KaitaiStream(message));
        console.log(record);
    }
});

client.bind(PORT);
       

//###########################################################################################################################
//                 Stratux AHRS message
//  -------------------------------------------------------------------------------------------------------------------------
//  BOM  |ID/TYP|resvd|roll  |  pitch  |   hdg   |slip |  yaw    |  G's | airspd  | palt   | vspd    | resvd   | chksm | EOM  
//  -------------------------------------------------------------------------------------------------------------------------                   
//  [126, 76, 69, 1, 1, 0, 0, 255, 255, 127, 255, 0, 0, 127, 255, 0, 10, 127, 255, 24, 162, 255, 255, 127, 255, 42, 249, 126]
//############################################################################################################################


