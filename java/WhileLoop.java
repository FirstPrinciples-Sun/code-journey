public class WhileLoop {
    public static void main(String a[]) {
        int i = 1;

        while(i <= 4) {
            
            System.err.println("Hi" + " " + i);
            int j = 1;
            while (j <= 4) {
                System.err.println("Hello" + " " + j);

                j++;
            }

            i++;
        }

        System.out.println("Bye" + " " + i);
    }
    
}
