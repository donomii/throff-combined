[![GoDoc](https://godoc.org/github.com/donomii/throff?status.svg)](https://godoc.org/github.com/donomii/throff)

# throff

A programming language that started as an improved Forth, and got out of control.  Has lexical scoping, continuations, effective tail calls and macros.  throff runs on a custom 3-stack engine. 

throff was an experiment in creating a small core and building the language up to a fully featured programming language.  This was successful, but there are already meany fully featured programming languages and this one doesn't add anything particularily new.

View the [manual](throffmain/README.md).

## Get it

the Throff programming language

    go get -v github.com/donomii/throff
    go build github.com/donomii/throff

 Or download a [precompiled binary](https://github.com/donomii/throff/releases)

## Use it

### Run from Command Line

	> throff ADD 1 1
	2

### Interactive shell

	> throff
	Welcome to the THROFF command shell v0.1. Type HELP for help.
	Throff »ADD 1 1
	2

### Run a program from a file

	throff -f sierpinski.thr

### Scripting interface for golang

``` golang
	result, _ := t.CallArgs1("ADD", 4, "5")
	fmt.Println("Calculation result: ", result)
```

## Throff is

Throff is a dynamically typed, late binding, homoiconic, concatenative programming language, taking inspiration from Forth, Joy and Scheme.  It has all the features of a modern language - [closures, lexical scopes](http://praeceptamachinae.com/post/throff_variables.html), [tail call optimisations](http://praeceptamachinae.com/post/throff_tail_call_optimisation.html), currying, and continuations.

Everything is a function, even language constructs like IF and FOR, which can be replaced and extended with your own versions.  It uses immutable semantics wherever possible to provide safe threading and continuations.  There is almost no lexer/tokeniser, and no parser in the traditional sense.  Commands are fed directly into the engine to be executed.  The programs are written _backwards_. 

There are many examples in the [examples](examples/) and [rosetta](rosetta/) directories.

