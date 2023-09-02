# Multi-part tutorial
<!-- TOC -->
* [Multi-part tutorial](#multi-part-tutorial)
  * [Step 1: Create a Go module](#step-1--create-a-go-module)
  * [Step 2: Call our code from another module](#step-2--call-our-code-from-another-module)
<!-- TOC -->

## Step 1: Create a Go module
Source: [Link](https://go.dev/doc/tutorial/create-module)

1. Create `greetings` dir for our Go module source code:
```
$ mkdir greetings
$ cd greetings
```

2. Start our module using the `go mod init` command:
```
$ go mod init greetings
go: creating new go.mod: module greetings

$ ls
go.mod

$ cat go.mod
module greetings

go 1.19
```
The `go mod init <....>` command creates a `go.mod` file to track our code's dependencies (currently, includes name of our module and Go version). 

3. Create `greetings.go` file with content provided in source link

## Step 2: Call our code from another module
Source: [Link](https://go.dev/doc/tutorial/call-module-code)
1. Create `hello` dir, where we will write our caller
```
$ cd ..
$ mkdir hello
$ cd hello
```
After we create this dir, we should both `greetings` and `hello` dir at same level in hierarchy:
```
<working_dir>/
├── greetings/
│   ├── go.mod
│   └── greetings.go
└── hello/
```
2. Enable dependency tracking for the code we are about to write
```
$ go mod init hello
go: creating new go.mod: module hello

$ ls
go.mod

$ cat go.mod
module hello

go 1.19
```
3. Create `hello.go` file with content provided in source link
4. Edit the `go.mod` in `hello` folder to use our local `greetings` module.

Redirect Go tools from its module path to local dir:
```
$ go mod edit -replace greetings=../greetings
$ cat go.mod
module hello

go 1.19

replace greetings => ../greetings
```
To synchrpnise the `hello` module's dependencies, adding those required by code, but not yet tracked in the module:
```
$ go mod tidy
go: found greetings in greetings v0.0.0-00010101000000-000000000000
$ cat go.mod
module hello

go 1.19

replace greetings => ../greetings

require greetings v0.0.0-00010101000000-000000000000
```
The `go mod tidy` command adds missing and removes unused modules/require statements.

The number following the module path is a *pseudo-version* number. It is a unique identifier for a specific commit in a module's history. The `go.mod` file uses the pseudo-version to identify a specific version of a module.

5. Run our code
```
$ go run hello.go # go run . also works
Hi, Gladys. Welcome!
```