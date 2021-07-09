# dse
dse
Distributed Search Engine (DSE), stage1

DSE project. Fuller description available at https://rorur.com.


## Installation.

At the moment, only Linux is supported. 

To install, add

```
deb [trusted=yes] https://rorur.com/debian ./
```

to 
```
/etc/apt/sources.list
```
file and then run

```
sudo apt update
```
and

```
sudo apt install dse
```
(this installs dse manager script). Then run
```
sudo dse install
```
which will install binaries.

Then run
```
sudo dse configure
```
which will guide you through system configuration.

Then run
```
sudo dse start
```

which will create DSE daemon. This is a systemd daemon. By default, it is added to autorestart and automatic start on boot.

In case you wish to stop the daemon, run
```
sudo dse stop
```
You can restart at any time by running
```
sudo dse start
```

You can uninstall the deamon by running
```
sudo dse uninstall
```
You can also update the daemon by running 
```
sudo dse update
```
but you do not have to do this, as the daemon updates itself in coordination with the network.


To change configuration, run 
```
sudo dse configure 
```

again
