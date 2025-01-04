# Tic-Tac-Toe

## Overview

This is a simple implementation of the classic Tic-Tac-Toe game in Golang. The game allows two players to play against each other on a 3x3 grid.

## Features

- Two-player mode
- Simple command-line interface
- Input validation
- Win detection
- Draw detection

## Installation

To install and run the game, follow these steps:

1. Clone the repository

2. Navigate to the project directory:
    ```sh
    cd ticTacToe
    ```
3. Build the project:
    ```sh
    go build
    ```
4. Run the executable:
    ```sh
    ./ticTacToe
    ```

## How to Play

1. The game starts with an empty 3x3 grid.
2. Players take turns to enter their moves.
3. Enter the row and column number to place your mark (X or O) on the grid.
4. The game checks for a win or a draw after each move.
5. The game ends when a player wins or the grid is full (draw).

## Example

```
Player 1 (X) - Enter row and column: 1 1
 X |   |  
-----------
   |   |  
-----------
   |   |  

Player 2 (O) - Enter row and column: 1 2
 X | O |  
-----------
   |   |  
-----------
   |   |  
```

