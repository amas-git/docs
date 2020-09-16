import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;
import java.util.Random;

class ListNode {
  int value;
  ListNode next;
  public ListNode(int value) {
    this.value = value;
    this.next = null;
  }

  public ListNode append(int value) {
      last().next = new ListNode(value);
      return this;
  }

  public ListNode last() {
    ListNode last = this;
    while (last.next != null) {
      last = last.next;
    }
    return last;
  }
  // 1,2,3,4,5
  /**
   * nth[-1] : 倒数第一
   * nth[-2] : 倒数第二
   * @param n
   * @return
   */
  public ListNode nth(int n) {
    ListNode r = this;
    int i = n;

    for (; r.next!=null && i>1; i--,r=r.next) {
        System.out.println("x="+r.value);
    }

    return  r;
  }


  public boolean isLast() {
    return this.next == null;
  }

  static ListNode fromArray(int[] xs) {
    ListNode root = new ListNode(xs[0]);
    for(int i=1; i<xs.length; ++i) {
      root.append(xs[i]);
    }
    return root;
  }

  public ArrayList<Integer> toArray() {
    ArrayList<Integer> xs = new ArrayList<Integer>();
    for (ListNode n = this; n != null; n = n.next) {
      xs.add(n.value);
    }
    return xs;
  }

  @Override
  public String toString() {
    return toArray().toString();
  }
}

public class Hello {
  public static void main(String[] args) {
    
    //List<Integer> xs = Arrays.asList(1,2,3,4,5,6);
    //System.out.println(Arrays.toString(suffle(xs).toArray()));

    //Collections.shuffle(arg0, arg1);

    ListNode xs = ListNode.fromArray(new int[] {1,2,3,4,5,6});
    System.out.println(xs);
    
    // System.out.println(xs.nth(1).value);
    // System.out.println(xs.nth(3).value);
    // System.out.println(xs.nth(6).value);
    //System.out.println(xs.nth(7).value);
    
    // System.out.println(xs.nth(-3).value);
    System.out.println(xs.nth(-1).value);
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