;;; Fizzbuzz

(defun fizzbuzz (n)
  (cond
    ((= n 1) (write-to-string 1))
    ((= (mod n 3) 0)
      (cond
        ((= (mod n 5) 0) (concatenate 'string "FizzBuzz" (fizzbuzz (- n 1))))
        ((concatenate 'string "Fizz" (fizzbuzz (- n 1))))))
    ((= (mod n 5) 0) (concatenate 'string "Buzz" (fizzbuzz (- n 1))))
    ((concatenate 'string (write-to-string n) (fizzbuzz (- n 1))))))

(print (fizzbuzz 100))
