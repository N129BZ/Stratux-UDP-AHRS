// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type AhrsRecord struct {
	Btid uint8
	Msgid uint8
	Mtype uint8
	Rsvd1 uint8
	Rsvd2 uint8
	Roll int16
	Pitch int16
	Heading uint16
	Slipskid int16
	Yaw int16
	Gs int16
	Airspeed uint16
	Palt uint16
	Vspeed uint16
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
	this.Roll = int16(tmp6)
	tmp7, err := this._io.ReadS2be()
	if err != nil {
		return err
	}
	this.Pitch = int16(tmp7)
	tmp8, err := this._io.ReadU2be()
	if err != nil {
		return err
	}
	this.Heading = uint16(tmp8)
	tmp9, err := this._io.ReadS2be()
	if err != nil {
		return err
	}
	this.Slipskid = int16(tmp9)
	tmp10, err := this._io.ReadS2be()
	if err != nil {
		return err
	}
	this.Yaw = int16(tmp10)
	tmp11, err := this._io.ReadS2be()
	if err != nil {
		return err
	}
	this.Gs = int16(tmp11)
	tmp12, err := this._io.ReadU2be()
	if err != nil {
		return err
	}
	this.Airspeed = uint16(tmp12)
	tmp13, err := this._io.ReadU2be()
	if err != nil {
		return err
	}
	this.Palt = uint16(tmp13)
	tmp14, err := this._io.ReadU2be()
	if err != nil {
		return err
	}
	this.Vspeed = uint16(tmp14)
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
