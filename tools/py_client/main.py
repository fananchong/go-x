import user
from util import *
import log
print = log.log

if __name__ == "__main__":
    args, cfg = parse_args("py_client")
    u1 = user.User(1, args, cfg)
    
    if u1.login() == False:
        print("u1 request login fail.")
        exit(0)
    
    print("over.")
    