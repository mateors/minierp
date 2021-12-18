## Fyne installation on Windows 
* step 1 [`Install Go`](https://youtu.be/hffMABwkW00)
* step 2 `Install Gcc`
> Install `MSYS2` from [msys2.org](https://www.msys2.org/) 


![image](images/1.jpg)
![image](images/2.jpg)

* step 3 - after install `msys2` Open “MSYS2 MinGW 64-bit” from start menu..
* Step 4 - use this command on `MinGW Terminal`
```bash
pacman -Syu
```
![image](images/3.jpg)
> press `y` to install

![image](images/3-1.jpg)

### step 5 - Another command use on MinGW Terminal
```bash
pacman -Su
```
![image](images/4.jpg)
### step 6 - next command use this ...
```bash
pacman -S --needed base-devel mingw-w64-x85_64-toolchain
```
![image](images/5.jpg)
### `importent` :=  Add `C:\msys64\mingw64\bin` to `PATH` in `User Variables` and in `System Variables`
![image](images/7.jpg)

### * now your `GCC` ready for play with `FYNE`.
