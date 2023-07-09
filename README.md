# go-sudoku-solver

This project contains a small go script that is able to solve sudoku riddles.

## Input

The file `field.json` provides an example input for a sudoku.
In general, you'll have to pass a matrix of size 9x9 which has to be represented as an array of an array of numbers.
Empty fields are filled with a zero.
For example:

```json
[
  [0,9,0,0,5,0,3,4,0],
  [1,4,0,6,0,0,8,0,0],
  [0,0,0,8,0,1,0,0,0],
  [0,0,0,0,3,0,0,6,9],
  [0,0,0,0,0,0,0,0,0],
  [8,2,0,0,1,0,0,0,0],
  [0,0,0,3,0,4,0,0,0],
  [0,0,6,0,0,9,0,7,5],
  [0,8,2,0,6,0,0,3,0]
]
```

## Output

The programm prints the given field after running the solving algorithm on it.
If it's not completely solved, you'll find the number of missing fields above it.

Given example from above returns:

```
---==[ CURRENT FIELD ]==---
┏6│9│8┃7│5│2┃3│4│1┓
┃1│4│5┃6│9│3┃8│2│7┃
┗2│7│3┃8│4│1┃5│9│6┛
┏5│1│7┃4│3│8┃2│6│9┓
┃3│6│9┃5│2│7┃4│1│8┃
┗8│2│4┃9│1│6┃7│5│3┛
┏9│5│1┃3│7│4┃6│8│2┓
┃4│3│6┃2│8│9┃1│7│5┃
┗7│8│2┃1│6│5┃9│3│4┛
```