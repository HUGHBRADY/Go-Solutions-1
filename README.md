# Go Solutions 1
This repository contains solutions to Go Problem Sheet 1, authored by Hugh Brady, a student in GMIT. This sproblem sheet is for the purpose of learning the Go programming language as part of our module, Data Representation & Querying.

## Running the Code
To run any Go programs on your machine you must download the [Go compiler](https://golang.org/dl/).

Once Go is installed, open your command prompt and follow the steps below.

1) Clone this repository using Git.
> git clone https://github.com/HUGHBRADY/Go-Solutions-1

2) Change directory into the src folder
> cd Go-Solutions-1/src

3) Run any of the files you wish
> go run <filename.go>



## Go Problem Sheet 1
Below are the actual problems assigned to us.


### 1. Kon’nichiwa, Sekai!

Source: https://tour.golang.org/welcome/1

Write a program that prints Hello, world! in Japanese (using Japanese characters) to the screen.


### 2. Current time

Source: https://tour.golang.org/welcome/4

Write a program that prints the current time and date to the console.


### 3. FizzBuzz

Source: http://wiki.c2.com/?FizzBuzzTest

Write a program that prints the numbers from 1 to 100, each on a new line, to the console, except for the following conditions. For multiples of three print Fizz instead of the number, and for multiples of five print Buzz. For numbers that are multiples of both three and five print FizzBuzz.


### 4. Factorial digit sum

Write a function that calculates the sum of the digits of the factorial of a number. n! means n x (n − 1) ... x 3 x 2 x 1. For example, 10! = 10 x 9 x ... x 3 x 2 x 1 = 3628800, and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27. Then find the sum of the digits in the number 100!.


### 5. Guessing game

Source: https://adriann.github.io/programming_problems.html

Write a guessing game where the user must guess a randomly generated number. After every guess the program tells the user whether their number was too high or too low. At the end, the number of tries should be printed. It counts only as one try if they input the same number multiple times consecutively.


### 6. Largest and smallest in list

Source: https://adriann.github.io/programming_problems.html

Write a function that returns the largest and smallest elements in a list.


### 7. Palindrome test

Source: https://adriann.github.io/programming_problems.html

Write a function that tests whether a string is a palindrome. A palindrome is a string that reads the same in reverse, e.g. radar.


### 8. Merge list and sort

Source: https://adriann.github.io/programming_problems.html

Write a function that merges two sorted lists into a new sorted list, e.g. merge([1,4,6], [2,3,5]) = [1,2,3,4,5,6].


### 9. Newton’s method for square roots

Source: https://tour.golang.org/flowcontrol/8

Write a function to calculate the square root of a number using Newton’s method. Newton’s method is to approximate the square root of a number x by picking a starting point z and then repeating the following operation.

z_next = z - ((z*z - x) / (2 * z))
This is repeated while the values of z_next and z are different. After each iteration the value of z is replaced with that of z_next.

Run a few tests to determine how close you are to the math.Sqrt value?

Hint: to declare and initialize a floating point value, give it floating-point syntax or use a conversion:

z := float64(1)
z := 1.0


### 10. Reverse string

Write a function to reverse a string in Go.
