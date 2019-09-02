public class Dead {

  static void SLEEP(int msec) {
    try {
      Thread.sleep(msec);
    } catch (Exception e) {
      //TODO: handle exception
    }
  }

  static void print(String msg) {
    System.out.println(msg);
  }

  static void goDeadLock() {
    final Object o1 = new Object();
    final Object o2 = new Object();

    Thread t1 = new Thread("t1") {
      public void run() {
        synchronized(o1) {
          SLEEP(1000);
          synchronized(o2) {
            print("T1");
          }
        }
      }
    };
    Thread t2 = new Thread("t2") {
      public void run() {
        synchronized(o2) {
          SLEEP(1000);
          synchronized(o1) {
            print("T2");
          }
        }
      }
    };
    t1.start();
    t2.start();
    print("ALL START");

    //SLEEP(2000);
    t1.interrupt(); t2.interrupt();
    JOIN(t1);
    JOIN(t2);
  }

  public static void JOIN(Thread t) {
    try {
      t.join();
    } catch (Exception e) {
      print(e.toString());
    }
  }

  public static void main(String[] args) {
    print("HELLO");
    goDeadLock();
  }
}