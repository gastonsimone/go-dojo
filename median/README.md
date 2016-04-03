# Median of two sorted arrays
## Description of the problem

There are two sorted arrays `s` and `t` of size `m` and `n` respectively. Find
the median of the two sorted arrays. The overall run time complexity should be
`O(log(m+n))`.

## Solution

The best solution I found (that really works! see the tests) runs in
`O(log(min(m,n)))` on average, and `O(log(m)+log(n))` in the worst case.

The idea is to scan only one of the arrays (slices in Go), using a binary
search, looking for the overall median. In every step we can calculate how many
elements are considered greater than the candidate median in this array. Then
we can calculate exactly how many elements must be greater than the candidate
median in the other array. Since the other array is also sorted, we only need
to verify two positions of the other array to confirm this.

For example, when scanning the array `s` the first guess is to say the element
in the middle of `s` is the overall median. In this case `len(s)/2` elements
are greater than this candidate in as (for `s` is sorted). If `N = len(s)
+ len(t)` is the total number of elements, we can say that this candidate
median must be greater than `c = T/2 - len(s)/2` elements from the array `t`.
Therefore, we only need to verify positions `i = len(t) - c` and `j = len(t)
- c - 1` in `t`, if they exist. The candidate is indeed the overall median if
and only if `t[i] <= candidate <= t[j]`.

### Optimization

In the cases where `s` is reasonable bigger than `t` there is a significant
improvement. We know the overall median will be close to the median of `s`.
Therefore there is no need to analyse all the elements of `s`, but only those
around its own median. If all the elements in `t` are greater (or lower) than
the median of `s`, then the overall median is lower (or higher) than the median
of `s` by `len(t)/2` elements. Hence, we only need to analyse `len(t)/2`
elements to the left and to the right of the median of `s`. This is implemented
in the `medianScan` function.

### Concurrency

Since the approach to solve this problem is to scan one of the arrays, and then
scan the other (only if the first scan did not succeed) there is a good
opportunity for concurrency here: Both arrays can be scanned concurrently,
using *goroutines* and channels. We only need the result of the first one that
finishes successfully. I tried this idea, but it was **30x slower** than the
non-concurrent approach. So I discarded it.
