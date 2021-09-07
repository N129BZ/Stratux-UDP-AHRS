// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package main

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type AhrsRecord struct {
	Btid     uint8
	Msgid    uint8
	Mtype    uint8
	Rsvd1    uint8
	Rsvd2    uint8
	Roll     float32
	Pitch    float32
	Heading  float32
	Slipskid float32
	Yaw      float32
	G        float32
	Airspeed float32
	Palt     float32
	Vspeed   float32
	Rsvd3    uint8
	Rsvd4    uint8
	Chksum   uint16
	Etid     uint8
}

func NewAhrsRecord() *AhrsRecord {
	return &AhrsRecord{}
}

func (*AhrsRecord) Read(io *kaitai.Stream, ahrsrecord *AhrsRecord) (err error) {

	const INVALID = 32767

	tmp1, err := io.ReadU1()
	if err != nil {
		return err
	}
	ahrsrecord.Btid = tmp1

	tmp2, err := io.ReadU1()
	if err != nil {
		return err
	}
	ahrsrecord.Msgid = tmp2

	tmp3, err := io.ReadU1()
	if err != nil {
		return err
	}
	ahrsrecord.Mtype = tmp3

	tmp4, err := io.ReadU1()
	if err != nil {
		return err
	}
	ahrsrecord.Rsvd1 = tmp4

	tmp5, err := io.ReadU1()
	if err != nil {
		return err
	}
	ahrsrecord.Rsvd2 = tmp5

	roll, err := io.ReadS2be()
	if err != nil {
		return err
	}
	if roll == INVALID {
		roll = 0
	}
	ahrsrecord.Roll = float32(roll / 10)

	pitch, err := io.ReadS2be()
	if err != nil {
		return err
	}
	if pitch == INVALID {
		pitch = 0
	}
	ahrsrecord.Pitch = float32(pitch / 10)

	hdg, err := io.ReadS2be()
	if err != nil {
		return err
	}
	if hdg == INVALID {
		hdg = 0
	}
	ahrsrecord.Heading = float32(hdg / 10)

	slip, err := io.ReadS2be()
	if err != nil {
		return err
	}
	if slip == INVALID {
		slip = 0
	}
	ahrsrecord.Slipskid = float32(slip / 10)

	yaw, err := io.ReadS2be()
	if err != nil {
		return err
	}
	if yaw == INVALID {
		yaw = 0
	}
	ahrsrecord.Yaw = float32(yaw / 10)

	g, err := io.ReadS2be()
	if err != nil {
		return err
	}
	if g == INVALID {
		g = 0
	}
	ahrsrecord.G = float32(g / 10)

	aspd, err := io.ReadS2be()
	if err != nil {
		return err
	}
	if aspd == INVALID {
		aspd = 0
	}
	ahrsrecord.Airspeed = float32(aspd)

	palt, err := io.ReadU2be()
	if err != nil {
		return err
	}
	if palt == INVALID {
		palt = 0
	}
	ahrsrecord.Palt = float32(palt - 5000)

	vspd, err := io.ReadS2be()
	if err != nil {
		return err
	}
	if vspd == INVALID {
		vspd = 0
	}
	ahrsrecord.Vspeed = float32(vspd)

	tmp15, err := io.ReadU1()
	if err != nil {
		return err
	}
	ahrsrecord.Rsvd3 = tmp15

	tmp16, err := io.ReadU1()
	if err != nil {
		return err
	}
	ahrsrecord.Rsvd4 = tmp16

	tmp17, err := io.ReadU2be()
	if err != nil {
		return err
	}
	ahrsrecord.Chksum = uint16(tmp17)

	tmp18, err := io.ReadU1()
	if err != nil {
		return err
	}
	ahrsrecord.Etid = tmp18

	return err
}
