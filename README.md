# Collider for CloudFoundry
A websocket-based signaling server in Go based on appartc project

# Building

1. Install the Go tools and workspaces as documented at [http://golang.org/doc/install](http://golang.org/doc/install) and [http://golang.org/doc/code.html](http://golang.org/doc/code.html) 

2. Checkout this repository
```
cd $GOPATH
go get github.com/tkausch/collidermain
```

3. Now you should be able to run ```mycollider``` locally
```
$ cd $GOPATH/src/github.com/tkausch/collidermain
$ go run mycollider.go 
2017/03/07 16:32:13 Starting collider: tls = false, port = 8090, room-server=https://appr.tc/
```

# Deploy on Cloud Foundry

1. Install the ```glide``` package manager 
```
brew install glide
```  

2. Update dependencies with glide. This installs all depencies so they can be pushed to cloudfoundry.
```
cd $GOPATH/src/github.com/tkausch/collidermain
glide up
glide install 
```

3. Login into Cloud Foundry and push the go app:
```
cf push mycollider
```

# Configure Collider Environment

To define your room server you can set the following environment variable 
```
ROOM_SRV = https://appr.tc
```
 
