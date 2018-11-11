# helloworld

This repo contains simple helloworld applications which helps you to set up, run your first application first using **go** CLI command and then enhacing your **helloworld** application using **glide**.

You can find two projects in this repository:

- **helloworld-basic**: A Hello World CLI Application which you use *go* CLI commands to build and run your application
- **helloworld-glide**: A Hello World CLI Application using glide package manager CLI commnads to build and run your application


## Runing this repo helloworld samples
In order to run this application without any error and running make sure you have already installed Go and Glide

### Installing Go
In order to install Go framework on visit my tutorial page here:
http://kabiliravi.com/index.php/software/programming/golang/install/

### Installing Go
In order to install Go framework on visit my tutorial page here:
http://kabiliravi.com/index.php/software/programming/golang/install/

### Installing Glide
For installing Glide you can visit this page to learn more about it:
http://kabiliravi.com/index.php/software/programming/golang/install-glide/

### Go workspace setup
Your go workspace directory should be under $GOPATH/src and $GOPATH will point to $HOME/go
```
$ mkdir $GOPATH/src/gotutorial
```
### Clone the helloworld git repo
Clone the helloworld git repository into gotutrial workspace folder by following commnad:
```
$ git clone https://github.com/gotutorial/helloworld.git
$ cd helloworld
```
### Build and run helloworld-basic
Build helloworld-basic by following command
```
$ cd $GOPATH/gototurial/helloworld/helloworld-basic
$ go build -o hello
```
Run **hello** executable file by following commnad
```
$ ./hello
Hello World
```
### Build and run helloworld-glide
Build helloworld-glide by following command
```
$ cd $GOPATH/gototurial/helloworld/helloworld-glide

$ glide install
[INFO]	Lock file (glide.lock) does not exist. Performing update.
[INFO]	Downloading dependencies. Please wait...
[INFO]	--> Fetching updates for github.com/dustin/go-humanize
[INFO]	Resolving imports
[INFO]	Downloading dependencies. Please wait...
[INFO]	Setting references for remaining imports
[INFO]	Exporting resolved dependencies...
[INFO]	--> Exporting github.com/dustin/go-humanize
[INFO]	Replacing existing vendor dependencies
[INFO]	Project relies on 1 dependencies.

$ go build -o hello
```
Run **hello** executable file by following commnad
```
$ ./hello
Hello World !!
Sample File Size in MB: 48 MB.
```

## See Also

You can learn more by visiting following pages in my website

### Write your first application
Find the tutoral of how I wrote these **helloworld** applications here:
http://kabiliravi.com/index.php/software/programming/golang/helloworld/

### Write your first application using Glide
Find the tutoral of how I wrote these **helloworld using Glide** applications here:
http://kabiliravi.com/index.php/software/programming/golang/helloworld-using-glide/
# weather-cli
