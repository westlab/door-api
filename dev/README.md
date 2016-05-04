Building interop2016 developing virtual environment stack and deploying Application using Ansible Playbooks and Vagrant.
-------------------------------------------

**CAUTION:** This doc does not support Windows system.

## Preparing
### Install VirtualBox
See [Official Page](https://www.virtualbox.org/wiki/Downloads). Donwload and Install VB.
### Install Vagrant
See [Official Docs INSTALL VAGRANT](https://www.vagrantup.com/docs/installation/)
### Add CentOS7 box

```
vagrant box add contos/7
```
### Install Ansible
#### If you have pip command, (MacOSX has pip from the beginning.)
```
sudo pip install ansible
```
#### Other, on CentOS, see [Latest Release Via Yum ](http://docs.ansible.com/ansible/intro_installation.html#latest-release-via-yum)  
#### on Ubuntu, see [Latest Release Via Ap (Ubuntu)](http://docs.ansible.com/ansible/intro_installation.html#latest-releases-via-apt-ubuntu)


## Main

### Chcek you're in `door-api/dev`
### Start Vagrant
```
vagrant up --provision
```
**CAUSION: You always have to start `vagrant up/reload` with `--provision` because of MariaDB issue...**

Installing centos7 and provisioning will be started.  
**You'll be required to input password. Type `vagrant`**

### SSH to devserver
```
vagrant ssh devserver
```

( following commands are executed in devserver )
### Go get/install glide
```
go get github.com/Masterminds/glide
go install github.com/Masterminds/glide
```
### Check you're in `/usr/local/go/src/github.com/westlab/door-api`
(You can see door-api components because of Vagrant synced function.)
### Run
**See [door-api how to run](https://github.com/westlab/door-api#how-to-run)**

### On your host, check it is working.
```
curl http://192.168.33.41:8080/v1/word_rank
```
(`192.168.33.41` is devserver's default IP.)

### **You can edit files in your host thanks to Vagrant synced function!!**

## Trouble Shooting
### Case1: Check MariaDB is working without problems
##### on devserer
```
sudo systemctl status mariadb
```
if mariadb is not working
##### on host
```
vagrant reload --provision
```
