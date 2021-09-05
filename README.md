# Parse Stratux UDP port 4000 AHRS messages 
These apps parse Stratux AHRS UDP messages on port 4000. All implementations use the Kaitai struct compiler to generate a go, csharp, python, or javascript class (and many other languages) that will decode the incoming byte array.

###
   At run time, requires a reference to KaitaiStream, which is installed as part of the Kaitai struct compiler. The compiler generates a class object in your preferred language from the ahrs.ksy yaml file. 
   ```
   Samples: 
       
       ksc -t javascript ahrs.ksy
       ksc -t python ahrs.ksy
       ksc -t csharp ahrs.ksy
       ksc -t go ahrs.ksy
       
   Or specify Json output:
   
      ksc -t python --ksc-json-output ahrs.ksy
      
      etc.
   
   ```
###
   Requirements, see: https://github.com/kaitai-io/kaitai_struct_compiler/releases
