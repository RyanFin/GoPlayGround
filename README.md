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

### Go Specific (Interview)

#### Go Routines 
- Go Routines vs Multithreading vs Multiprocessing:

#### Go WaitGroups
- Goroutine sync.WaitGroups: https://yourbasic.org/golang/wait-for-goroutines-waitgroup/

##### GoRoutine Deadlocks
- Simple deadlock example: https://yourbasic.org/golang/detect-deadlock/

#### Go Channels
- Channels are objects which allow you to store data in a first-in-first-out(FIFO) queue.
- Go Channels tutorial: https://tutorialedge.net/golang/go-channels-tutorial/
- Unbuffered vs Buffered channels in Go: https://www.ardanlabs.com/blog/2014/02/the-nature-of-channels-in-go.html
- Channel buffering: https://gobyexample.com/channel-buffering
  - By default channels are buffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.

- **Unbuffered channels have no capacity set**. When attempting to *send* a resource, the unbuffered channel will lock if there is no goroutine waiting to receive the resource from the channel. When attempting to *receive* a resource, the unbuffered channel will lock if there is no goroutine waiting to send a resource to the channel.

- **Buffered channels have a set capacity** which determines how much data can be stored in the channel. In a buffered channel, if the channel is full and the goroutine attempts to *send* data, the goroutine will lock until there is space in the channel. If the buffered channel attempts to *receive* data and the channel is empty, the goroutine will lock as it waits for the buffer to fill with data.

#### Go Mutex

#### Go Interfaces

#### Critical Sections

####  Race Conditions

#### Go Defer Statements
- The deferred call's arguments are evaluated immediately, however the deferred code is not executed until the surrounfing function returns.
- Go Defer statement: https://tour.golang.org/flowcontrol/12

#### Go os.Signal
- Channels that take os.signal objects: s := make(chan os.Signal, 1). This is a buffered channel with a capacity of 1.

#### Anonymous Functions in Go

#### Atomic Counters
- Atomic Counters: https://gobyexample.com/atomic-counters
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




