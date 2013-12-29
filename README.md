go-fish - Yet another Go REPL (read, eval, print loop)
============================================================================

[![Build Status](https://travis-ci.org/rocky/go-fish.png)](https://travis-ci.org/rocky/go-fish)

This project provides an interactive environment for evaluating go
expressions.

Yeah, we know of some others. I think this one has more promise. The
heavy lifting for eval is provided by the Carl Chatfield's
[eval](https://github.com/0xfaded/eval) package.

Setup
-----

* Make sure our GO environment is setup, e.g. *$GOBIN*, *$GOPATH*, ...
* Make sure you have go a 1.2ish version installed.

```
   $ go get github.com/rocky/go-fish
   $ cd <place where go-fish got copied>
   $ make go-fish
   $ cp go-fish $GOBIN/
```

If you have
[go-gnureadline](https://code.google.com/p/go-gnureadline/) installed
and want the GNU readline support. In addition to the above try:

```
   $ make go-fish-grl
   $ cp go-fish-grl $GOBIN/
```

Using
-----

Run `go-fish` or `go-fish-rl`. For now, we have only a static
environment provided and that's exactly the environment that *eval*
uses for itself. (In other words this is ideally suited to introspect
about itself). Since the *eval* package is a reasonable size program,
many of the packages like *os*, *fmt*, *strconv*, *errors*, etc. are
available. Look at the import list in file *eval_imports.go* for the
exact list.

Two global variables have been defined: *env*, and *results*. *env*
the environment that is defined, again largely by
*eval_imports.go*. As you enter expresions, the results are saved in
slice *results*. To quit, enter `Ctrl-D` (EOF) or the word `quit`.

Here's a sample session:

```

$ ./go-fish
=== A simple Go eval REPL ===

Results of expression are stored in variable slice "results".
The environment is stored in global variable "env".

Enter expressions to be evaluated at the "go>" prompt.

To see all results, type: "results".

To quit, enter: "quit" or Ctrl-D (EOF).
go> 10 + len("abc" + "def")
Kind = Type = int
results[0] = 16
go> len(results)
Kind = Type = int
results[1] = 1
go> os.Args[0]
Kind = Type = string
results[2] = "./go-fish"
go> strings.HasPrefix(os.Args[0], "./")
Kind = Type = bool
results[3] = true
go> quit
```

See Also
--------

* [What's left to do?](https://github.com/rocky/go-fish/wiki/What%27s-left-to-do%3F)

[![endorse rocky](https://api.coderwall.com/rocky/endorsecount.png)](https://coderwall.com/rocky)
