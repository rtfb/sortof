
(defparameter *address* #(0 0 0 0))
(defparameter *port* 8080)

; Read data from a connection.
(defun read-from-connection (connection)
  (multiple-value-bind (buffer length) (sb-bsd-sockets:socket-receive connection nil 1024)
    (let (data)
      (if (> length 0)
          (subseq buffer 0 length)
          data))))

; Handle an incoming connection: read the data and echo it back, but replace
; the second character with 'z'.
(defun handle-connection (connection)
  (let ((data (read-from-connection connection)))
    (setf (aref data 1) #\z)
    (sb-bsd-sockets:socket-send connection data nil)
    (sb-bsd-sockets:socket-close connection)))

(defun start-server ()
  (let ((server (make-instance 'sb-bsd-sockets:inet-socket :type :stream :protocol :tcp)))
    ; Set socket options
    (setf (sb-bsd-sockets:sockopt-reuse-address server) t)
    (setf (sb-bsd-sockets:non-blocking-mode server) t)

    ; Bind to an address
    (sb-bsd-sockets:socket-bind server *address* *port*)
    (sb-bsd-sockets:socket-listen server 1)
    (loop
      (let ((connection (sb-bsd-sockets:socket-accept server)))
        (when connection
          (handle-connection connection))))))

(start-server)
