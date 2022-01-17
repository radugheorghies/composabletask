# How I approach the solution

For this problem, I took into consideration that, at a specific moment, we need to analyze at least 4 elements and maximum 5 elements from the slice A.

From the optimization perspective, I used a linked list where I put those 4 elements. At any moment, my list has only 4 or 5 elements. In this way, I'll have the result with only one iteration thru elements. The time complexity is O(n) and the space complexity is also O(n). If the data is available on a stream, the space complexity will be 1, because I can have a maximum of 5 elements at a specific moment in my list.

If the list has only 4 elements, I choose what element must be cut, I eliminate it from the list, and I load the next element from A to the back of the list.
If the elements in the list don't need to be cut, I remove the first element from the list and I load the next element from A to the back of the list.
If the list contains 5 elements, it means that I was in a position to cut one of the elements in the middle, and in this situation, I removed the cut element and the first from the list, and I loaded the next element from A to the back of the list.

The process is repeating until I finish consuming elements from A.

At any momment the first 4 elements in my  list could exist in only one of the following situations:

e1 < e2 < e3 < e4 => this will fail <br />
e1 < e2 < e3 = e4 => wrong dataset (2 equals) <br />
e1 < e2 < e3 > e4 => this should be analyzed <br />

e1 < e2 = e3 < e4 => wrong dataset (2 equals) <br />
e1 < e2 = e3 = e4 => wrong dataset (2 equals) <br />
e1 < e2 = e3 > e4 => wrong dataset (2 equals) <br />

e1 < e2 > e3 < e4 => this is the ideal situation, continue <br />
e1 < e2 > e3 = e4 => wrong dataset (2 equals) <br />
e1 < e2 > e3 > e4 => this should be analyzed <br />

....... <br />

e1 > e2 > e3 > e4 # this will fail <br />

So we have 19 different situations (9 for e1 > e2, 1 for e1 = e2 and 9 for e1 < e2), whith some of them in the mirror.

### Please see the comments in the code
