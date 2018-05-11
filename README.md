# PsOff
### _A command-line tool for Mac and Linux to clear your ports_

Ever try to launch a process and are blocked by an error like...

```Error: listen EADDRINUSE 127.0.0.1:8080``` ?

With PsOff you can tell that process to get out of the way:

```sh
> psoff 8080                                             0|21:07:39
Process 41945: [ node ] is running on port 8080
Would you like to kill it? [Y/n] Y
The process has been killed without mercy!
```

To install do the following:

#### Direct binary installation (Mac only):

just enter the following into a bash terminal:

```git clone https://github.com/stantronic/psoff.git && cd psoff && ./install.sh```

#### Build and install (Mac and Linux):
_requires that you have go installed_

Enter the following into a bash terminal:

```git clone https://github.com/stantronic/psoff.git && cd psoff && ./build_and_install.sh```
