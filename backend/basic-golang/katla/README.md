# Katla

- Please try this first: <https://katla.vercel.app/arsip>
- [Important] Give students chance to try themselves until they really understand the game rule
- Please explain step by step, slowly
  - Start from zero code and iteratively add the solution

[![Introduction Video](https://img.youtube.com/vi/aqks78z2nSw/maxresdefault.jpg)](https://youtu.be/aqks78z2nSw)

## First challenge: Katla Engine

[![Katla Engine](https://img.youtube.com/vi/h3kcr8CTMNo/maxresdefault.jpg)](https://youtu.be/h3kcr8CTMNo)

- This problem is meant for tutorial, please explain the step by step solution to the students
- `engine/main.go` is CLI based Katla Engine
- The engine will randomly choose a word from the dictionary and ask the user to guess the word
- You have 6 chances to guess the word
  - First, type the word you think it is
  - The engine will return the hints (see Katla's rules):
    - "G": Green, correct position
    - "Y": Yellow, correct word but wrong position
    - "X": Grey, wrong word
- Enjoy the game :)
- In the next lesson, when we would have learned HTTP server, we will learn how to render nicely on the browser :)
## Second challenge: Katla Solver

<!-- TODO: answer here -->

- This problem is meant for challenge, let the students try themselves first
- `solver/main.go` is Katla Solver
- The solver will help us eliminate words from the dictionary and return the remaining possible words