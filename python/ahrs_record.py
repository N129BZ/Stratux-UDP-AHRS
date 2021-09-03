from pkg_resources import parse_version
import kaitaistruct
from kaitaistruct import KaitaiStruct, KaitaiStream, BytesIO


if parse_version(kaitaistruct.__version__) < parse_version('0.9'):
    raise Exception("Incompatible Kaitai Struct Python API: 0.9 or later is required, but you have %s" % (kaitaistruct.__version__))

class AhrsRecord(KaitaiStruct):
    def __init__(self, _io, _parent=None, _root=None):
        self._io = _io
        self._parent = _parent
        self._root = _root if _root else self
        self._read()

    def _read(self):
        self.btid = self._io.read_u1()
        self.msgid = self._io.read_u1()
        self.mtype = self._io.read_u1()
        self.rsvd1 = self._io.read_u1()
        self.rsvd2 = self._io.read_u1()
        r = self._io.read_s2be()
        self.roll = 0 if r == 32767 else r / 10 
        p = self._io.read_s2be()
        self.pitch = 0 if p == 32767 else p / 10
        h = self._io.read_s2be()
        self.heading = 0 if h == 32767 else h / 10
        s = self._io.read_s2be()
        self.slipskid = 0 if s == 32767 else s / 10
        y = self._io.read_s2be()
        self.yaw = 0 if y == 32767 else y / 10
        gs = self._io.read_s2be()
        self.g = 0 if gs == 32767 else gs / 10
        a = self._io.read_s2be()
        self.airspeed = 0 if a == 32767 else a
        pa = self._io.read_u2be()
        self.palt = 0 if pa == 32767 else pa - 5000
        v = self._io.read_s2be()
        self.vspeed = 0 if v == 32767 else v
        self.rsvd3 = self._io.read_u1()
        self.rsvd4 = self._io.read_u1()
        self.chksum = self._io.read_u2be()
        self.etid = self._io.read_u1()
