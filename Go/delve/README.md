![Delve](https://raw.githubusercontent.com/go-delve/delve/master/assets/delve_horizontal.png)

Installation
------------

```bash
go install github.com/go-delve/delve/cmd/dlv@master
```

Getting Started
---------------

```bash
dlv debug github.com/me/foo/cmd/foo -- -arg1 value // pass arguments to the program
```
```bash
dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient // start a headless debug server
```

```bash
# Remote debug command on vm, so that you can debug on your local Goland IDE with go remote debug configuration.
dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient ./cmd -- -c ./conf/app_binliu.yaml
```
