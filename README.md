About
---
TCP port scanner written in go from the book 'Black Hat Go'. Hard coded to scan scanme.nmap.com
The implementation uses a worker pool with 100 channels. Returns a list of open ports


Build
----
`make build`

binary is in bin 
 
 
 
 To Do
----
- [ ] allow users to specify number of workers as an option
- [ ] allow url to be passed as argument
- [ ] ability to parse port strings like `80, 443, 8080` similar to nmap

Enhancements
---

It's not neccessary to send on the _results_ channel for every port scanned. We can
use an additional channel not only to track the workers but also to prevent a race condition 
by ensuring the completion of all gathered results.

check [Here](https://github.com/blackhat-go/bhg/blob/master/ch-2/scanner-port-format)
