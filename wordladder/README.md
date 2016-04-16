# Word ladder
## Description of the problem

Given two words (`start` and `end`), and a dictionary, find the length of
shortest transformation sequence from `start` to `end`, such that only one
letter can be changed at a time and each intermediate word must exist in the
dictionary.

## Solution

### Data structures

To solve this problem three data structures are used:

1. `WordSet` represents a set of words, implemented as a hash table to provide
   O(1) searches. It is used for two purposes: first, to represent the
   dictionary of known words; and second to keep track of the words that have
   already been visited, in order to avoid infinite loops.
2. `wordLadder` is a linked list of words, starting from the last visited word
   and ending on the `start` word. This approach has two results: on the good
   side, when visiting a new word, we can add it in the ladder by inserting it at
   the beginning of the linked list in O(1). On the bad side, when we finally get
   to our `end` word, we need to list the linked list in reverse order to
   correctly show the transformation sequence. This is done in the `climbLadder`
   method.
3. A `queue` of `wordLadders` is used to implement a Breadth-First Search (BFS)
   on the graph of all the possible conversions beginning from the `start`
   word. Doing a BFS ensures we will find on of the shortest transformations
   first, as requested by the problem.

### Summary on how the process works

The process starts with a single `wordLadder` in the queue. This ladder is
called the `root` and only contains the `start` word, which is marked as
visited right away.

Then, for each `wordLadder` in the queue, the last word in it (that is, the
first word in the linked list) is transformed in all the possible ways. This
means each character in the word is changed for any other character in the
English alphabet. Note this means that `len(word) * 25` new words are tried at
this stage. For each of these potential new words: If the `end` word is found,
the process ends by *climbing* the ladder. Otherwise, if the new word has not
been visited yet and it exists in the dictionary, it is marked as visited,
added at the end of the ladder (that is, at the beginning of the linked list)
and the `wordLadder` is enqueued back again at the end of the queue for another
iteration later in the BFS.
