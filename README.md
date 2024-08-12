## canjump
This program determines if it's possible to reach and stay at the last index of the array
starting from the first index, based on the steps you need to advance. Each value of the 
slice represents the exact number of steps you must take forward from that position.


## How the program works
For example; if the program is given the
unsigned integer slice:
 []uint{2, 3, 1, 1, 4}

then:
Position: 0  1  2  3  4
Steps:    2  3  1  1  4
          ^

// Starting from position 0, you have 2 steps to move forward. This means you will move to positions 2.

Position: 0  1  2  3  4
Steps:    2  3  1  1  4
                ^

// From position 2, you have 1 step, so you will move to position 3.

Position: 0  1  2  3  4
Steps:    2  3  1  1  4
                   ^

// Finally, from position 3, you have 1 step to reach the last index at position 4 confirming that it's possible so the output will be "True".

Position: 0  1  2  3  4
Steps:    2  3  1  1  4
                      ^


## Author
- [Barrack Kope Otieno](https://www.linkedin.com/in/barrack-kope-otieno-064a43244)