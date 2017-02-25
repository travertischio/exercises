;;; Fizzbuzz

(defun fizzbuzz (n)
  (cond
    ((= n 1) (cons n nil))
    ((= (mod n 3) 0)
      (cond
        ((= (mod n 5) 0) (cons  "FizzBuzz" (fizzbuzz (- n 1))))
        ((cons "Fizz" (fizzbuzz (- n 1))))))
    ((= (mod n 5) 0) (cons "Buzz" (fizzbuzz (- n 1))))
    ((cons n (fizzbuzz (- n 1))))))

(print (reverse (fizzbuzz 100)))
