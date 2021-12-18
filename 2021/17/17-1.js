/*
If the behavior of the probe on the y-axis is studied.
It can be observed that it passes through the same points on the way up and on the way down.
At some point, "y" will return to value 0, the next value will be -(initial "y" velocity + 1).

Knowing this, we can figure out that the velocity of the highest y will be the one that falls on the edge of the target area after passing through 0.
As the y target area range is -103...-58, a velocity of 102 will fall just right on the edge (-103).
We can calculate the highest value of its movement using the summatory function of 102.
*/
console.log((102 * 103)/2);
