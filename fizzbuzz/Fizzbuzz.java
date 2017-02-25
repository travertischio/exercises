// Fizzbuzz

public class Fizzbuzz {

  public static void main(String[] args) {
    for(int x = 1; x <= 100; x++) {
      if (x % 3 == 0 && x % 5 == 0) {
        System.out.print("Fizzbuzz\n");
      } else if (x % 3 == 0) {
        System.out.print("Fizz\n");
      } else if (x % 5 == 0) {
        System.out.print("Buzz\n");
      } else {
        System.out.print(x + "\n");
      }
    }
  }
}
