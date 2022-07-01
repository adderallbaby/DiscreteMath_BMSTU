import java.util.Scanner;

public class VisMealy {
    public static void outPut(char[][] exit, int[][] crossover, int start){
        System.out.println("digraph {");
        System.out.println("\trankdir = LR");
        for (int i = 0; i < crossover.length; i++) {
            for (int j = 0; j < crossover[i].length; j++) {
                if (crossover[i][j] != 0) {
                    System.out.print("\t" +i + "->" + +crossover[i][j] + "[label = \"" + (char) (j + 97) + "(" + (exit[i][j]) + ")" + "\"]\n");
                }
            }
        }
        System.out.println("}");
    }
    public static void main(String[] args) {
        int n,m,start,cr;
        char sym;
        Scanner scanner = new Scanner(System.in);

        n = scanner.nextInt();
        m = scanner.nextInt();
        start = scanner.nextInt();
        int[][] crossover = new int[n][n];
        char[][] exit = new char[n][n];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                crossover[i][j] =scanner.nextInt();
            }
        }
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                exit[i][j] =scanner.next().charAt(0);

            }
        }
        outPut(exit,crossover,start);
    }

}
