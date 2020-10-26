# Linked List



`Node`have the following fields and functions:

* `Val`: the node's value (we will use `interface{}`).
* `Next()`: pointer to the next node.
* `Prev()`: pointer to the previous node.
* `First()`: pointer to the first node (head).
* `Last()`: pointer to the last node (tail).

 `List`  have the following fields and functions:

* `NewList(args ...interface{}) *List`: creates a new linked list preserving the order of the values.
* `First()`: pointer to the first node (head).
* `Last()`: pointer to the last node (tail).
* `PushBack(v interface{})`: insert value at back.
* `PopBack() (interface{}, error)`: remove value at back.
* `PushFront(v interface{}) `: insert value at front.
* `PopFront() (interface{}, error)`: remove value at front.
* `Reverse()`: reverse the linked list.


## Running the tests

To run the tests run the command `go test` from within the exercise directory.

If the test suite contains benchmarks, you can run these with the `--bench` and `--benchmem`
flags:

    go test -v --bench . --benchmem



