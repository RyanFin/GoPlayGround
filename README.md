# GoPlayGround

## Interview Practice
  - 4 TIPS FOR PREPARING FOR A CODING INTERVIEW: https://generalassemb.ly/blog/4-tips-for-preparing-for-a-coding-interview/
    - [Hackerrank](https://www.hackerrank.com/)
    - [LeetCode](https://leetcode.com/)
    - [Codewars](https://www.codewars.com/)
    - [AlgoExpert](https://www.algoexpert.io/)

## Go General
- Naming Conventions for Golang Functions: https://www.golangprograms.com/naming-conventions-for-golang-functions.html
- Anonymous functions: https://www.geeksforgeeks.org/anonymous-function-in-go-language/
- Array increment types, arr[i]++ vs arr[i++]: https://stackoverflow.com/questions/7595247/array-increment-types-in-c-arrayi-vs-arrayi
- Unsigned Integers (uints) are always non-negative (either zero or positive): https://www.cs.utah.edu/~germain/PPS/Topics/unsigned_integer.html#:~:text=Unsigned%20Integers%20(often%20called%20%22uints,will%20always%20be%20non%2Dnegative.
- For loop  without a condition will loop repeatedly until you break out of the loop or return from the enclosing function: https://gobyexample.com/for

### Go Specific (Interview Preparation)

#### Go Routines
- Go Routines vs Multithreading vs Multiprocessing:
  Process: Generated from the executable (go build) when you run it 
  Multiprocessing: Multiple Processes (applications) running at the same time
  Thread: 
  GoRoutine: Main goroutine as well as concurrent ones you can create: https://gobyexample.com/goroutines
    - Need to sleep to run goroutines: https://stackoverflow.com/questions/15771232/why-is-time-sleep-required-to-run-certain-goroutines
  - Why goroutines instead of threads?: https://golang.org/doc/faq#goroutines

#### Go WaitGroups
- To wait for multiple go routines to finish, use a wait group. 
- Goroutine sync.WaitGroups: https://yourbasic.org/golang/wait-for-goroutines-waitgroup/

##### GoRoutine Deadlocks
- Simple deadlock example: https://yourbasic.org/golang/detect-deadlock/

#### Go Channels
- New channels tutorial: https://www.scaler.com/topics/golang/buffered-and-unbuffered-channel-in-golang/
- Channels are objects which allow you to store data in a first-in-first-out(FIFO) queue.
- Go Channels tutorial: https://tutorialedge.net/golang/go-channels-tutorial/
- Unbuffered vs Buffered channels in Go: https://www.ardanlabs.com/blog/2014/02/the-nature-of-channels-in-go.html
- Channel buffering: https://gobyexample.com/channel-buffering
  - By default channels are buffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.

- **Unbuffered channels have no capacity set**. When attempting to *send* a resource, the unbuffered channel will lock (deadlock) if there is no goroutine waiting to receive the resource from the channel. When attempting to *receive* a resource, the unbuffered channel will lock (deadlock) if there is no goroutine waiting to send a resource to the channel.

- **Buffered channels have a set capacity** which determines how much data can be stored in the channel. In a buffered channel, if the channel is full and the goroutine attempts to *send* data, the goroutine will lock until there is space in the channel. If the buffered channel attempts to *receive* data and the channel is empty, the goroutine will lock as it waits for the buffer to fill with data.

#### Go Mutex
- Avoid race conditions by syncronising goroutines access to the state (variables) 

#### Go Atomic Operations
- Atomic counters: https://gobyexample.com/atomic-counters

#### Go Worker Pools

#### Go Interfaces
- Interfaces are named collections of method signatures: https://gobyexample.com/interfaces 
<!-- - Interfaces work identical to the ones in Java, as they define contracts for what the concrete implementation  must have -->
- Interfaces in Go enable developers to decide what functions they want as a generic function. That is, it enables multiple  instances of concrete Go Structs to implement them in different ways with the same method name.
- How to use interfaces in Go: https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
  - **interface vs interface{} type**
  - **pointer receivers vs value receivers** when defining functions that belong to a struct

#### Critical Sections
- How to use a mutex to define critical sections of code and fix race conditions: https://www.golangprograms.com/how-to-use-a-mutex-to-define-critical-sections-of-code-and-fix-race-conditions.html
  - A mutex is used to create a critical section around code that ensures that only one goroutine is able to execute that code section. This solves race condition issues.

####  Race Condition
- Golang: Concurrency is Hard; So What Can We Do About It? | Medium: https://medium.com/dm03514-tech-blog/golang-candidates-and-contexts-a-heuristic-approach-to-race-condition-detection-e2b230e70d08
- A race condition in Go is where more than one GoRoutine attempt to access the same variable memory address simultaneously; one of which is writing. The application must be synchronised to ensure that this does not happen.

#### Go Defer Statements
- The deferred call's arguments are evaluated immediately, however the deferred code is not executed until the surrounfing function returns.
- Go Defer statement: https://tour.golang.org/flowcontrol/12

#### Go os.Signal
- Signals in golang: https://gobyexample.com/signals
- os/signal package: https://golang.org/pkg/os/signal/
- Channels that take os.signal objects: s := make(chan os.Signal, 1). This is a buffered channel with a capacity of 1.
- Notify: https://golang.org/pkg/os/signal/#Notify
  - Notify causes package signal to relay incoming signals to c.
If no signals are provided (second parameter, third param, fourth param...), all incoming signals will be relayed to c.
Otherwise, just the provided signals will.

#### Continue Statement
- Forces the next iteration  of the loop to take place: https://www.tutorialspoint.com/go/go_continue_statement.htm

#### Anonymous Functions in Go

#### Atomic Counters
- Atomic Counters: https://gobyexample.com/atomic-counters

#### Go Select
- GoByExample Select: https://gobyexample.com/select

#### Duck Typing
- Duck typing: https://en.wikipedia.org/wiki/Duck_typing

## Coverage
- Generate coverage report: $ go test -coverprofile=coverage.out
- View coverage report in web browser: $ go tool cover -html=coverage.out 

## Linked Lists
- LeetCode Breakdown: https://leetcode.com/explore/learn/card/linked-list/209/singly-linked-list/1287/
- Go Official Linked List Doc: https://golang.org/pkg/container/list/
- Linked List Tutorial Junmin Lee
  - Part 1: https://www.youtube.com/watch?v=1S0_-VxPLJo&list=PL0q7mDmXPZm7s7weikYLpNZBKk5dCoWm6&index=3&ab_channel=JunminLee
  - Part 2: https://www.youtube.com/watch?v=8QoynPUY9_8&list=PL0q7mDmXPZm7s7weikYLpNZBKk5dCoWm6&index=4&ab_channel=JunminLee
  - LinkedList Deletion at a particular position: https://www.youtube.com/watch?v=f1r_jxCyOl0&ab_channel=NesoAcademy
  - Reversing a linked list: https://medium.com/outco/reversing-a-linked-list-easy-as-1-2-3-560fbffe2088
  - Insertion into a sorted doubly linked list: https://www.geeksforgeeks.org/insert-value-sorted-way-sorted-doubly-linked-list/

## Tries
- Data Structures in Golang (Junmin Lee)
  - Trie Implementation: https://www.youtube.com/watch?v=nL7BHR5vJDc&ab_channel=JunminLee

## iota
- https://yourbasic.org/golang/iota/

## Variadic Functions
- https://gobyexample.com/variadic-functions


## Data Structures and Algorithms [Robert Lafore]
  - Supporting Workshop Applets:
    - https://cs.brynmawr.edu/Courses/cs206/spring2004/lafore.html
    - https://cs.ccsu.edu/~markov/ccsu_courses/501Syllabus.html
      - Run the Java applet with the following command: $ appletviewer <html file>

## Problem Solving

### Diagonal Difference Challenge
- https://rishabh1403.com/posts/coding/hackerrank/2018/08/hackerrank-solution-of-diagonal-difference-in-golang/

### Plus Minus Challenge

#### Roundup.go
- https://gist.github.com/DavidVaini/10308388
- https://play.golang.org/p/Xog_9wXSqj

### Staircase
- https://rishabh1403.com/posts/coding/hackerrank/2018/09/hackerrank-solution-of-staircase-in-golang/

### Left-rotation
- https://www.hackerrank.com/challenges/array-left-rotation/forum




