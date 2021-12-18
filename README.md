# go-functools

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/ShauryaAg/go-functools/Go)

Golang package containing **functools** using Go generics

`functools` is a set of functions for functional programming in Golang.

It contains the following methods:

- **arrayEquals**: checks if two arrays are equal
- **copyWithin**: copy a value from target to destination within the array
- **every**: check if every element satisfies the given condition
- **fill**: fill the given range with the given value
- **filter**: pick the values that satisfy the given condition
- **find**: find the first value that satisfies the condition
- **findIndex**: find the index of the first value that satisfies the condition
- **forEach**: run a given function for all the values in the array
- **includes**: check if the given value is present in the array
- **join**: join the array as a string using a delimiter
- **map**: run a given method for all the elements and return the result
- **reduce**: reduce the array into a single value using a given condition
- **reverse**: reverse the array
- **shift**: remove the first element from the array
- **some**: check if some element(s) of the array satisfies the condition
- **sort**: sort the array
- **unshift**: add values in front of the array

## Todo:

- [ ] Add tests
- [ ] Add tools for `map`
- [ ] Add more efficient implementations
