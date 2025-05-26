import java.util.LinkedList;

public class Queue {
    LinkedList<Integer> list = new LinkedList<>();

    public void enqueue(int item) {
        list.addLast(item);
    }

    public int dequeue() {
        if (list.isEmpty()) {
            System.out.println("Queue is empty");
            return -1;
        }
        return list.removeFirst();
    }

    public boolean isEmpty() {
        return list.isEmpty();
    }

    public static void main(String[] args) {
        Queue q = new Queue();
        q.enqueue(10);
        q.enqueue(20);
        q.enqueue(30);
        System.out.println(q.dequeue()); // 10
        System.out.println(q.dequeue()); // 20
        System.out.println(q.isEmpty()); // false
    }
}

