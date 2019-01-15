# Reasons for concurrency and parallelism


To complete this exercise you will have to use git. Create one or several commits that adds answers to the following questions and push it to your groups repository to complete the task.

When answering the questions, remember to use all the resources at your disposal. Asking the internet isn't a form of "cheating", it's a way of learning.

 ### What is concurrency? What is parallelism? What's the difference?
 > Parallelism is when multiple copies of the same data are run at exactly the same time, but they use different data. I.e If I have two machines that are botting in world of warcraft with the same botting program, they run in "parallell".

> Concurrency can be the same as parallelism but they have the ability to communicate with each other.  Two tasks are concurrent when the start end end time overlap.

 ### Why have machines become increasingly multicore in the past decade?
 > When we only have one core, several processes need to use the same one. This might result in bottlenects. Also one core can only execute 1 action at any time, so to have real parallell execution you need more than 1 core.   

> We also have the power/temperature problem with transistors. When the transistors are getting smaller, we have the oportunity to have more of them. In turn this produces more heat and use more power. Since we can't increase the clock speed as much, we add cores.

 ### What kinds of problems motivates the need for concurrent execution?
 (Or phrased differently: What problems do concurrency help in solving?)
 > Mainly, concurrency improves performance, even on single-core systems that can't utilize parallelism.  I.e i/o operations traditionally is a slow operation, the CPU is idle for several clock cycles when fetching from the memory / sending to GPU… etc. This is wasted clock cycles that could be used to compute other tasks. This is the main problem concurrency solves.

 > Concurrency is used to hide latency, so the perceived flow of a program is more smooth.

 ### Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
 (Come back to this after you have worked on part 4 of this exercise)
> Both.
> Harder:
Concurrency makes the job of the programmer harder. When programming concurrently machine code interleaves the instructions, and that is non-determenistic. This means that the order that the code is executed in, can be different for each time you execute it. Even with the same input.  This makes it challenging to debug a program.
Race Conditions may occur. Race Conditions is when two threads access the same resource at the same time and there happen to be a conflict. You might get bugs that only happens ONCE and never again,

> Easier:
When programming programs / services that need a lot of data intensive work or a lot of connection, concurrent programming makes life easier. Say you have a web-server and clients try to connect to your page. A concurrent program will make many connections feel much more smooth for the end user, and makes the server able to handle more connections.

> It also helps the program be alot faster if programmed correctly.

 ### What are the differences between processes, threads, green threads, and coroutines?
> A processes is a instance of a running program. Like the "universe" of the running program. It's "slow" to switch context between processes. Managed by the OS.

> A thread: Threads are sequential instructions within a particular process, with their own smaller context. Fast to context switch. Managed by the OS.

> Green threads: Green threads are not actually threads, but using a runtime library or an virtual enviromnent to emulate threads in a multithreaded enviroment.

> Coroutines: Coroutines are cooperating functions. I.e if you have two functions they cooperate on the timing on which they are executed and can also pass values between each others.

 ### Which one of these do `pthread_create()` (C/POSIX), `threading.Thread()` (Python), `go` (Go) create?
> pthread_create(): This makes a thread
thread.Thread(): This makes a thread
Go(go) is creating green thread. They handles the scheduling and allocation of the goroutines internally.

 ### How does pythons Global Interpreter Lock (GIL) influence the way a python Thread behaves?
> GIL is a mutex/lock that only allows one thread to run in the python program at any time. Even though this does not affect single-core users, this means that you can't run the code in parallel on multiple cores. This is a major bottleneck.  

 ### With this in mind: What is the workaround for the GIL (Hint: it's another module)?
> You can either use alternative interpreters, or you can use multi-processing.  Multiprocessing is a technique where you use multiple python processes that gets its own python interpreter and memory space.

 ### What does `func GOMAXPROCS(n int) int` change?

 > How many CPU's that can be executed simultaneously. If you use runtime.NumCPU(), then you can get the local machines number of logical CPU's
