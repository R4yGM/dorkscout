![](https://api.travis-ci.org/R4yGM/dorkscout.svg)

# dorkscout
dokrscout is a tool to automate the finding of vulnerable applications or secret files around the internet throught google searches, dorkscout first starts by fetching the dorks lists from https://www.exploit-db.com/google-hacking-database and then it scans a given target or everything it founds

# Installation
dorkscout can be installed in different ways:

## **Go Packages**

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png" width=95 height=35>

 throught [Golang Packages](https://github.com/rust-lang/cargo) (golang package manager)
 
```bash
go get github.com/R4yGM/dorkscout
```
**this will work for every platform**

## **Docker**

<img src="https://cdn3.iconfinder.com/data/icons/logos-and-brands-adobe/512/97_Docker-512.png" width=35 height=35>

  if you don't have docker installed you can follow their [guide](https://docs.docker.com/engine/install/)
  
 first of all you have to pull the docker image (only **17.21 MB**) from the docker registry, you can see it [here](https://hub.docker.com/r/r4yan/dorkscout), if you don't want to pull the image you can also clone the repository and then build the image from the Dockerfile
 ```bash
docker pull r4yan/dorkscout:latest
  ```
 
  if you don't want to pull the image you can download or copy the dorkscout Dockerfile that can be found [here](https://github.com/R4yGM/dorkscout/blob/1.0/Dockerfile) and then build the image from the Dockerfile
  
  then if you want to launch the container you have to first create a volume to share your files to the container
  
  ```bash
  docker volume create --name dorkscout_data
  ``` 
 using docker when you launch the container it will automatically install the dork lists inside a directory called "dorkscout" :
   ```bash
-rw-r--r-- 1 r4yan r4yan   110 Jul 31 14:56  .dorkscout
-rw-r--r-- 1 r4yan r4yan 79312 Aug 10 20:30 'Advisories and Vulnerabilities.dorkscout'
-rw-r--r-- 1 r4yan r4yan  6352 Jul 31 14:56 'Error Messages.dorkscout'
-rw-r--r-- 1 r4yan r4yan 38448 Jul 31 14:56 'Files Containing Juicy Info.dorkscout'
-rw-r--r-- 1 r4yan r4yan 17110 Jul 31 14:56 'Files Containing Passwords.dorkscout'
-rw-r--r-- 1 r4yan r4yan  1879 Jul 31 14:56 'Files Containing Usernames.dorkscout'
-rw-r--r-- 1 r4yan r4yan  5398 Jul 31 14:56  Footholds.dorkscout
-rw-r--r-- 1 r4yan r4yan  5568 Jul 31 14:56 'Network or Vulnerability Data.dorkscout'
-rw-r--r-- 1 r4yan r4yan 49048 Jul 31 14:56 'Pages Containing Login Portals.dorkscout'
-rw-r--r-- 1 r4yan r4yan 16112 Jul 31 14:56 'Sensitive Directories.dorkscout'
-rw-r--r-- 1 r4yan r4yan   451 Jul 31 14:56 'Sensitive Online Shopping Info.dorkscout'
-rw-r--r-- 1 r4yan r4yan 29938 Jul 31 14:56 'Various Online Devices.dorkscout'
-rw-r--r-- 1 r4yan r4yan  2802 Jul 31 14:56 'Vulnerable Files.dorkscout'
-rw-r--r-- 1 r4yan r4yan  4925 Jul 31 14:56 'Vulnerable Servers.dorkscout'
-rw-r--r-- 1 r4yan r4yan  8145 Jul 31 14:56 'Web Server Detection.dorkscout'
  ```
  so that you don't have to install them
  then you can start scanning by doing :
  ```bash
docker run -v Dorkscout:/dorkscout r4yan/dorkscout scan <options>
  ```
  replace the `<options>` with the options/arguments you want to give to dorkscout,
  example :
   ```bash
docker run -v dorkscout_data:/dorkscout r4yan/dorkscout scan -d="/dorkscout/Sensitive Online Shopping Info.dorkscout" -H="/dorkscout/a.html"
  ```
  
  **Always save your results inside the volume and not in the container because then the results will be deleted! you can save them by writing the same volume path of the directory you are saving the results**
 
 if you added this and did everything correctly at the end of every scan you'd find the results inside the folder `/var/lib/docker/volumes/dorkscout_data/_data`
  
  
  **this will work for every platform**
  
  ## Executable
  you can also download the already compiled binaries [here](https://github.com/R4yGM/dorkscout/releases) and then execute them

# Usage

```bash
dorkscout -h
```
```bash
Usage:
  dorkscout [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  delete      deletes all the .dorkscout files inside a given directory
  help        Help about any command
  install     installs a list of dorks from exploit-db.com
  scan        scans a specific website or all the websites it founds for a list of dorks

Flags:
  -h, --help   help for dorkscout

Use "dorkscout [command] --help" for more information about a command.
```
to start scanning with a wordlist and a proxy that will then return the results in a HTML format 
```bash
dorkscout scan -d="/dorkscout/Sensitive Online Shopping Info.dorkscout" -H="/dorkscout/a.html" -x socks5://127.0.0.1:9050
```
results : 

![](https://i.imgur.com/hGz7b9n.png)

## Install wordlists
to start scanning you'll need some dork lists and to have these lists you can install them through the install command
```bash
dorkscout install --output-dir /dorks
```
and this will fetch all the available dorks from exploit.db
```
[+] ./Advisories and Vulnerabilities.dorkscout
[+] ./Vulnerable Files.dorkscout
[+] ./Files Containing Juicy Info.dorkscout
[+] ./Sensitive Online Shopping Info.dorkscout
[+] ./Files Containing Passwords.dorkscout
[+] ./Vulnerable Servers.dorkscout
[+] ./Various Online Devices.dorkscout
[+] ./Pages Containing Login Portals.dorkscout
[+] ./Footholds.dorkscout
[+] ./Error Messages.dorkscout
[+] ./Files Containing Usernames.dorkscout
[+] ./Network or Vulnerability Data.dorkscout
[+] ./.dorkscout
[+] ./Sensitive Directories.dorkscout
[+] ./Web Server Detection.dorkscout
2021/08/11 19:02:45 Installation finished in 2.007928 seconds on /dorks
```
