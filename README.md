# Golang Concurrent/Single-Threaded Data Structures Library w/ Generics

Welcome to the Golang Data Structures Library! This library provides a variety of efficient and thread-safe data structures implemented in Go, designed to enhance performance and concurrency handling in your applications.

## Features

- **Queue**: A dynamically resizing queue with efficient enqueue and dequeue operations.
- **Deque**: A double-ended queue that allows insertion and deletion from both ends.
- **Stack**: A standard stack with push and pop operations.
- **Concurrent Queue (cqueue)**: A thread-safe queue using `sync.RWMutex` for concurrent access.
- **Concurrent Stack (cstack)**: A thread-safe stack implementation with `sync.RWMutex`.
- **Concurrent Deque (cdeque)**: A thread-safe double-ended queue with `sync.RWMutex`.

## Why use my data structures?

- I implemented them all optimally. For queue and deque, I use a circular array under the hood and use a formula to resize based on how large the structure becomes. They also have escapes from the initial data structures, by means of a ToSlice() function, to allow the you to look at the underlying data if necessary. 

## Installation

To install the library, use the following command:

```bash
go get github.com/Shreyas-Adireddy/data_structures
```

## Usage

### Queue

A standard queue with dynamic resizing.

**Functions:**

- `Enqueue(value T)`: Adds an element to the rear of the queue.
- `Dequeue() (T, error)`: Removes and returns the front element of the queue.
- `Front() (T, error)`: Returns the front element without removing it.
- `Back() (T, error)`: Returns the rear element without removing it.
- `Size() int`: Returns the number of elements in the queue.
- `IsEmpty() bool`: Checks if the queue is empty.
- `Clear()`: Clears all elements in the queue.

**Example:**

```go
import "github.com/Shreyas-Adireddy/data_structures"

func main() {
    q := queue.New[int]()
    q.Enqueue(10)
    q.Enqueue(20)
    fmt.Println(q.Dequeue()) // Outputs: 10
}
```

### Deque

A double-ended queue that supports insertion and deletion from both ends.

**Functions:**

- `AddFront(value T)`: Adds an element to the front of the deque.
- `AddRear(value T)`: Adds an element to the rear of the deque.
- `RemoveFront() (T, error)`: Removes and returns the front element.
- `RemoveRear() (T, error)`: Removes and returns the rear element.
- `Front() (T, error)`: Returns the front element without removing it.
- `Rear() (T, error)`: Returns the rear element without removing it.
- `Size() int`: Returns the number of elements in the deque.
- `IsEmpty() bool`: Checks if the deque is empty.
- `Clear()`: Clears all elements in the deque.

**Example:**

```go
import "github.com/Shreyas-Adireddy/data_structures/deque"

func main() {
    d := deque.New[int]()
    d.AddFront(10)
    d.AddRear(20)
    fmt.Println(d.RemoveFront()) // Outputs: 10
}
```

### Stack

A standard stack with push and pop operations.

**Functions:**

- `Push(value T)`: Adds an element to the top of the stack.
- `Pop() (T, error)`: Removes and returns the top element of the stack.
- `Peek() (T, error)`: Returns the top element without removing it.
- `Size() int`: Returns the number of elements in the stack.
- `IsEmpty() bool`: Checks if the stack is empty.
- `Clear()`: Clears all elements in the stack.

**Example:**

```go
import "github.com/Shreyas-Adireddy/data_structures/stack"

func main() {
    s := stack.New[int]()
    s.Push(10)
    s.Push(20)
    fmt.Println(s.Pop()) // Outputs: 20
}
```

### Concurrent Queue (cqueue)

A thread-safe queue using `sync.RWMutex` for concurrent access.

**Functions:**

- `Enqueue(value T)`: Adds an element to the rear of the queue.
- `Dequeue() (T, error)`: Removes and returns the front element of the queue.
- `Front() (T, error)`: Returns the front element without removing it.
- `Back() (T, error)`: Returns the rear element without removing it.
- `Size() int`: Returns the number of elements in the queue.
- `IsEmpty() bool`: Checks if the queue is empty.
- `Clear()`: Clears all elements in the queue.

**Example:**

```go
import "github.com/Shreyas-Adireddy/data_structures/cqueue"

func main() {
    cq := cqueue.New[int]()
    cq.Enqueue(10)
    cq.Enqueue(20)
    fmt.Println(cq.Dequeue()) // Outputs: 10
}
```

### Concurrent Stack (cstack)

A thread-safe stack implementation.

**Functions:**

- `Push(value T)`: Adds an element to the top of the stack.
- `Pop() (T, error)`: Removes and returns the top element of the stack.
- `Peek() (T, error)`: Returns the top element without removing it.
- `Size() int`: Returns the number of elements in the stack.
- `IsEmpty() bool`: Checks if the stack is empty.
- `Clear()`: Clears all elements in the stack.

**Example:**

```go
import "github.com/Shreyas-Adireddy/data_structures/cstack"

func main() {
    cs := cstack.New[int]()
    cs.Push(10)
    cs.Push(20)
    fmt.Println(cs.Pop()) // Outputs: 20
}
```

### Concurrent Deque (cdeque)

A thread-safe double-ended queue.

**Functions:**

- `AddFront(value T)`: Adds an element to the front of the deque.
- `AddRear(value T)`: Adds an element to the rear of the deque.
- `RemoveFront() (T, error)`: Removes and returns the front element.
- `RemoveRear() (T, error)`: Removes and returns the rear element.
- `Front() (T, error)`: Returns the front element without removing it.
- `Rear() (T, error)`: Returns the rear element without removing it.
- `Size() int`: Returns the number of elements in the deque.
- `IsEmpty() bool`: Checks if the deque is empty.
- `Clear()`: Clears all elements in the deque.

**Example:**

```go
import "github.com/Shreyas-Adireddy/data_structures/cdeque"

func main() {
    cd := cdeque.New[int]()
    cd.AddFront(10)
    cd.AddRear(20)
    fmt.Println(cd.RemoveFront()) // Outputs: 10
}
```

## Contributing

We welcome contributions to improve this library! Here are some ways you can help:

1. **Report Bugs**: If you encounter any issues, please create a GitHub issue.
2. **Suggest Features**: Have an idea for a new feature? Let us know by opening an issue.
3. **Submit Pull Requests**: If you'd like to fix a bug or implement a feature, feel free to submit a pull request. Make sure to include tests and documentation for your changes.
4. **Improve Testcases**: Help me write some more testcases for each of the data structures.

### Getting Started

1. **Fork the repository**: Click the "Fork" button at the top right of this page.
2. **Clone your fork**: `git clone https://github.com/yourusername/yourrepository.git`
3. **Create a branch**: `git checkout -b my-feature-branch`
4. **Make your changes**: Implement your feature or bug fix.
5. **Run tests**: Ensure all tests pass.
6. **Commit your changes**: `git commit -m 'Add some feature'`
7. **Push to the branch**: `git push origin my-feature-branch`
8. **Create a Pull Request**: Open a pull request with a clear description of your changes.

Thank you for your interest in contributing to our project!

---

Feel free to customize the content to match your specific project and preferences.
