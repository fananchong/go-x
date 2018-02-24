from user_login import *
import log
print = log.log

class User():
    def __init__(self, index, args, cfg):
        self.index = index
        self.args = args
        self.cfg = cfg
        self.login_module = LoginMode(index, args, cfg)
        
    def login(self):
        self.login_module.login()