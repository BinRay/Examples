# Go installation
https://go.dev/doc/install
1. `wget https://go.dev/dl/go1.21.4.linux-amd64.tar.gz`
2. `rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz`
3. `export PATH=$PATH:/usr/local/go/bin`

# Go evnironment
1.`go env -w GOINSECURE=private.repo.com` // set private repo, If this variable is set to a non-null value, then the Go program will allow HTTP requests over untrusted HTTPS certificates.
2.`go env -w GOPROXY=https://goproxy.io,direct` // set proxy
3.`go env -w GOPROXY=https://goproxy.cn` // another better proxy
4.`go env -w GONOPROXY=private.repo.com` //  bypass the proxy and download modules directly from specific hosts or domains

# Build
- CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
