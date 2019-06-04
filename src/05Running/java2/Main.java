// 28032
// 89436

public class Main {

    public static void main(String[] args) throws InterruptedException{
        for (int i = 0; i < 1000000; i++) {
            MyThread myThread = new MyThread();
            myThread.start();
        }
    }
}

