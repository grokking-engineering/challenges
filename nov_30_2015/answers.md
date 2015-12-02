# Question 1 (10pts)

In most programming languages, the following expression will not return 1.0

```ruby
1.0 + 2.0/3.0 - 2.0/3.0
```

For example, this is true in [Ruby](https://repl.it/B4aR/0), [Javascript](https://repl.it/B4aN), [Golang](http://play.golang.org/p/zYnQuxiJ_d), [Java](https://repl.it/B4aM)

What do you think is the cause of this? Explain in a short answer.

## Suggested answers

Basically, this is because there's a limit to how accurate floating point numbers are presented.
As all things on computer are stored based on a base 2 system (0 & 1),
anything that can't be represented as a sum of powers of 2 wil be rouned up/down.

When you perform arithmetic with these rounded numbers, you will get wrong result.
This is why scientist, programmers working with money, and anyone who care about accuracy should be extremely careful with floating point numbers 

## Scoring scheme:

Condition                               | Points |
----------------------------------------|--------|
Did not mention floating point accuracy | 0 pts  |
----------------------------------------|--------|
Mention floating point accuraccy        | 10 pts | 

****

# Question 2 (10pts)

__Choose all the correct answers:__

What action(s) __will likely__ leading to "too many open files" error on your server in Linux?

1. Open and read text file (i.e: ~/.bashrc)
2. Send HTTP request and read their body
3. Running multiple sub-processes and create pipes to access their stdin, stdout
4. Open and close a text file multiple times
5. Adding more processes to your PostgreSQL database instance

## Suggested answers

It's (1), (2), (3) and (5).

In Linux, the error "too many open files" is related to the 
limited number of __kernel handles__ that processes can use at a time.

Historically, the only thing that a process often open a lot is file,
hence the error name. Thus (1) can be chosen

A socket connection (to read/write HTTP request) is also a kernel handle,
and you can forget to close it (see: [this question](http://unix.stackexchange.com/questions/157351/why-are-tcp-ip-sockets-considered-open-files)).
Thus answer (2) is chosen

A pipe is also considered a file descriptor -> answer (3) is chosen

(5) is a little bit far-fetched (we admit). It goes like this:
If other processes are using most of the file descriptors on your machine, and you try
to add a couple more PostgreSQL processes, the extra processes will cause the error.
The reasons is being a persistent-to-disk database, the extra PostreSQL processes will
need to open files too :) (but we asked for likely cause)

(4) is not chosen because we explicitly closed the file
 
If you ever see this error, remember these and run `lsof` to debug :)

****

# Question 3 (10pts)

Choose all correct answers:

In theory, __TCP/IP__ provides guaranteed delivery of packets. However, in reality, when  you try to send a request to a server (say, using `curl -X GET 'http://www.google.com'`), you might never receive a response. This is possibly because:

1. TCP/IP doesn't work in reality (oops)
2. HTTP request are processed via the OSI stack and not via TCP/IP stack
3. Client or server(s) might stop handling the request
4. There is a firewall blocking traffic coming to you
5. Network delays might cause different packets to arrive out of order, making the response un-processable

## Suggested answers

The answer is (3) and (4).

(1), (2) and (5) are just crap we write as diversion, ignore them

While TCP/IP does guarantee delivery and ordering, it does not guarantee __response time__. This mean, a packet can go on for a very long amount of time (till next year), and it's okay.

Obviously, things won't wait for a year on the Internet, so people will timeout things if it takes too long. Thus (3) is correct.

Also, when the packet goes to your machine, it can be blocked by firewall on the OS (on Linux, you can use iptables). Realistic usecase might be to block traffic coming from a ceterin IP (of a bad service, hacker, DDOS-er). Thus (4) is correct

****

# Question 4 (30pts)

Given the following programs:

```c
#include <stdio.h>
#include <stdlib.h>

void
f(void)
{
    int a[4];
    int *b = malloc(16);
    int *c;
    int i;

    c = a;
    for (i = 0; i < 4; i++)
  a[i] = 100 + i;
    c[0] = 200;
    printf("1: a[0] = %d, a[1] = %d, a[2] = %d, a[3] = %d\n",
     a[0], a[1], a[2], a[3]);

    c[1] = 300;
    *(c + 2) = 301;
    3[c] = 302;
    printf("2: a[0] = %d, a[1] = %d, a[2] = %d, a[3] = %d\n",
     a[0], a[1], a[2], a[3]);

    c = c + 1;
    *c = 400;
    printf("3: a[0] = %d, a[1] = %d, a[2] = %d, a[3] = %d\n",
     a[0], a[1], a[2], a[3]);

    c = (int *) ((char *) c + 1);
    *c = 500;
    printf("4: a[0] = %d, a[1] = %d, a[2] = %d, a[3] = %d\n",
     a[0], a[1], a[2], a[3]);

    b = (int *) a + 1;
    c = (int *) ((char *) a + 1);
    printf("5: a = %p, b = %p, c = %p\n", a, b, c);
}

int
main(int ac, char **av)
{
    f();
    return 0;
}
```

[You can run it here](https://repl.it/B4a1/1).

For each line printer (1-5), explain why the result is that way

## Suggested answers

The loop on lines 13 and 14 sets each element of `a` to `100`, `101`, `102` and `103`, respectively. After line 10, `c` and `a` references the same memory location, so `c[0] = 200` modifies the same memory location `a` is pointing to, i.e. it sets `a[0]` to `200`.

Line 19 sets `c[1] = 300`, evidently from the code. Line 20 sets `c[2]` to `301` as `c[2]` is equivalent to `*(c + 2)`. Line 21 sets `c[3]` to `302` as `c[3]` is equivalent to `3[c]`. Keep in mind that `c` still refers to the same memory location as `a`.

Line 25 makes `c` point to the second element of `a`, i.e. `a[1]`, and line 26 sets `a[1]` to `400`.

Line 30 sets `c` to point to the second element of `a`, off by 1 byte, so setting the memory pointed to by `c` to `500` (`00000000 00000000 00000001 11110100`) is undefined behavior. What is observed differs based on which CPU architecture the code is compiled for and is run on. On some ARM CPUs, or on Intel x86 CPUs with unaligned memory access trap turned on, the code will generate a hardware fault (usually leads to a SIGBUS on Linux), stopping the program or summoning up the debugger. On architectures whose configuration allows unaligned memory access, what values are observed in `a[1]` and `a[2]` depend on which endianness the architecture has. If you see `128144` (`00000000 00000001 11110100 10010000`) for `a[1]` and `256` (`00000000 00000000 00000001 00000000`) for `a[2]`, you're running on a little endian architecture.

Line 35 makes `b` point to the second element of `a`, i.e. `a[1]`. Line 36 makes `c` point to the first element of `a`, off by 1 byte.

****

# Question 5 (50pts)

We will play i a game called " CodeBreaker". How it works is described as follow:

The game have 2 role: __Coder__ and __Breaker__. Each game, the __Coder__ sets a code of 4 spots, consisting of colors in the set (Red, Green, Blue, Yellow, Black, White) (order matters). There is no limit on the number of times a color can re-appear (i.e: Black - White - Blue - Green or Red - Red - Red - Black are valid codes)

The __Breaker__ have 10 turns to break this code. Each turn, he make a guess consisting of colors from the set and give it to the __Coder__. The __Coder__ then tell him 2 things: the number of spots that match both color and position; as well as number of colors that match, but have the wrong position (excluding those that match both).

For example:
```
Coder:   Red - Green - Black - Black
Breaker: Red - White - Black - Green

There are 2 spots that match both color and position: 1st (Red) and 3rd (Black).
In the remaining, the 4th spot in the guess (Green) match with the 2nd spot in the code, but have wrong position
```

The __Breaker__ then proceeds to make another guess, until he found the code, or 10 turns have ended.

You are to write a program to simulate the __Breaker__, following these spec


```
// terminology
// ----------------------------------------------------
//  - We have the this main program (driver), and your subprogram (solver)
//  - They talk via pipes:
//    - driver's stdout > solver's stdin
//    - solver's stdout > driver's stdin
//  - All commands are terminated with the character '\n' (your separator)
//  - To make it simpler to write examples, whenever you see:
//    - '>:' it indicates that driver is writing into solver's stdin
//    - '<:' it indicates that solver is writing to driver's stdin
//    - The data written is on the right of the ':'
//  - When we test your program, we expect it to terminates when the game end
//    e.g, when you receive either TIMEOUT or CORRECT
//  - After you have written the guesses to stdout, flush your output buffer
//    so we can read it
//
// convention
// ----------------------------------------------------
// colors are indicated using 1 ASCII character, as followed:
//  - Red:       R
//  - Green:     G
//  - Blue:      U
//  - Yellow:    Y
//  - Black:     B
//  - White:     W
//
// input/output
// ----------------------------------------------------
//  - >:START\n                 // indicate new game
//  - <:RRRR\n                  // solver make a first guess
//  - >:0 1 2 CONTINUE\n        // driver returns the following 4 things:
//                              //   <turn no>
//                              //   <no of spot matching color & position (1)>
//                              //   <no of spot matching color ONLY, excluding (1)>
//                              //   <a state>, where
//                              //      CONTINUE indicate you can make another guess
//                              //      CORRECT indicate you have guessed correctly
//                              //      TIMEOUT indicate you didn't find the answer
//
// example: let's say the code is RWBY
// ----------------------------------------------------
// >:START\n
// <:RRRR\n
// >:1 1 0 CONTINUE\n           // 1st spot matched
// <:RBBB\n
// >:2 2 0 CONTINUE\n           // 1st & 3rd spot matched
// <:RBWW\n
// >:3 1 2 CONTINUE\n           // 1st spot matched, 2nd & 3rd are in wrong position
// <:RWBB\n
// >:4 3 0 CONTINUE\n           // 1st, 2nd & 3rd matched
// <:RWBY\n
// >:5 4 0 CORRECT\n            // all matched
//
//
// the same game might exceed 10 turns, in which case we return a TIMEOUT
//
// ... some other steps ..
// <:RWBB\n
// >:9 3 0 CONTINUE\n           // 1st, 2nd & 3rd matched
// <:RWBU\n
// >:10 3 0 TIMEOUT\n           // 1st, 2nd & 3rd matched, timeout
//
```

You can write your __Breaker__ program in any languages, as long as it follows the rule about stdin and stdout and can be run on a Linux machine.

Assume we are on a new Linux box, please provides instructions as comment in your code to help us build/run your program.
For example, if you are running Ruby and require gems, please include the Gemfile, your ruby version (irb/jruby, etc) as comment

The source code for the test driver, as well as a sample client is given here, together with a binary compiled for Linux (GOARCH=amd64)
<http://tinyurl.com/grokkingchallenge2-code>

## Suggested answers

This is just a matter of writing code :D

You an see how we run the tests from this video (.ogv)

<https://github.com/grokking-engineering/challenges/raw/master/nov_30_2015/guide.ogv>
