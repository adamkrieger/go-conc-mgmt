public class MyThread extends Thread {
    public void run() {
        while(true){
                try {
                    Thread.sleep(5000);
                } catch(InterruptedException ex){
                }
            }
    }
}