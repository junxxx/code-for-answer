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
4. process, thread, goroutines
    - Stack
        Each OS thread has a fixed-size `stack`(often 2MB). But a goroutine start life with a small stack, typically 2KB.It grows and shrinks as needed.

    - Scheduling
        Thread scheduled by OS kernel. Next a switch to kernel context
        Go runtime contains its own scheduler that use a technique known as `m:n scheduling`.It multiplexes `m` goroutines on `n` OS threads. Do not need a switch to kernel context, rescheduling a goroutine is much cheaper than rescheduling a thread.
    

