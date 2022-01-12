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

### if show error this then produce next command
![toolchain_error](images/toolchain_error.png)

### If you want MSYS2's MinGW-w64 x86_64 GCC then you should install the `mingw-w64-x86_64-toolchain` package group via:

```bash
pacman -S mingw-w64-x86_64-toolchain
```
![toolchain_error_fixed](images/toolchain_error_fixed.png)
### `importent` :=  Add `C:\msys64\mingw64\bin` to `PATH` in `User Variables` and in `System Variables`
![image](images/7.jpg)

### * now your `GCC` ready for play with `FYNE`.
* Step 7 - For Fyne Download
```bash
go get fyne.io/fyne/v2
```
* Step 8 - To finish your module’s set up, you now need to tidy the module file to correctly reference Fyne as a dependency. You do this by using the following command (can be skipped if you are not using modules): 





```bash
go mod tidy
```

### Run the demo
If you want to see the Fyne toolkit in action before you start to code your own application, you can see our demo app running on your computer by executing:

```bash
go run fyne.io/fyne/v2/cmd/fyne_demo
```

