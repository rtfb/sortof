(asdf:oos 'asdf:load-op :FiveAM)
(defpackage :it.bese.FiveAM.example
  (:use :common-lisp
        :it.bese.FiveAM))
(in-package :it.bese.FiveAM.example)

(defun add-2 (n)
  (+ n 2))

(test add-2
 "Test the ADD-2 function" ;; a short description
 ;; the checks
 (is (= 2 (add-2 0)))
 (is (= 0 (add-2 -2))))

(run! 'add-2)
