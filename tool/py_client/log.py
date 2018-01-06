
Enable=True

def log(*args, **kwargs):
    global Enable
    if Enable==True:
        print(*args, **kwargs)

