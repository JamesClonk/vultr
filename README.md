# vultr
[Vultr](https://www.vultr.com) CLI and API client library, written in [Go](https://golang.org)

[![GoDoc](https://godoc.org/github.com/JamesClonk/vultr/lib?status.png)](https://godoc.org/github.com/JamesClonk/vultr/lib) [![Build Status](https://travis-ci.org/JamesClonk/vultr.png?branch=master)](https://travis-ci.org/JamesClonk/vultr)

### Screenshot

![Screenshot](https://github.com/JamesClonk/vultr/raw/master/screenshot.png "Screenshot")

Everybody likes screenshots, even of command line tools.. :smile:

### Usage

Here a some usage examples:

---

##### show help text for a command
```sh
$ vultr snapshot --help
```
```
Usage: vultr snapshot  COMMAND [arg...]

modify snapshots

Commands:
  create       create a snapshot from an existing virtual machine
  delete       delete a snapshot
  list         list all snapshots on current account

Run 'vultr snapshot COMMAND --help' for more information on a command
```

---

##### list available plans for region
```sh
$ vultr plans -r 9
```
```
VPSPLANID NAME                                VCPU  RAM   DISK  BANDWIDTH PRICE
30        1024 MB RAM,20 GB SSD,2.00 TB BW    1     1024  20    2.00      7.00
29        768 MB RAM,15 GB SSD,1.00 TB BW     1     768   15    1.00      5.00
3         2048 MB RAM,40 GB SSD,3.00 TB BW    2     2048  40    3.00      15.00
28        8192 MB RAM,120 GB SSD,5.00 TB BW   4     8192  120   5.00      70.00
27        4096 MB RAM,65 GB SSD,4.00 TB BW    2     4096  65    4.00      35.00
```

---

##### add SSH public key
```sh
$ vultr sshkey create -n sampleKey --key="$(cat ~/.ssh/id_rsa.pub)"
```
```
SSH key create success!

SSHKEYID        NAME        KEY
24c81f53be692   sampleKey   ssh-rsa AAAB3NzaC1yc2EQABAQClpsNAM+huOB2dpxM..
```

---

##### create new virtual machine
```sh
$ vultr server create -n "test-server" -r 9 -p 29 -o 127
```
```
Virtual machine create success!

SUBID           NAME            DCID    VPSPLANID       OSID
1685097         test-server     9       29              127
```

---

##### show information about virtual machine
```sh
$ vultr server show 1685097
```
```
Id (SUBID):             1685097
Name:                   test-server
Operating system:       CentOS 6 x64
Status:                 active
Power status:           running
Location:               Frankfurt
Region (DCID):          9
VCPU count:             1
RAM:                    768 MB
Disk:                   Virtual 15 GB
Allowed bandwidth:      1000
Current bandwidth:      0
Cost per month:         5.00
Pending charges:        0.01
Plan (VPSPLANID):       29
IP:                     107.62.131.240
Netmask:                255.255.254.0
Gateway:                107.62.131.1
Internal IP:
IPv6 IP:
IPv6 Network:           ::
IPv6 Network Size:      0
Created date:           2015-02-08 12:36:36
Default password:       sbiecxo8yk!5
Auto backups:           no
KVM URL:                https://my.vultr.com/subs/vps/novnc/api.php?data=ILXS..
```

---

#### TODO

:scream:
- [x] ~~add options (ServerOption struct) to server create command~~
- [x] ~~implement all server subcommands~~
  - [x] ~~server rename~~
  - [x] ~~server start~~
  - [x] ~~server halt~~
  - [x] ~~server reboot~~
  - [x] ~~server reinstall~~
  - [x] ~~server change-os~~
    - [x] ~~server change-os --list (to show possible choices)~~
  - [x] ~~server bandwidth~~
- [x] ~~implement ssh command~~
- [ ] implement snapshot commands
- [x] ~~implement script commands~~
- [ ] Vultr auth
  - [ ] read apiKey (and other stuff like default OS image, default region, etc..) from ~/.vultr.json file if present
  - [ ] command line arguments like --api-key take precendence over values from ~/.vultr.json though!
  - [ ] add "vultr auth" command, which prompts for apiKey, then stores it into ~/.vultr.json
- [ ] add usage guide for command line tool
- [ ] add documentation on how to use library in other projects
- [ ] create github.io page

