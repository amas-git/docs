import java.util.Arrays;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;
import java.util.Random;

public class Hello {
  public static void main(String[] args) {
    
    List<Integer> xs = Arrays.asList(1,2,3,4,5,6);
    System.out.println(Arrays.toString(suffle(xs).toArray()));
  }


  static Random random = new Random(System.currentTimeMillis()); 
  public static<T> List<T> suffle(List<T> xs) {
    Collections.sort(xs, new Comparator<T>() {
      @Override
      public int compare(T a, T b) { 
        return random.nextInt(3) - 1;
      }
    });
    return xs;
  }
}