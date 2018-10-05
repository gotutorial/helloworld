# helloworld

This repo contains simple helloworld applications which helps you to set up, run your first application first using **go** CLI command and then enhacing your **helloworld** application using **glide**.

You can find two projects in this repository:

- **helloworld-basic**: A Hello World CLI Application which you use *go* CLI commands to build and run your application
- **helloworld-glide**: A Hello World CLI Application using glide package manager CLI commnads to build and run your application

## Installing Go

In order to install Go framework follow these instructions:

### Mac OS
1) Create Directories
```
mkdir $HOME/go
mkdir -p $HOME/go/src/github.com/user
```
2) Setup your paths, you can have these exports in your user home **.bash_profile** file
```
export GOPATH=$HOME/go
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```
3) Install Go
```
brew install go
```

## Write you first application
1) Create a folder anywhere you like and call it anything you want. I called it **helloworld-basic** and you can find it in this repository
2) Create a file with the extension of **.go**. You can give any name to that file. I named it **main.go**.
