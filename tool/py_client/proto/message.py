import hashlib
from enum import Enum
from util import *

SIGN1 = "5UY6$f$h"
SIGN2 = "3wokZB%q"
SIGN3 = "%2Fi9TRf"
    
def msgurl(addr, port, version, cmd, msg):
    t = str(timestamp())
    c = str(cmd)
    s = hashlib.md5((SIGN1 + c + SIGN2 + t + SIGN3 + version).encode('utf-8')).hexdigest()
    d = msg
    url = "http://%s:%d/msg?c=%s&t=%s&d=%s&s=%s" % (addr, port, c, t, d, s)
    return url
    
class LoginCmd(Enum):
    Login=1
    