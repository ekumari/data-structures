import java.util.ArrayList;

public class ArraysArrayList {
    public static void main(String[] args) {
        // Array: fixed size
        int[] arr = {1, 2, 3};
        System.out.print("Array: ");
        for (int i : arr) {
            System.out.print(i + " ");
        }
        System.out.println();

        // ArrayList: dynamic size
        ArrayList<Integer> list = new ArrayList<>();
        list.add(4);
        list.add(5);
        list.add(6);
        System.out.println("ArrayList: " + list);

        // Add more elements
        list.add(7);
        list.add(8);
        System.out.println("ArrayList after adding elements: " + list);

        // Access elements
        System.out.println("Element at index 2: " + list.get(2));

        // Size of ArrayList
        System.out.println("Size of ArrayList: " + list.size());
    }
}

