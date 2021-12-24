This file contains some of useful go and fyne commands.

Locate Go build cache directory

The default location for build cache data is a directory named go-build in the standard cache directory. For me, on Linux, it is ~/.cache/go-build/. You can find out yours using this command
//
go env GOCACHE
//
Go also conveniently allows you to modify the cache location using the environment variable


Check disk usage of Go build cache
Here's a quick one-liner to see the disk usage of the build cache
//
du -hs $(go env GOCACHE)
//


Finally, to address the question of this article, you can run this command
//
go clean --cache
//
This command removes all the subdirectories inside go-build directory and leaves out just two files
 
                 ** How To Run Exe without having cmd to launch every time **
 
 Windows applications load from a command prompt by default, which means if you click an icon you may see a command window. while opening exe that was created using GO, To fix this add the parameters" -ldflags -H=windowsgui" to your run or build commands.
command Examples :
go build -ldflags -H=windowsgui "your_go_file_name".go
go run -ldflags -H=windowsgui "your_go_file_name".go
// in my case my go file name is "hello.go " so I will use these commands like this 
go build -ldflags -H=windowsgui hello.go
go run -ldflags -H=windowsgui hello.go

