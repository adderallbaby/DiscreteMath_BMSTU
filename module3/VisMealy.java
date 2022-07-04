import java.util.Scanner;

public class VisMealy {
    public static void outPut(String[][] exit, int[][] crossover, int start){
        System.out.println("digraph {");
        System.out.println("    rankdir = LR");
        for (int i = 0; i < crossover.length; i++) {
            for (int j = 0; j < crossover[i].length; j++) {
                if (!(exit[i][j]==null)){
                    System.out.print("    "+i + " -> " + +crossover[i][j] + "[label = \"" + (char) (j + 97) + "(" + (exit[i][j]) + ")" + "\"]\n");
                }
            }
        }
        System.out.println("}");
    }
    public static void main(String[] args) {
        int n,m,start,cr;
        String sym;
        Scanner scanner = new Scanner(System.in);

        n = scanner.nextInt();
        m = scanner.nextInt();
        start = scanner.nextInt();
        int[][] crossover = new int[2*n][m];
        String[][] exit = new String[2*n][m];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                crossover[i][j] =scanner.nextInt();
            }
        }
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                exit[i][j] =scanner.next();

            }
        }
        outPut(exit,crossover,start);
    }

}
