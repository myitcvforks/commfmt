# commfmt
A standard for Go comments.

## Installation

```
$ go get -u github.com/codingconcepts/commfmt
```

## Usage

```
$ commfmt -h
Usage of commfmt:
  -path string
        the relative/absolute path of the root directory. (default ".")
  -width int
        the maximum width of comments. (default 80)
```

## Example

Here's an example of running the commfmt tool against a directory containing the following code block:

```
$ commfmt -path . -width 80
```

**Before**:

``` go
// Aaa aaa aaa aaa aaa aaa aaa aaa aaa aaa aaa aaa aaa aaa aaa aaa.
//
//     if err != nil {
//             return err
//     }
//
// Bbb bbb bbb bbb bbb bbb bbb bbb bbb bbb bbb bbb bbb bbb bbb bbb.
```

**After**:

``` go
// Aaa  aaa aaa aaa aaa aaa aaa aaa aaa aaa
// aaa aaa aaa aaa aaa aaa.
// 
//     if err != nil {
//             return err
//     }
// 
// Bbb  bbb bbb bbb bbb bbb bbb bbb bbb bbb
// bbb bbb bbb bbb bbb bbb.
```

## Todos

* Fix additional commented line when last paragraph is a code block.
* Currently supports `FuncDecl`, need to support more.
* Consider /**/ code blocks.