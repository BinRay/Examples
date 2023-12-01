A distributed, reliable key-value store for the most critical data of a distributed system

# cmd
```bash
etcdctl --endpoints=192.168.12.248:2379  

# get all keys
etcdctl get --prefix ""

# get the value of key Greeter
etcdctl get Greeter    
```