meta:
  id: ahrs_record
  endian: be
seq:
  - id: btid
    type: u1
  - id: msgid
    type: u1
  - id: mtype
    type: u1
  - id: rsvd1
    type: u1
  - id: rsvd2
    type: u1
  - id: roll
    type: s2
  - id: pitch
    type: s2
  - id: heading
    type: s2
  - id: slipskid
    type: s2
  - id: yaw
    type: s2
  - id: gs
    type: s2
  - id: airspeed
    type: s2
  - id: palt
    type: u2
  - id: vspeed
    type: s2   
  - id: rsvd3
    type: u1
  - id: rsvd4
    type: u1
  - id: chksum
    type: u2
  - id: etid
    type: u1
