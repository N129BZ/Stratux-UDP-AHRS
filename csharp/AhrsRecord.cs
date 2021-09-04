// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild
using System;
using Newtonsoft.Json;


namespace Kaitai
{
    public partial class AhrsRecord : KaitaiStruct
    {
        public static AhrsRecord FromFile(string fileName)
        {
            return new AhrsRecord(new KaitaiStream(fileName));
        }

        public AhrsRecord(KaitaiStream p__io, KaitaiStruct p__parent = null, AhrsRecord p__root = null) : base(p__io)
        {
            m_parent = p__parent;
            m_root = p__root ?? this;
            _read();
        }

        public Record Record { 
            get { return _record; }
        }

        private void _read()
        {
            m_io.ReadU1();
            m_io.ReadU1();
            m_io.ReadU1();
            m_io.ReadU1();
            m_io.ReadU1();
            _roll = m_io.ReadS2be();
            _pitch = m_io.ReadS2be();
            _heading = m_io.ReadS2be();
            _slipskid = m_io.ReadS2be();
            _yaw = m_io.ReadS2be();
            _gs = m_io.ReadS2be();
            _airspeed = m_io.ReadS2be();
            _palt = m_io.ReadU2be();
            _vspeed = m_io.ReadS2be();
            
            _record = new Record(_roll, _pitch, _heading, _slipskid, _yaw, _gs, _airspeed, _palt, _vspeed);
        }
        
        private short _roll;
        private short _pitch;
        private short _heading;
        private short _slipskid;
        private short _yaw;
        private short _gs;
        private short _airspeed;
        private ushort _palt;
        private short _vspeed;
        private AhrsRecord m_root;
        private KaitaiStruct m_parent;
        private Record _record;

        
    }

    public class Record 
    {
        public Record(short roll, short pitch, short heading, 
                      short slip, short yaw, short gforce, short airspd, 
                      ushort palt, short vspd) {
            Roll = (float)(roll == 32767 ? 0 : roll / 10);
            Pitch = (float)(pitch == 32767 ? 0 : pitch / 10);
            Heading = (float)(heading == 32767 ? 0 : heading / 10);
            Slipskid = (float)(slip == 32767 ? 0 : slip / 10);
            Yaw = (float)(yaw == 32767 ? 0 : yaw / 10);
            GForce = (float)(gforce == 32767 ? 0 : gforce / 10);
            Airspeed = (short)(airspd == 32767 ? 0 : airspd);
            Palt = (float)(palt == 32767 ? 0 : palt - 4999.5);
            Vspeed = (short)(vspd == 32767 ? 0 : vspd);
        }

        public override string ToString()
        {
            return JsonConvert.SerializeObject(this, Formatting.None);
        }

        public float Roll { get; set; }
        public float Pitch { get; set; } 
        public float Heading { get; set; } 
        public float Slipskid { get; set; } 
        public float Yaw { get; set; } 
        public float GForce { get; set; }
        public short Airspeed { get; set; } 
        public float Palt { get; set; } 
        public short Vspeed { get; set; } 
    }
}
