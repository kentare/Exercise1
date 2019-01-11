Results Part 4
--------------------
Finally, create a new file called RESULT.md inside this directory explaining what happens and why. Remember to add, commit and push the new file.

Every program creates two threads and execute the functions concurrently. Because we use a share memory (i) there is conflict and they access the same resource at the same time (race condition). Therefore the answer is different every time we run the program.
To battle this problem we might look into using a mutex structure so incrementing function gets access to the resource first(i), then the decrementing function get access to the resource (i) when the incrementing is done. This way the answer always will be 0.  
