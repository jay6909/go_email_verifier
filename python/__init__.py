
import ctypes
import os
import platform


_dir=os.path.dirname(__file__)
_system=platform.system()

if _system == 'Windows':
    _lib_name="libverifier.dll"
elif _system == 'Darwin':
    _lib_name="libverifier.dylib"
else:
    _lib_name="libverifier.so"
    
_lib_path=os.path.join(_dir, _lib_name)

#load the library 
_lib=ctypes.CDLL(_lib_path)



def verify_email(email:str)->bool:
    _lib.Verify.argtypes=[ctypes.c_char_p]
    _lib.Verify.restype=ctypes.c_bool
    
    result =_lib.Verify(email.encode('utf-8'))
    return bool(result)