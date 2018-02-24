import argparse
import os
import json
import time
import log
print = log.log

def parse_args(title):
    parser = argparse.ArgumentParser(description=title, formatter_class=argparse.RawDescriptionHelpFormatter)
    parser.add_argument("--cfg", default="./_cfg.json", help="configure file", type=str)
    args = parser.parse_args()

    jsonfile = os.path.abspath(os.path.dirname(__file__)) + "/" + args.cfg
    if os.path.exists(jsonfile) == False:
        print("ERROR: cfg file not found. path = ", jsonfile)
        exit(0)

    f = open(jsonfile, 'rt')
    cfg = json.loads(f.read())
    f.close()
    return args, cfg
    
    
def timestamp():
    return int(round(time.time() * 1000))