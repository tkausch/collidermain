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

2. Update and install dependencies with glide. When an app is deployed App Cloud reads the glide.yaml file and installs the appropriate depencencies. Run `glide install```command in your local directory to install the dependencies, preparing your system for running the app locally. The command will create a glide.lock file. It is good practice to push this lock file to the cloud as well since it specifies which versions of depencencies to install exactly. 

```
cd $GOPATH/src/github.com/tkausch/collidermain
glide up
(optional:) glide install 
```

3. Login into Cloud Foundry and push the go app. When running ```cf push``, we will also push our whole dependencies folder ```vendor```. Since Cloud Foundry runs ```glide install``` anyways, this is redundant. To save bandwidth and time you can create a ```.cfignore``` file which works just like a ```.gitignore```file and tells CloudFoundry what files should be excluded when pushing.

```
cf push mycollider
```

# Configure Collider Environment

To define your room server you can set the following environment variable 
```
ROOM_SRV = https://appr.tc
```
 
