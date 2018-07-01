Gocker (Go + Docker)

This is a bash wrapper for turning Golang's official docker image 
into a container that acts like the native command line compiler.

Performs the following:
-mounts host GOPATH dir to container's /go folder
-adds a tmpfs file system in rwx mode
-sets up container's environment variables for Go
-drops container user to UID=1000 for security purposes
-reflects current working directory in container FS
-destroys container after running

To use:
1) Download and install DockerCE
2) Ensure docker daemon is running
3) Change first variable "godir" to your host's relative GOPATH dir
(ex. if GOPATH is /home/user/code/go, then godir='code/go')

Usage examples:
go version
go build project
go run main.go

Enjoy!
