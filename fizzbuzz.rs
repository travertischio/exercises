// fizzbuzz

fn main() {
  for x in 1..101 {
    let out =
      if x%3 == 0 && x%5 == 0 {
        "FizzBuzz"
      } else if x%3 == 0 {
        "Fizz"
      } else if x%5 == 0 {
        "Buzz"
      } else {
        "num" // x.to_string();
      };

    if out == "num" {
      println!("{}", x);
    } else {
      println!("{}", out);
    }
  }
}
