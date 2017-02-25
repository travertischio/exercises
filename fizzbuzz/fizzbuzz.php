// Fizzbuzz

<?php
  for ($x = 1; $x <= 100; $x++) {
    if ($x % 3 == 0 && $x % 5 == 0) {
      echo "FizzBuzz\n";
    } elseif ($x % 3 == 0) {
      echo "Fizz\n";
    } elseif ($x % 5 == 0) {
      echo "Buzz\n";
    } else {
      echo "$x\n";
    }
  }
?>
