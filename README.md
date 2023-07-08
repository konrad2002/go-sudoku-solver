# go-sudoku-solver

This project contains a small go script that is able to solve easy sudoku riddles.
If the riddle is too hard, it won't be able to find a solution.

## Input

The file `field.json` provides an example input for a sudoku.
In general, you'll have to pass a matrix of size 9x9 which has to be represented as an array of an array of numbers.
Empty fields are filled with a zero.
For example:

```json
[
  [5,3,0,0,0,4,7,0,0],
  [0,0,0,0,0,7,0,3,0],
  [0,0,0,3,6,0,0,0,1],
  [6,7,0,0,1,0,0,4,9],
  [0,0,0,9,0,6,0,0,0],
  [4,5,0,0,8,0,0,2,6],
  [2,0,0,0,7,9,0,0,0],
  [0,6,0,5,0,0,0,0,0],
  [0,0,3,6,0,0,0,1,4]
]
```

## Output

The programm prints the given field after running the solving algorithm on it.
If it's not completely solved, you'll find the number of missing fields above it.

Given example from above returns:

```
---==[ CURRENT FIELD ]==---
|5|3|2|1|9|4|7|6|8|
|1|4|6|8|5|7|9|3|2|
|9|8|7|3|6|2|4|5|1|
|6|7|8|2|1|5|3|4|9|
|3|2|1|9|4|6|8|7|5|
|4|5|9|7|8|3|1|2|6|
|2|1|5|4|7|9|6|8|3|
|8|6|4|5|3|1|2|9|7|
|7|9|3|6|2|8|5|1|4|
```