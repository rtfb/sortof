
(ql:quickload :bt-semaphore)

(defvar *lock* (bt:make-lock))

(defun make-thread-runner (top-level number)
  (lambda ()
    (let* ((curr-thread (bt:current-thread))
           (curr-thread-name (bt:thread-name curr-thread)))
      (dotimes (i number)
        (bt:with-lock-held (*lock*)
          (format top-level "[~a:~a] Counter: ~a~%" curr-thread-name curr-thread i))
        (sleep 1)))))

(defun do-counting-threads ()
  (let* ((top-level *standard-output*)
         (counting-thread-1 (bt:make-thread (make-thread-runner top-level 7)))
         (counting-thread-2 (bt:make-thread (make-thread-runner top-level 13)))
         )
    (bt:join-thread counting-thread-1)
    (bt:join-thread counting-thread-2)
    (format top-level "Done~%"))
  nil)

(do-counting-threads)
(quit)
