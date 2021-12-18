const targetX = [265, 287];
const targetY = [-103, -58];

// [(n*(n+1))/2 < 265 (n < 23), targetX[1]]
// lower than 23 it does not horizontally reach the target area
const xVelRange = [23, 287];

//[targetY[0], result of 17-1]
const yVelRange = [-103, 102];

let count = 0;

const fallIntoTargetArea = (x, y) => {
    let pos = [0,0];
    let xVel = x;
    let yVel = y;
    while (true) {
        pos = [pos[0] + xVel, pos[1] + yVel];

        if (pos[0] >= targetX[0] && pos[0] <= targetX[1] && pos[1] >= targetY[0] && pos[1] <= targetY[1]) {
            return true;
        } else if (pos[0] > targetX[1] || pos[1] < targetY[0]) {
            return false;
        }

        if (xVel > 0) xVel--;
        yVel--;
    }
}

for (let x = xVelRange[0]; x <= xVelRange[1]; x++) {
    for (let y = yVelRange[0]; y <= yVelRange[1]; y++) {
        if (fallIntoTargetArea(x, y)) {
            count ++;
        }
    }
}

console.log(count);





