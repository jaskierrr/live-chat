import valkey

def get_user(id: int):
    try:
        r = valkey.Valkey(host='localhost', port=6379, decode_responses=True)
        r.set('foo', id)
        result = r.get('foo')
    except:
        result = None
        print("can't connect to redis")

    return result
