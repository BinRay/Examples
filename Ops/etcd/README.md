A distributed, reliable key-value store for the most critical data of a distributed system

# cmd

etcdctl --endpoints=192.168.12.248:2379  
etcdctl get --prefix "" // get all keys  
etcdctl get Greeter // get the value of key Greeter  