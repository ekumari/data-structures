import java.util.HashMap;

public class HashMapExample {
    public static void main(String[] args) {
        HashMap<String, Integer> map = new HashMap<>();

        // Insert
        map.put("apple", 10);
        map.put("banana", 20);
        map.put("orange", 30);

        System.out.println("Map: " + map);

        // Access
        System.out.println("Apple price: " + map.get("apple"));

        // Delete
        map.remove("banana");
        System.out.println("After deletion: " + map);

        // Check if key exists
        if (map.containsKey("banana")) {
            System.out.println("Banana price: " + map.get("banana"));
        } else {
            System.out.println("Banana not found");
        }
    }
}

