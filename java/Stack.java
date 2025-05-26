public class Stack {
    private int maxSize;
    private int[] stackArray;
    private int top;

    public Stack(int size) {
        maxSize = size;
        stackArray = new int[maxSize];
        top = -1;
    }

    public void push(int item) {
        if (isFull()) {
            System.out.println("Stack is full");
            return;
        }
        stackArray[++top] = item;
    }

    public int pop() {
        if (isEmpty()) {
            System.out.println("Stack is empty");
            return -1;
        }
        return stackArray[top--];
    }

    public int peek() {
        if (isEmpty()) {
            System.out.println("Stack is empty");
            return -1;
        }
        return stackArray[top];
    }

    public boolean isEmpty() {
        return (top == -1);
    }

    public boolean isFull() {
        return (top == maxSize - 1);
    }

    public int size() {
        return top + 1;
    }

    public static void main(String[] args) {
        Stack stack = new Stack(10);

        stack.push(10);
        stack.push(20);
        stack.push(30);

        System.out.println("Stack size: " + stack.size());    // Output: 3
        System.out.println("Top element: " + stack.peek());   // Output: 30

        System.out.println("Popped: " + stack.pop());         // Output: 30
        System.out.println("Popped: " + stack.pop());         // Output: 20
        System.out.println("Stack empty? " + stack.isEmpty()); // Output: false

        System.out.println("Popped: " + stack.pop());         // Output: 10
        System.out.println("Stack empty? " + stack.isEmpty()); // Output: true
    }
}

