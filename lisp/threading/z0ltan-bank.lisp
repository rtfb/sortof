;;; Modify a shared resource from multiple threads

(defclass bank-account ()
  ((id :initarg :id
       :initform (error "id required")
       :accessor :id)
   (name :initarg :name
         :initform (error "name required")
         :accessor :name)
   (balance :initarg :balance
            :initform 0
            :accessor :balance)))

(defgeneric deposit (account amount)
  (:documentation "Deposit money into the account"))

(defgeneric withdraw (account amount)
  (:documentation "Withdraw amount from account"))

(defmethod deposit ((account bank-account) (amount real))
  (incf (:balance account) amount))

(defmethod withdraw ((account bank-account) (amount real))
  (decf (:balance account) amount))

(defparameter *rich*
  (make-instance 'bank-account
                 :id 1
                 :name "Rich"
                 :balance 0))

(defun demo-race-condition ()
  (loop repeat 100
     do
       (sb-thread:make-thread
        (lambda ()
          (loop repeat 10000 do (deposit *rich* 100))
          (loop repeat 10000 do (withdraw *rich* 100))))))

(defvar *lock* (sb-thread:make-mutex))

(defun demo-race-condition-locks ()
  (loop repeat 100
     do
       (sb-thread:make-thread
        (lambda ()
          (loop repeat 10000 do (sb-thread:with-mutex (*lock*)
                                  (deposit *rich* 100)))
          (loop repeat 10000 do (sb-thread:with-mutex (*lock*)
                                  (withdraw *rich* 100)))))))

