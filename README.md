# intro to golang presentation

configure gopath
```
$ cat .bash_profile

#Golang settings
export PATH=$PATH:/usr/local/opt/go/libexec/bin
export GOPATH=$HOME/Projects/go  #/Users/halyph/Projects/go
export PATH=$PATH:$GOPATH/bin
```

install dependencies
```
$ go get golang.org/x/net
$ go get golang.org/x/tools
$ go install golang.org/x/tools/cmd/present
```

```
$ present
```

kudos: http://halyph.com/blog/2015/05/18/golang-presentation-tool.html
