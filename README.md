# go-dojo

Here I store exercises using the [Go programming language](https://golang.org/).
It has become a good source of Go examples too.

Most of the exercise ideas are taken from [Program Creek](http://www.programcreek.com/2012/11/top-10-algorithms-for-coding-interview/).
But the solutions to each exercise may not be exactly the same, as I am trying
to use Go's idioms. If you have any suggestion for a new exercise, or an
improvement on any solution (bug fix, better performance or better usage of
Go), please raise an issue or send a pull request. Thank you. However I do not
promise I will address the request in any specific period of time. This is just
for fun and personal training.

## Exercises

### Strings and arrays

#### Rotate an array

**Description:**
Rotate an array of `n` elements to the right by `k` steps.
For example, with `n = 7` and `k = 3`, the array `[1,2,3,4,5,6,7]` is rotated
to `[5,6,7,1,2,3,4]`.

**Solution**: [rotate](/rotate).

#### Evaluate Reverse Polish Notation

**Description:**
Evaluate the value of an arithmetic expression in Reverse Polish Notation.
Valid operators are +, -, `*`, / and ^ (power). Each operand may be an number or
an operator. For example:
* `["2", "1", "+", "3", "*"]` is `((2 + 1) * 3) = 9`
* `["4", "13", "5", "/", "+", "10", "*"]` is `(4 + (13 / 5)) * 10 = 66`
* `["2", "3", "^"]` is `2 ^ 3 = 8`

**Solution**: [rpn](/rpn).

#### Isomorphic strings

**Description:**
Given two strings `s` and `t`, determine if they are isomorphic. Two strings are
isomorphic if the characters in `s` can be replaced to get `t`.
For example, "egg" and "add" are isomorphic, but "foo" and "bar" are not.

**Solution**: [isomorphic](/isomorphic).

#### Word ladder

**Description:**
Given two words (`start` and `end`), and a dictionary, find the length of
shortest transformation sequence from `start` to `end`, such that only one
letter can be changed at a time and each intermediate word must exist in the
dictionary.

**Solution**: [wordladder](/wordladder).

#### Median of two sorted arrays

**Description:**
There are two sorted arrays `s` and `t` of size `m` and `n` respectively. Find
the median of the two sorted arrays. The overall run time complexity should be
`O(log(m+n))`.

**Solution**: [median](/median).

**Note**: The best solution I found (that really works! see the tests) runs in
`O(log(min(m,n)))` on average, and `O(log(m)+log(n))` in the worst case.

#### Merge intervals

**Description:**
Given a collection of intervals, merge all overlapping intervals.
For example:
* Given: `[1,3], [2,6], [8,10], [15,18]`
* Result: `[1,6], [8,10], [15,18]`.

**Solution**: [mergeint](/mergeint).
