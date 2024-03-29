# Advanced Looping In Go
+++
title = "Advanced Looping in Go"
date = "FIXME"
tags = ["golang"]
categories = ["golang"]
url = "FIXME"
author = "mikit"
+++

### Overview

Looping seems like a basic topic: Write a `for` loop with a termination condition, and you're done. However there's a lot of ways you can write a `for`  loop in Go. Knowing more about the different versions of `for` loops will help you choose the best option to accomplish your tasks and will also help you prevent some bugs.

### Some Assembly Required

What kind of code is generated by the compiler for a `for` loop? To make the assembly output simple, I will write an empty `for` loop:

**Listing 1: An Empty `for` Loop**  
```
01 package main
02 
03 func main() {
04     for i := 0; i < 3; i++ {
05     }
06 }
```

Listing 1 shows a simple program with an empty `for` loop that iterates 3 times.

You can produce the assembly for that Go code using the following command:

`go build -gcflags=-S asm.go > asm.s 2>&1`

Now look at the portion of the assembly that is associated with the `for` loop. I've modified the output to make it clearer.

**Listing 2: Assembly Output**  
```
06     0x0000 00000 asm.go:3    XORL   AX, AX
07     0x0002 00002 asm.go:4    JMP    7
08     0x0004 00004 asm.go:4    INCQ   AX
09     0x0007 00007 asm.go:4    CMPQ   AX, $3
10     0x000b 00011 asm.go:4    JLT    4
```

Listing 2 shows the assembly output of the `for` loop.

On line 06, the `XORL` commands sets the `AX` register to `0`. Then on line 07, the code jumps to address `7` which is line 09. On line 09, the code compares `AX` to the value of 3 and then on line 10, the code jumps to address`4` (line 08) if the result of the comparison on line 09 is less than `3`.

On line 08, the code increments `AX` by one. Then the program will continue to move to line 09, do the comparison again and so forth. This will continue until `JLT` instruction on line 10 won't execute because `AX` is `3`.   At that point the `for` loop exits.

If you're not familiar with assembly, this logic seems backwards. But remember that in assembly you only have jumps for flow control, so this is how looping works at that level.

If you want a nice visual display of assembly when you have to read it, check out this cool tool.

[lensm](https://github.com/loov/lensm).

### Using More Than One Variable

Go supports having more than one variable in the `for` loop when necessary. This is convenient when you need to write an algorithm to check if a string is a palindrome.

**Listing 3: Two Variable `for` Loop**  
```
25 // IsPalindrome returns true if `s` is a palindrome.
26 func IsPalindrome(s string) bool {
27     rs := []rune(s) // convert to runes
28
29     for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
30         if rs[i] != rs[j] {
31             return false
32         }
33     }
34 
35     return true
36 }
```

Listing 3 shows a `for` loop with two variables. On line 29, the code initializes `i` and `j` in the init statement, then in the condition part the code checks if `i` is less than `j`, and finally in the post statement, the code increments `i` and decrements `j`.

Note that you can't write `i++, j--` in the post statement, it's not valid syntax.

### Label Breaks

There is a situation most developers eventually run into when they have a loop surrounding a switch statement. Can you detect the bug in the following code?

**Listing 4: Handling Log**  
```
11 type Log struct {
12     Level   string `json:"level"`
13     Message string `json:"message"`
14 }
15 
16 func logHandler(r io.Reader) error {
17     dec := json.NewDecoder(r)
18
19     for {
20         var log Log
21         err := dec.Decode(&log)
22
23         switch {
24         case errors.Is(err, io.EOF):
25             break
26         case err != nil:
27             return err
28         default:
29             fmt.Printf("log: %+v\n", log)
30         }
31     }
32 
33     return nil
34 }
```

Listing 4 shows a log handler that reads a stream of log messages from a reader. On lines 11-14, the code declares a `Log` type that represents a log message. On line 19, the code starts a "forever" loop and on line 21, the code reads a log message from the reader and  decodes the message into a `Log` value. . On line 23, the code declares a switch statement that checks two things. On line 24, a check is performed to see if the reader is empty and on line 26, a check is performed to see if there was an error reading the stream or decoding the log message. If there are no errors, then line 29 will execute and print the log message.




If you run this code, the bug is that the loop will never terminate. The reason is that the `break` on line 23 breaks out of the `case` and not out of the `for` loop.

To solve this issue, you can use a label break [a label](https://go.dev/ref/spec#Labeled_statements).

**Listing 5: Label Breaks**  
```
16 func logHandler(r io.Reader) error {
17     dec := json.NewDecoder(r)
18
19 loop:
20     for {
21         var log Log
22         err := dec.Decode(&log)
23
24         switch {
25         case errors.Is(err, io.EOF):
26             break loop
27         case err != nil:
28             return err
29         default:
30             fmt.Printf("log: %+v\n", log)
31         }
32     }
33 
34     return nil
35 }
```

Listing 5 shows how to fix the bug. On line 19, the code adds a label before the `for` statement and on line 26, the code can use the label to break out of the `for` loop from within the `switch` statement..

### Range Loop Semantics

Say you want to give a $1,000 bonus to your VIP bank members. There is a special `for` loop in Go called a `for range` that is perfect for this scenario.

**Listing 6: Bonus Time**  
```
05 type Account struct {
06     Name    string
07     Type    string
08     Balance int
09 }
10 
11 func main() {
12     acts := []Account{
13         {"donald", "regular", 123},
14         {"scrooge", "vip", 1_000_001},
15     }
16 
17     for _, a := range acts {
18         if a.Type == "vip" {
19             a.Balance += 1_000
20         }
21     }
22
23     fmt.Println(acts)
24 }
```

Listing 6 shows code that adds a bonus to the account of every VIP member. On lines 05-09, the code defines an `Account` type. Then on lines 17-21, the code iterates over the accounts and adds a $1,000 bonus to any VIP member.

However, when the code prints the accounts on line 23, you won’t see the changes. This is because the `a` variable declared inside the loop on line 17 is not a reference to the `Account` value being iterated over, but a copy of the `Account` value. So on line 19, the update is happening to a copy and can’t be seen outside the `for range` statement.  This form of the `for range` is called the value semantic version since the loop is operating on copies of the data.

To solve this bug, you can change the slice to hold a reference to each `Account` value using pointers ([]*Account), but don’t do that. Instead a better solution is to use the pointer semantic  version of the `for range`.

**Listing 7: Pointer Semantics**  
```
17     for i := range acts {
18         if acts[i].Type == "vip" {
19             acts[i].Balance += 1_000
20         }
21     }
22
23     fmt.Println(acts)
```

Listing 7 shows the pointer semantic version of the `for range`. The adding of the bonus on line 19 will now update the actual value in the slice, and it will show when printing the accounts on line 23.

Another solution to solve this bug is to use the read/modify/write pattern.

**Listing 8: Read, Modify, Write**  
```
17     for i, a := range acts {
18         if a.Type == "vip" {
19             a.Balance += 1_000
20             acts[i] = a
21         }
22     }
```

Listing 8 shows how to implement the  "read, modify ,write" pattern using the `for range` loop. On line 17, the code gets its own copy of the `Account` value stored in variable `a`. Then on line 19, the code adds the bonus to its local copy. Then on line 20, the code stores the local copy with the changes  into the slice.

This solution is great if you need to perform "transaction like" changes. You do several modifications on the local copy of the data, then you check those changes for validity, and only if everything is OK, you replace the original value in the data set with the changes.

### Bonus Map Range

What happens when you change the data set from being a slice (like in listing 6) to being a `map`?.

**Listing 9: Updating VIP Accounts**  
```
05 type Account struct {
06     Name    string
07     Type    string
08     Balance int
09 }
10 
11 func main() {
12     acts := map[string]Account{
13         "donald":  {"donald", "regular", 123},
14         "scrooge": {"scrooge", "vip", 1_000_001},
15     }
16 
17     for _, a := range acts {
18         if a.Type == "vip" {
19             a.Balance += 1_000
20         }
21     }
22
23     fmt.Println(acts)
23 }
```

Listing 9 shows an attempt to update a map value inside of a map iteration. On lines 12-15, the code declares the `acts` variable as a map with a key of type string representing a login name and a value of type `Account`. On lines 17-21,the codeI gives a bonus to every VIP member like you saw in listing 6.

This code has the same issue as in listing 6: The code is  updating a local copy of each account and the update won't be reflected back in the map.

You might be tempted to use the same solution as in listing 7.

**Listing 10: Trying Pointer Semantics**  
```
17     for k := range acts {
18         if acts[k].Type == "vip" {
19             acts[k].Balance += 1_000
20         }
21     }
```

Listing 10 shows an attempt to iterate over only the keys and reference each value in the map. However, this code does not compile, you will see the following error:

`./bank_map_err.go:19:4: cannot assign to struct field bank[k].Balance in map`.

If you try to use a reference (as in `(&acts[k]).Balance += 1_000`), it will fail as well with `invalid operation: cannot take address of acts[k]`.

The solution you can use is the read/modify/write pattern from listing 8.

**Listing 11: Using Read/Modify/Write with a Map**  
```
17     for k, a := range acts {
18         if a.Type == "vip" {
19             a.Balance += 1_000
20             acts[k] = a
21         }
22     }
```

Listing 11 shows you how to use the read/modify/write pattern to update a map in a range loop.

### Range Over Numbers

Something new in Go 1.22 is the ability to  `range` over an integer.

**Listing 12: Looping Over Integers**  
```
09     for i := range 3 {
10         fmt.Println(i)
11     }
```

Listing 12 shows a range over an integer. Line 09 is the equivalence of:

`for i := 0; i < 3; i++ {`

I found this new syntax handy for writing benchmark loops:

**Listing 13: Benchmark Loop**  
```
22 func BenchmarkEvent_Validae(b *testing.B) {
23     evt := NewEvent("elliot", "read", "/etc/passwd")
24
25     for range b.N {
26         err := evt.Validate()
27         if err != nil {
28             b.Fatal(err)
29         }
30     }
31 }
```

Listing 13 shows a benchmark. On line 24, the code loops `b.N` times using the new syntax instead of the old `for i := 0; i < b.N; i++`.

_NOTE: Go 1.22 also added `range over function` experiment, but this is a topic for another blog post._

### The goto Statement

The `for` statement is not the only looping construct in Go, there's also the `goto` keyword.
Looking at the source code for Go at version 1.22, you can see about 650 uses of `goto`:

**Listing 18: Number of `goto` in the Standard Library**  
```
01 $ find ~/sdk/go1.22.0/src -type f -name '*.go' -not -path '*test*' -not -name '*_test.go' | \
02     xargs grep -E 'goto\s+' | wc -l
03 657
```

Listing 18 shows you how to look for the `goto` keyword in the standard library. On line 01, you use the `find` utility to find non-test files in the Go 1.22 sources and then on line 02, you use `grep` and `wc` to find any line that is using `goto`.

Assuming some false positives, this is still a significant use of the `goto` keyword. Even considering [the dangers using goto](https://xkcd.com/292/),
it signals there are valid cases for using `goto`. However, if I see a `goto` in a code review, I will ask why. In my years of writing Go code, I haven't written a single `goto` statement in production code.

Let's change the event processing loop from listing 5 to use `goto`:

**Listing 19: Using `goto`**  
```
16 func logHandler(r io.Reader) error {
17     dec := json.NewDecoder(r)
18
19     for {
20         var log Log
21         err := dec.Decode(&log)
22
23         switch {
24         case errors.Is(err, io.EOF):
25             goto done
26         case err != nil:
27             return err
28         default:
29             fmt.Printf("log: %+v\n", log)
30         }
31     }
32 
33 done:
34     return nil
35 }
```

Listing 19 shows an example of using a `goto` instead of a label break. One line 31, the code defines a `done` label and then line 23 when there's no more data, the code jumps to the `done` label using a `goto` statement.

### Conclusion

There's much more to looping than just a traditional `for` loop. Next time you're about to start a loop, think about all the options you have in Go and pick the right one.

What looping idioms did I miss? Tell me at [miki@ardanlabs.com](mailto:miki@ardanlabs.com).
