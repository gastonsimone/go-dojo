# Evaluate Reverse Polish Notation
## Description of the problem

Evaluate the value of an arithmetic expression in Reverse Polish Notation.
Valid operators are +, -, `*`, / and ^ (power). Each operand may be an number
or an operator. For example:
* `["2", "1", "+", "3", "*"]` is `((2 + 1) * 3) = 9`
* `["4", "13", "5", "/", "+", "10", "*"]` is `(4 + (13 / 5)) * 10 = 66`
* `["2", "3", "^"]` is `2 ^ 3 = 8`

## Solution

The strategy to solve this problem is to keep an *operands stack*, which begins
empty, and then start consuming one token at a time:

1. If the next token is a number, it is pushed onto the stack.
2. If the next token is an operator, as many numbers (operands) as needed by
   the operator are popped out from the stack, the calculation is done and the
   result is pushed onto the stack. If the stack does not contain enough operands
   for the operator, an error is returned.

When all the tokens are processed the final result is on top of the stack.
