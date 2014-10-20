golang-tdd
==========

Developing in Golang using Test Driven Development



Build
=====

If this repository is checked out in a directory structure organized accordingly to the standard described in [the official documentation](http://golang.org/pkg/go/build/), the ```go``` tool can be used to install and test the code.

In order to build the code, run

```
go install github.com/golang4dummies/golang-tdd
``` 

This will compile the package and generate an executable in ```$GOPATH/bin```.

Alternatively, you can cd into the repository's root with

```
cd $GOPATH/golang4dummies/golang-tdd
```

and then simply run

```
go install
```


Note that the executable will be named after the repository name, not the Go file.


Run
===

Install the package, then run it with

```
$GOPATH/bin/golang-tdd
```

If ```$GOPATH/bin``` has been included in ```$PATH```, as suggested in [the official documentation](http://golang.org/pkg/go/build/), you can simply run the program with

```
golang-tdd
```


Uninstall
=========

Run

```
go clean -i
```