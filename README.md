# This is a simple golang program to build a TCP connection

## Let's start it
download it to your gopath(inside src) and switch to its location.
### 1、`go build` --compile and generate an executable file
### 2、`./TCP_golang -c server` --start server in a terminal
### 3、`./TCP_golang -c client` --start client in another terminal
### 4、`./TCP_golang -m` or `sudo ./TCP_golang -m` --start a monitor program, usually you need to use root authority 
### 5、client terminal can input some strings to server.

## more
### 1、update the package in main.go, or you can't compile it successfully
### 2、default parameters: 
     - server: ip 127.0.0.1 port: 8087
     - client: you can't set ip and port, you need server_ip and server_port to connect it
     - network_card: eth0 or ens33, you can execute `ifconfig` to check your ip and network_card.
### 3、use 127.0.0.1 localhost, you can start a TCP connection easily.
        