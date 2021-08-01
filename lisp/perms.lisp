
(defun punchout (l skip)
  (loop for item in l
        for i from 0
        when (not (equal i skip))
        collect item into result
        finally (return result)))

(defun perms (l)
  (if (eq (length l) 2)
      (list l (reverse l))
      (let ((result '()))
        (dotimes (i (length l))
          (dolist (subperm (perms (punchout l i)))
            (push (push (nth i l) subperm) result)))
        result)))


(print (perms '(2)))
(print (perms '(3 4)))
(print (perms '(5 7 9)))

