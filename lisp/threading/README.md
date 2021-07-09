
`sudo apt install sbcl`

Then install quicklisp by following instructions here:
https://www.quicklisp.org/beta/

And that should be it:

```
CL-USER> (ql:quickload :bt-semaphore)
To load "bt-semaphore":
  Load 1 ASDF system:
    bt-semaphore
; Loading "bt-semaphore"

(:BT-SEMAPHORE)

CL-USER> bt:*supports-threads-p*
T
```

NOTE! The above will not work on a Mac since SBCL does not support threads there
out of the box. It needs compiling from source, as described
[here](https://z0ltan.wordpress.com/tag/sbcl-thread-support-mac-os/), but I
didn't try that.

