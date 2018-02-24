import urllib.request
import traceback
from proto.message import *
import proto.login_pb2
from util import *
import log
print = log.log

class LoginMode():
    def __init__(self, index, args, cfg):
        self.index = index
        self.args = args
        self.cfg = cfg
        self.addr = self.cfg["addr"]
        self.port = self.cfg["port"]
        self.version = self.cfg["version"]
        self.cookie = ""
        self.account_name = ""
        self.account_id = 0
        
    def login(self):
        return self.login_detail(           \
            self.cfg["login"]["account"],   \
            self.cfg["login"]["password"],  \
            self.cfg["login"]["mode"],      \
            self.cfg["login"]["userdata"],  \
            )
        
    def login_detail(self, account, password, mode, userdata):
        try:
            cmd = LoginCmd.Login.value
            msg = proto.login_pb2.MsgLogin()
            msg.Account = account
            msg.Password = password
            msg.Mode = mode
            #msg.userdata = userdata
            request, result = self.send(cmd, msg)
            pmsg = proto.login_pb2.MsgLoginResult()
            pmsg.ParseFromString(result)
        
            print("MsgLoginResult.Err:", pmsg.Err)
            
        except Exception as e:
            print(traceback.format_exc())
            return False
        return True
            
    def send(self, cmd, msg):
        url, postdata = msgurl(self.addr, self.port, self.version, cmd, msg)
        request = urllib.request.urlopen(url, postdata)
        result = request.read()
        return request, result
        