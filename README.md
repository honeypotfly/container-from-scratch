# Container From Scratch
Using Go to build a container from scratch.

### Introduction
If you ran

```
docker run --rm -it ubuntu /bin/bash
```
You would be starting up a container that uses the ubuntu image and spins up a shell. It would be an interactive container (-it) and it would be removed once the container is exited (--rm). If you were to look at the processes in this interactive container, it would look like this:

```
root@97256822e3f1:/# ps
    PID TTY          TIME CMD
      1 pts/0    00:00:00 bash
      9 pts/0    00:00:00 ps
```

Notice the low process ids and and the new hostname. This in essence is what signifies a container, and we're going to implement it in Go.

### The Objective
Typically, in docker we would run something like
```
docker run image <cmd> <params>
```

Well, in this version, we'll need something similar:
```
go run main.go run <cmd> <params>
```

### What is a Namespace?
A namespace is where we limit what a process can see. So our docker container could not see the processes on the host, this is bc the docker container has a namespace for process IDs. Similarly, the container could only see its own hostname, thats because it it had a namespace for its hostname.

We can set up namespaces using syscalls. This is a big part of containerization, since it is what allows us to restrict what the container can see and access.


### Credits
This project was presented by Liz Rice at GOTO 2018. I highly recommend [watching it](https://www.youtube.com/watch?v=8fi7uSYlOdc).
