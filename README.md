code for get the right answer

Talk is cheap, show me code.

The porpose of this repository is that get the right answer use Golang code.

1. What's the difference between nil slice and empty slice
2. `make` VS `new`

    `make` is a built-in function, which only worked on slice, map and channel, return a regular values(slice, map or channel).

    `new` only returns pointers to initialised memory.
3. Concurrency problems
    - Is `map` concurrent-safe?
    - How to implement a concurrent-safe `map`?

    Do not communicate by sharing memory; instead, share memory by communicating.

    A `data race` occurs whenever two goroutines access the same variable concurrently and at least one of the accesses is a write.

    Data structures that are never modified or are immutable are inherently conrurrency-safe and need no synchronization.

    - Not to write the variable
    - Avoid accessing the variable from multiple goroutines
    - Allow many goroutines to access the variable, but only one at a time(mutual exclusion)
    

