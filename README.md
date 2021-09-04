# Parse Stratux UDP port 4000 AHRS messages 
These apps parse Stratux AHRS UDP messages on port 4000. All implementations use the Kaitai struct compiler to generate a javascript, python, or c# class that will decode the incoming byte array.

###
   At run time, requires a reference to KaitaiStream, which is installed as part of the Kaitai struct compiler. The compiler generates your preferred language from the ahrs.ksy yaml file. 
   ```
   Samples: 
       
       ksc -t javascript ahrs.ksy
       ksc -t python ahrs.ksy
   ```
###
   Requirements, see: https://github.com/kaitai-io/kaitai_struct_compiler/releases
