# GoLang Learning Repository

This repository contains a collection of Go programs and projects that I created while learning the Go programming language. The programs are organized in a sequence that progressively covers various aspects of Go, from basic syntax to advanced topics such as concurrency and web development.

## Table of Contents

1. [Introduction](#introduction)
2. [Folder Structure](#folder-structure)
3. [Individual Programs](#individual-programs)
4. [Projects](#projects)
   - [App](#app)
   - [App2](#app2)
   - [WebApp](#webapp)
5. [How to Run](#how-to-run)
6. [Credits](#credits)

## Introduction

This repository serves as a personal reference and a showcase of what I have learned in Go. It includes simple programs demonstrating basic concepts, as well as more complex projects that involve file I/O, concurrency, and web development.

## Folder Structure

```
.
├── app
│ ├── go.mod
│ ├── main.go
│ └── mypackage
│ └── mypackage.go
├── app2
│ ├── go.mod
│ ├── testemail.go
│ └── testemail_test.go
├── learning
│ ├── 00-hello-go.go
│ ├── 01-user-input.go
│ ├── 02-variables.go
│ ├── 03-data-types.go
│ ├── 04-type-casting.go
│ ├── 05-if-conditions.go
│ ├── 06-strings.go
│ ├── 07-runes.go
│ ├── 08-time.go
│ ├── 09-random-values.go
│ ├── 10-math-functions.go
│ ├── 11-formatted-print.go
│ ├── 12-for-loop.go
│ ├── 13-while-loop.go
│ ├── 14-guessing-game.go
│ ├── 15-range.go
│ ├── 16-array.go
│ ├── 17-slicing.go
│ ├── 18-functions.go
│ ├── 19-varadic-functions.go
│ ├── 20-array-parameter.go
│ ├── 21-pass-by-value.go
│ ├── 22-defined-types-01.go
│ ├── 22-pointers.go
│ ├── 23-pass-by-reference.go
│ ├── 24-file-io.go
│ ├── 25-append-file.go
│ ├── 26-command-line
│ ├── 26-command-line.go
│ ├── 27-maps.go
│ ├── 28-generics.go
│ ├── 29-structs.go
│ ├── 30-structs-member-func.go
│ ├── 31-struct-composition.go
│ ├── 32-defined-types-01.go
│ ├── 33-defined-types-02.go
│ ├── 34-interfaces.go
│ ├── 35-concurrency-goroutines.go
│ ├── 36-goroutines-channels.go
│ ├── 37-mutex-lock.go
│ ├── 38-closures.go
│ ├── 39-recursion.go
│ ├── 40-regular-expression.go
│ ├── data.txt
│ └── Readme.md
└── webapp
├── new.html
├── todos.txt
├── view.html
└── webapp.go
```


## Individual Programs

Each file in the root directory (e.g., `00-hello-go.go`, `01-user-input.go`) is a standalone Go program that demonstrates a specific concept or feature of the Go language. These include:

- Basic syntax and structure
- Variables and data types
- Control structures (if, loops)
- Functions and methods
- Arrays, slices, and maps
- Pointers and memory management
- Concurrency with Goroutines and Channels
- Error handling and file I/O
- Custom types and interfaces

## Projects

### App

The `app` folder contains a project that demonstrates the use of pointers, pass by reference, and handling date/time errors in Go. This project is designed to deepen the understanding of these concepts by applying them in a practical scenario.

### App2

The `app2` folder contains a project focused on testing in Go. It includes examples of automated tests written using Go's testing package. This project showcases how to write and run tests to ensure code reliability and correctness.

### WebApp

The `webapp` folder contains a small web application project that demonstrates how to build a simple To-Do list application in Go. It includes basic HTML templates and Go code to handle web requests, showcasing Go's capabilities for web development.

## How to Run

To run any of the individual programs or projects, navigate to the respective directory and use the Go command:

```sh
go run filename.go
```

For the projects in the app, app2, and webapp folders, you may need to use go build to compile the application and then run the executable:

```sh
go build main.go
./main
```

To run tests in the app2 folder:
```sh
go test
```

## Credits

This learning journey was inspired by Derek Banas' excellent YouTube tutorials on Go. If you're interested in learning more about Go or other programming languages, I highly recommend checking out his content.

- **YouTube Channel**: [Derek Banas](https://www.youtube.com/@derekbanas)
- **Go Tutorial Video**: [Go Programming Tutorial](https://youtu.be/YzLrWHZa-Kc?si=vBWZiOU7qJfhnMmz)
