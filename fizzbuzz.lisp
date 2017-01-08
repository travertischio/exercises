;;; Fizzbuzz

(defun fizzbuzz (n)
  (cond
    ((= (mod n 3) 0)
      (cond
        ((= (mod n 5) 0) (print "FizzBuzz"))
        ((print "Fizz"))))
    ((= (mod n 5) 0) (print "Buzz"))
    ((print n)))
  (cond
    ((< n 100) (fizzbuzz (+ n 1)))))

(fizzbuzz 1)
