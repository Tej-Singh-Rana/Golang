### Install `Go` in Linux system

- Download the archive file and extract it into `/usr/bin/` directory.
   - In case, if we are in different directory instead of home directory. Change with `cd` command.  
     ```sh
     $ cd /root/
     ```
   - Download archive with `wget` command.
     ```sh
     $ wget https://golang.org/dl/go1.15.8.linux-amd64.tar.gz
     --2021-02-13 10:37:13--  https://golang.org/dl/go1.15.8.linux-amd64.tar.gz
     ```
   - Extract it from `tar` command.
     ```sh
     $ pwd
     /root
     
     $ ls
     go1.15.8.linux-amd64.tar.gz
     
     $ tar -xvf go1.15.8.linux-amd64.tar.gz
     go/
     go/AUTHORS
     go/CONTRIBUTING.md
     go/CONTRIBUTORS
     go/LICENSE
     go/PATENTS
     go/README.md
     go/SECURITY.md
     go/VERSION
     go/api/
     ...
     
     $ ls
     go  go1.15.8.linux-amd64.tar.gz 
     ```
   - Copy/Move it extracted directory into `/usr/bin/` directory. 
     > In your case might be different, so please check before copy.
     ```sh
     $ pwd
     /root 
     
     $ cp -r go /usr/bin/
     ```
     > Note: To know the already existing `Go` path, use `which` command. Before installing a new version of `Go`, remove old one. Even you can manage multiple versions of `Go` in a same machine. <br>
     > [Instructions to manage multiple versions of `Go`](https://golang.org/doc/manage-install#installing-multiple)
     ```sh
     $ which go
     /usr/bin/go/bin/go
     ```
   - Set the __PATH__ in user's `$HOME/.bashrc` file. 
     ```sh
     $ echo 'export PATH="$PATH:/usr/bin/go/bin"' >> ~/.bashrc
     
     $ . ~/.bashrc
     ```
     > Error:- `go: cannot find GOROOT directory`. To fix this issue, set the `GOROOT` path. 
   - In my case, 
     ```sh
     $ echo 'export GOROOT="/usr/bin/go"' >> ~/.bashrc
     
     $ . ~/.bashrc
     ```
   - Check the `go` version.
     ```sh
     $ go version
     go version go1.15.8 linux/amd64
     ```
> Note :- __GOPATH__ - Your Go source code directory _{src, pkg, bin}_ and default for linux (`$HOME/go`), __GOROOT__ - Your Go installation path. 

- [GOROOT and GOPATH](https://www.jetbrains.com/help/go/configuring-goroot-and-gopath.html)
- [Small introduction of Root Directory in Linux](https://youtu.be/ucctAawSnj8)
- [Settings for Mac/Linux](https://learn.gopherguides.com/courses/preparing-your-environment-for-go-development/modules/setting-up-mac-linux/)

### Install `Go` in Windows system

- Download the MSI file and follow the given instructions. 
  > Note: By default, the installer will install `Go` to `Program Files` or `Program Files (x86)`. You can change that location if needed.
  ```cmd
  C:\Users\User>go version
  go version go1.15.8 windows/amd64
  ```
  
- [Instructions from Official page](https://golang.org/doc/install#install)
- [To set __Environment variables__ in Windows Machine](https://www.computerhope.com/issues/ch000549.htm)
- [Settings for Windows](https://learn.gopherguides.com/courses/preparing-your-environment-for-go-development/modules/setting-up-windows/)

