# Isomorphic strings
## Description of the problem

Given two strings `s` and `t`, determine if they are isomorphic. Two strings are
isomorphic if the characters in `s` can be replaced to get `t`.
For example, "egg" and "add" are isomorphic, but "foo" and "bar" are not.

## Solution

The strategy to solve this problem is to keep a map of characters from `s` to
`t`. The process scans each character in `s`, along with its counterpart in `t`:
1. When a character that has not been seen before in `s` is found, it is mapped
   to the corresponding character in `t` (the character in `t` in the same
   position).
2. But when a character that had already been seen before in `s` is found, the
   process checks that the previously mapped character is equal to the
   corresponding character in `t`. If the characters are equal, all is fine and
   the process continues. Otherwise, the process are not isomorphic.

If the entire string `s` is processed and both `s` and `t` have the same
length, then they are isomorphic.

This implementation supports UTF-8 strings.
