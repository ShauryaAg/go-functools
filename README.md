# go-functools

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/ShauryaAg/go-functools/Go)

Golang package containing **functools** using Go generics

`functools` is a set of functions for functional programming in Golang.

It contains the following methods:

- [`arrayEquals`](tools/arrayEquals.go): checks if two arrays are equal
- [`copyWithin`](tools/copyWithin.go): copy a value from target to destination within the array
- [`every`](tools/every.go): check if every element satisfies the given condition
- [`fill`](tools/fill.go): fill the given range with the given value
- [`filter`](tools/filter.go): pick the values that satisfy the given condition
- [`find`](tools/find.go): find the first value that satisfies the condition
- [`findIndex`](tools/findIndex.go): find the index of the first value that satisfies the condition
- [`forEach`](tools/forEach.go): run a given function for all the values in the array
- [`includes`](tools/includes.go): check if the given value is present in the array
- [`join`](tools/join.go): join the array as a string using a delimiter
- [`map`](tools/map.go): run a given method for all the elements and return the result
- [`reduce`](tools/reduce.go): reduce the array into a single value using a given condition
- [`reverse`](tools/reverse.go): reverse the array
- [`shift`](tools/shift.go): remove the first element from the array
- [`some`](tools/some.go): check if some element(s) of the array satisfies the condition
- [`sort`](tools/sort.go): sort the array
- [`unshift`](tools/unshift.go): add values in front of the array

## Todo:

- [ ] Add tests
- [ ] Add tools for `map`
- [ ] Add more efficient implementations
