// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type AhrsRecord struct {
	Btid uint8
	Msgid uint8
	Mtype uint8
	Rsvd1 uint8
	Rsvd2 uint8
	Roll float32
	Pitch float32
	Heading float32
	Slipskid float32
	Yaw float32
	Gs float32
	Airspeed int16
	Palt uint16
	Vspeed float32
	Rsvd3 uint8
	Rsvd4 uint8
	Chksum uint16
	Etid uint8
	_io *kaitai.Stream
	_root *AhrsRecord
	_parent interface{}
}
func NewAhrsRecord() *AhrsRecord {
	return &AhrsRecord{
	}
}

func (this *AhrsRecord) Read(io *kaitai.Stream, parent interface{}, root *AhrsRecord) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Btid = tmp1
	tmp2, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Msgid = tmp2
	tmp3, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Mtype = tmp3 
	tmp4, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Rsvd1 = tmp4
	tmp5, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Rsvd2 = tmp5
	tmp6, err := this._io.ReadS2be()
	if err != nil {
		return err
	}
	this.Roll = float32(tmp6 / 10)
	tmp7, err := this._io.ReadS2be()
	if err != nil {
		return err
	}
	this.Pitch = float32(tmp7 / 10)
	tmp8, err := this._io.ReadU2be()
	if err != nil {
		return err
	}
	this.Heading = float32(tmp8 / 10)
	tmp9, err := this._io.ReadS2be()
	if err != nil {
		return err
	}
	this.Slipskid = ifloat32(tmp9 / 10)
	tmp10, err := this._io.ReadS2be()
	if err != nil {
		return err
	}
	this.Yaw = float32(tmp10 / 10)
	tmp11, err := this._io.ReadS2be()
	if err != nil {
		return err
	}
	this.Gs = float32(tmp11 / 10)
	tmp12, err := this._io.ReadU2be()
	if err != nil {
		return err
	}
	this.Airspeed = int16(tmp12)
	tmp13, err := this._io.ReadU2be()
	if err != nil {int16
		return err
	}
	this.Palt = uint16(tmp13 - 4999.5)
	tmp14, err := this._io.ReadU2be()
	if err != nil {
		return err
	}
	this.Vspeed = int16(tmp14)
	tmp15, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Rsvd3 = tmp15
	tmp16, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Rsvd4 = tmp16
	tmp17, err := this._io.ReadU2be()
	if err != nil {
		return err
	}
	this.Chksum = uint16(tmp17)
	tmp18, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Etid = tmp18
	return err
}
