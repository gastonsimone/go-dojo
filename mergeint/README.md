# Merge Intervals
## Description of the problem

Given a collection of intervals, merge all overlapping intervals.
For example:
* Given: `[1,3], [2,6], [8,10], [15,18]`
* Result: `[1,6], [8,10], [15,18]`.

## Solution

For this problem I have provided two solutions.

### Sorting first

This solution is implemented in the `SortFirst` function. The method begin by
sorting the set of intervals by their `start` value, as the first step in the
process. Then it iterates on each interval in the set, checking if it can be
merged with the last merged interval, or it can simply be added as part of the
final result. The initial sort ensures only the last merged interval needs to
be checked for overlapping, which reduces the nomber of comparisons that need
to be performed.

### No sorting

This solution is implemented in the `NoSort` function. This solution does not
need to sort the intervals before merging, but it needs to check all the
already-merged intervals with every new interval it has to process (this is why
it has a nested `for` loop). It also removes and re-inserts elements in the
merged set as it finds more intervals that can be merged.

### Comparison of both solutions

First of all, the sorting-first solution is much simpler, as it can be seen in
the source code.

However, it seems the initial sort produces a bigger work than the extra
overhead needed in the no-sorting solution. In my tests (see the benchmarks in
[merge\_test.go](./merge_test.go)) the no-sorting solution is between 30% and
40% faster than the sorting-first solution.
