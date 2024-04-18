If written answers are required, you can add them to this file. Just copy the
relevant questions from the root of the repo, preferably in
[Markdown](https://guides.github.com/features/mastering-markdown/) format :)


### stuff 1

What happens if you remove the go-command from the Seek call in the main function?
If the go-command is removed the program will execute sequentially and not parallel. A new thread wont be created.

What happens if you switch the declaration wg := new(sync.WaitGroup) to var wg sync.WaitGroup and the parameter wg *sync.WaitGroup to wg sync.WaitGroup?
You would pass the value of the wait group instead of the pointer to the wait group. This would mean that the wait group would be copied and the wait group in the main function would not be updated.

What happens if you remove the buffer on the channel match?
It would result in the channel blocking. The channel would block until a receiver is ready to receive the message. This would result in a deadlock.

What happens if you remove the default-case from the case-statement in the main function?
If the default case is removed the program will block until a message is received. This would result inthe program hanging and not completing.



### stuff 2

Before optimization:

time go run julia.go

real    0m9,866s
user    0m10,556s
sys     0m0,294s

--------------------------------------------------

After optimization (one async in main):

time go run julia.go

real    0m6,847s
user    0m10,369s
sys     0m0,161s

From 9.866s to 6.847s, 30% faster.


### stuff 3

mapreduce:

total time: 87 ms average time/run: 0.87 ms

single:

total time: 479 ms average time/run: 4.79 ms

