# Fizzbuzz

# Strict and warnings are recommended.
use strict;
use warnings;

for( $a = 1; $a <= 100; $a++ ){
  if ($a % 3 == 0 && $a % 5 == 0) {
    print "FizzBuzz\n";
  } elsif ($a % 3 == 0) {
    print "Fizz\n";
  } elsif ($a % 5 == 0) {
    print "Buzz\n";
  } else {
    print "$a\n";
  }
}
