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
// Aaa bbb ccc ddd eee fff ggg hhh iii jjj kkk lll mmm nnn ooo ppp qqq rrr sss ttt uuu vvv www xxx yyy zzz.
//
//     if err != nil {
//             return err
//     }
//
// Aaa bbb ccc ddd eee fff ggg hhh iii jjj kkk lll mmm nnn ooo ppp qqq rrr sss ttt uuu vvv www xxx yyy zzz.
```

**After**:

``` go
// Aaa  bbb  ccc  ddd eee fff ggg hhh iii jjj kkk lll mmm nnn ooo ppp qqq
// rrr sss ttt uuu vvv www xxx yyy zzz.
// 
//     if err != nil {
//             return err
//     }
// 
// Aaa  bbb  ccc  ddd eee fff ggg hhh iii jjj kkk lll mmm nnn ooo ppp qqq
// rrr sss ttt uuu vvv www xxx yyy zzz.
```

## Todos

* Currently supports `FuncDecl`, need to support more.