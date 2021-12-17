## How to install Fyne on your machine

* https://github.com/msys2/msys2-installer/releases/download/2021-11-30/msys2-x86_64-20211130.exe

## After C Compiler installation
> open git bash
> if you dont know what a git bash is follow this link  -> https://git-scm.com/


  * $ pacman -Syu
  * $ pacman -S git mingw-w64-x86_64-toolchain
  
  * echo "export PATH=$PATH:/c/Program\ Files/Go/bin:~/Go/bin" >> ~/.bashrc
  
  
  * go mod init minerp
  
  * go get fyne.io/fyne/v2
  
  
  * go run fyne.io/fyne/v2/cmd/fyne_demo
  
  * go install fyne.io/fyne/v2/cmd/fyne_demo@latest
