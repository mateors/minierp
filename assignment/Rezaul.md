## Fyne installation on Windows 
* step 1 [`Install Go`](https://youtu.be/hffMABwkW00)
* step 2 `Install Gcc`
> Install `MSYS2` from [msys2.org](https://www.msys2.org/) 


![image](https://user-images.githubusercontent.com/77927449/146638508-f40647d3-6530-445a-afec-333e72ec3e73.png)
![image](https://github.com/irezaul/minierp-1/blob/main/assignment/fyne_install/2.jpg?raw=true)

* step 3 - after install `msys2` Open “MSYS2 MinGW 64-bit” from start menu..
* Step 4 - use this command on `MinGW Terminal`
```bash
pacman -Syu
```
![image](https://github.com/irezaul/minierp-1/blob/main/assignment/fyne_install/3.jpg?raw=true)
> press `y` to install

![image](https://github.com/irezaul/minierp-1/blob/main/assignment/fyne_install/3-1.jpg?raw=true)

### step 5 - Another command use on MinGW Terminal
```bash
pacman -Su
```
![image](https://github.com/irezaul/minierp-1/blob/main/assignment/fyne_install/4.jpg?raw=true)
### step 6 - next command use this ...
```bash
pacman -S --needed base-devel mingw-w64-x85_64-toolchain
```
![image](https://github.com/irezaul/minierp-1/blob/main/assignment/fyne_install/5.jpg?raw=true)
### `importent` :=  Add `C:\msys64\mingw64\bin` to `PATH` in `User Variables` and in `System Variables`
![image](https://github.com/irezaul/minierp-1/blob/main/assignment/fyne_install/7.jpg?raw=true)

### * now your `GCC` ready for play with `FYNE`.
