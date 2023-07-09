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

## Functionality - How it works

### Data Model

A given sudoku riddle is stored as a matrix of numbers, represented by an array of an array of numbers (`var [9][9]int`)
The field is read from a json file, see [Input](#input)

### Helper functions

First of all let me introduce two important helper functions: `checkField()` and `possibleNumbers(x,y)`.
The `checkField()` function returns two values. A `boolean`, that tells if the given field is a valid sudoku 
game and is set to true if there is any violation of the games rules, like two 8 in a row. The method 
`possibleNumbers(x,y)` returns an array of numbers that contains all the numbers that could be inserted into the 
field at position (x,y) without violating the game rules.

### `solve()` function

This function is finally used to solve a sudoku. It is called recursively which means, that it can solve a sudoku 
at any stage by calling itself.

First of all the functions iterates over all the numbers in the sudoku and looks for the field with the smallest amount
of possible numbers, by calling `possibleNumbers(x,y)` for every field. If coordinates are found where the given number 
only has one option, the number at this position is replaced with the only possibility and the solve functions calls
itself again with the modified sudoku.

If after iterating over every field no field with only one possible option has been found, a second iteration over the fields starts.
This time the function looks for the first field with minimal amount of possible numbers to be placed there.
After this position is found it loops over all the possibilities for this field and for each option continues solving
with the `solve()` function.

The function returns the field after solving and a boolean that is `false` if solving the given field has failed.
At the point where the function is testing every possible option of a field it first assigns the first option
to this field and calls the `solve()` function with this modified field. If the solving algorithm returns `false`
it moves on to the next possible number for this field.

An example execution can be retraced when looking into the example log output mentioned unter [Output](#output) (`example_output_txt`)