# getmyip
gitmyip is a command-line alternative to opening a browser to use sites such 
as whatismyip.com to get your external IP address, it works in a similar 
fashion by getting your external IP address by doing a web request to 
https://ip.5ec.nl (OR https://ip4.5ec.nl OR https://ip6.5ec.nl). 

getmyip is written in [Go](https://golang.org/), prebuilt binaries
are available for Linux and Windows from the 
[releases section](https://github.com/mrkcor/getmyip/releases). If you are using
another platform you can download the sources and build it yourself using:

```
go build getmyip.go
```

Once you have a compiled version copy it to a convenient location from where you
can then execute it.

Use the -h switch to display the usage instructions:

```
➜  ./getmyip -h
Usage: getmyip [options]
     -4 Get your external IPv4 address
     -6 Get your external IPv6 address
     -v Verbose output
``` 