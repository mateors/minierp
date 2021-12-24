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

# Requirment 
==================

![image](https://github.com/irezaul/minierp-1/blob/main/assignment/images/Screenshot%20from%202021-12-24%2020-33-42.png?raw=true)

order status: pending
order status: in process/working
order status: ready with fix
order status: ready without fix
order status: delivered

---------------------------------------------------------------------
new client
=======================
Name: Rahim
Mobile:
Email*
Address:
Problem Details: 
========================

Delivery date: 1 days, 2 days 3 days, 4 days
Warranty: 6 months

---------------------------------------------------------------------
sms ==> Name*ProblemDetails
sms ==> 12121212

Dear Rahim
Your order no is: 12121212
Your expected delivery date is: 30 november 2021

-------------------
update: ready
Dear Rahim
Your order no is: 12121212
Your item/device is ready for delivery please collect from our store.
-------------------

update: delivered

Dear Rahim
Your order no is: 12121212 order status: delivered.
Warranty upto: 23 May 2022
---------------------------------------------------------------------


sms software -> incoming+outgoing
android sms so
topup software
hosting software+domain
email sender


