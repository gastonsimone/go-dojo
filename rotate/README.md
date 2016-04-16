# Rotate an array
## Description of the problem

Rotate an array of `n` elements to the right by `k` steps.
For example, with `n = 7` and `k = 3`, the array `[1,2,3,4,5,6,7]` is rotated
to `[5,6,7,1,2,3,4]`.

## Solution

This implementation uses an already famous and efficient technique to rotate
arrays. If you want to rotate an array `k` steps, you can:

1. Take `s := len(s) - k` as the *split position*.
2. Reverse the first `s` positions of the array as a sub-array.
3. Reverse the rest of the array as another sub-array.
4. Reverse the entire array.

This solution supports rotations *to the right* with `k > 0` and rotations *to
the left* with `k < 0`.

### Efficiency

All the movements and swapping are done in place. No auxiliary copies of the
array are done in memory. So, in terms of memory consumption, the algorithm is
O(1).

No matter which value of `k` is used, this implementation does the minimum
number of necessary movements to get the requested result, by taking `k modulo
len(s)`. Therefore, the algorithm is O(n), regardless of the value of `k`.
