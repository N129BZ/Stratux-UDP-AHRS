// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

(function (root, factory) {
  if (typeof define === 'function' && define.amd) {
    define(['kaitai-struct/KaitaiStream'], factory);
  } else if (typeof module === 'object' && module.exports) {
    module.exports = factory(require('kaitai-struct/KaitaiStream'));
  } else {
    root.AhrsRecord = factory(root.KaitaiStream);
  }
}(this, function (KaitaiStream) {
var AhrsRecord = (function() {
  function AhrsRecord(_io, _parent, _root) {
    this._io = _io;
    this._parent = _parent;
    this._root = _root || this;

    this._read();
  }
  AhrsRecord.prototype._read = function() {
    
    this.btid     = this._io.readU1();
    this.msgid    = this._io.readU1();
    this.mtype    = this._io.readU1();
    this.rsvd1    = this._io.readU1();
    this.rsvd2    = this._io.readU1();
    var r         = this._io.readS2be();
    this.roll     = r != 32767 ? r / 10 : 0;
    var p         = this._io.readS2be();
    this.pitch    = p != 32767 ? p / 10 : 0;
    var h         = this._io.readS2be();
    this.heading  = h != 32767 ? h / 10 : 0;
    var s         = this._io.readS2be();
    this.slipskid = s != 32767 ? s / 10 : 0;
    var y         = this._io.readS2be();
    this.yaw      = y != 32767 ? y / 10 : 0;
    var g         = this._io.readS2be();
    this.gs       = g != 32767 ? g / 10 : 0;
    var a         = this._io.readS2be();
    this.airspeed = a != 32767 ? a : 0;
    this.palt     = (this._io.readU2be()) - 5000;
    var v         = this._io.readS2be();
    this.vspeed   = v != 32767 ? v : 0;
    this.rsvd3    = this._io.readU1();
    this.rsvd4    = this._io.readU1();
    this.chksum   = this._io.readU2be();
    this.etid     = this._io.readU1();
  }

  return AhrsRecord;
})();
return AhrsRecord;
}));
