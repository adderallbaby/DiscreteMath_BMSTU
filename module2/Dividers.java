import java.util.Arrays;
import java.util.Scanner;


public class Dividers {
    public static long exist(long[] mas, long x){
        for (long ma : mas) {
            if (ma == x) {
                return x;
            }
        }
        return -1;
    }
    static void reverseArray(long[] intArray, int size) {
        int i, k;
        long temp;
        for (i = 0; i < size / 2; i++) {
            temp = intArray[i];
            intArray[i] = intArray[size - i - 1];
            intArray[size - i - 1] = temp;
        }
    }
        public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        long n = scanner.nextLong();

        if(n == 1){
            System.out.println("graph {");
            System.out.println("\t" + 1 );
            System.out.println("}");
        }else {


            long[] allDividers = new long[1000000];
            int limit = (int) Math.sqrt(n);
            int counter = 0;
            for (long i = 1; i <= limit; i++) {
                if (n % i == 0) {
                    allDividers[counter] = i;
                    counter++;
                    if (exist(allDividers,  n / i) == -1) {
                        allDividers[counter] = n / i;
                        counter++;
                    }
                }
            }


            long[] newDiv = new long[counter];
            for (long i = 0; i < counter; i++) {

                newDiv[Math.toIntExact(i)] = allDividers[Math.toIntExact(i)];
            }

            Arrays.sort(newDiv);
            reverseArray(newDiv, counter);
            System.out.println("graph {");
            for (int i = 0; i < counter; i++) {
                System.out.println("\t" + newDiv[i]);
            }
            if (counter == 1) {
                System.out.println(n + " -- " + 1);
            } else {
                for (int u = 0; u < counter; u++) {
                    for (int v = u + 1; v < counter; v++) {
                        if (newDiv[u] % newDiv[v] == 0) {
                            int flag = 1;
                            for (int w = u + 1; w < v; w++) {
                                if (newDiv[u] % newDiv[w] == 0 && newDiv[w] % newDiv[v] == 0) {
                                    flag = 0;
                                    break;
                                }
                            }
                            if (flag == 1) {
                                System.out.println("\t" +newDiv[u] + " -- " + newDiv[v]);


                            }
                        }
                    }
                }
            }
            System.out.println("}");
        }
    }

}


