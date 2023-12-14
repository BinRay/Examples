gomonkey is a library for mocking functions in Go.

## Test Method
```bash
$ cd test 
$ go test -gcflags=all=-l
# debug remotely
$ dlv test --headless --listen=:2346 --api-version=2 --accept-multiclient -gcflags=all=-l apply_method_test.go -- -test.v
```
The -gcflags flag is used to pass options to the Go compiler. In this case, all=-l is being passed, which instructs the compiler to disable inlining of functions. 