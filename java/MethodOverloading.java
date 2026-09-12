public class MethodOverloading {
    public static void main(String[] args) {
        int n1 = 20;
        int n2 = 30;
        int n3 = 40;

        Calculatez calsz = new Calculatez();
        
        int result1 = calsz.add(n1, n2);
        int result2 = calsz.add(n1, n2, n3);

        System.out.println(result1);
        System.out.println(result2);
    }
}

class Calculatez {
    public int add(int a1, int a2) {
        return a1 + a2;
    }

    public int add(int a1, int a2, int a3) {
        return a1 + a2 + a3;
    }
}