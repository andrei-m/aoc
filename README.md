Advent of Code
==============

My solutions to [Advent of Code](https://adventofcode.com/).

Solutions are organized by year and day. Each day is a runnable Go program (with a `main()` method) that accepts the day's input file from stdin.

Example:
```
# in the 2024/4 folder:
$ cat input.txt | go run xmas.go
```

Some solutions contain optional debug logging that can be enabled using the `DEBUG` environment variable. Example:

```
$ cat input.txt | DEBUG=1 go run xmas.go
```
