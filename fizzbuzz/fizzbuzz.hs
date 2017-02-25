-- fizzbuzz

main = do
  mapM_ fizzbuzz [1..100]

fizzbuzz x
  | x `mod` 3 == 0 && x `mod` 5 == 0 = print "FizzBuzz"
  | x `mod` 3 == 0 = print "Fizz"
  | x `mod` 5 == 0 = print "Buzz"
  | otherwise = print $ show x
