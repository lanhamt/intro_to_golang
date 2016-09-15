# intro to golang presentation

## installing go
```
$ brew install go
```
install rep and configure gopath (in your home directory, otherwise adjust as needed)
```
$ git clone https://github.com/lanhamt/intro_to_golang
$ export PATH=$PATH:/usr/local/opt/go/libexec/bin
$ export GOPATH=$HOME/intro_to_golang
$ export PATH=$PATH:$GOPATH/bin
```
install dependencies
```
$ cd into_to_golang
$ go get golang.org/x/net
$ go get golang.org/x/tools
$ go install golang.org/x/tools/cmd/present
```

## start presentation
```
$ present
```
[[http://localhost:3999]]

kudos: http://halyph.com/blog/2015/05/18/golang-presentation-tool.html
