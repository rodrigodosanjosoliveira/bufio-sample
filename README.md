### Bufio NewWriter

```bash
bufio.Writer
```

To avoid the overhead of many small write operations Golang is shipped with bufio.Writer. Data, instead of going straight to destination (implementing io.Writer interface) are first accumulated inside the buffer and send out when buffer is full:

Let’s visualise how buffering works with nine writes (one character each) when buffer has space for 4 characters:

```bash
producer         buffer           destination (io.Writer)

   a    ----->   a
   b    ----->   ab
   c    ----->   abc
   d    ----->   abcd
   e    ----->   e      ------>   abcd
   f    ----->   ef               abcd
   g    ----->   efg              abcd
   h    ----->   efgh             abcd
   i    ----->   i      ------>   abcdefgh
```
