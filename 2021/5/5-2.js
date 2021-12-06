const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const vectors = data.split('\n')
        .filter(line => line.length)
        .map(parseVector);

    const grid = Array(1000).fill().map(() => Array(1000).fill(0));

    vectors.forEach(vector => {
        if (vector.point1.x === vector.point2.x) {
            const lower = vector.point1.y < vector.point2.y ? vector.point1.y : vector.point2.y;
            const greater = vector.point1.y < vector.point2.y ? vector.point2.y :  vector.point1.y;
            for (let i = lower; i <= greater; i++) {
                grid[vector.point1.x][i] += 1;
            }
        } else if (vector.point1.y === vector.point2.y) {
            const lower = vector.point1.x < vector.point2.x ? vector.point1.x : vector.point2.x;
            const greater = vector.point1.x < vector.point2.x ? vector.point2.x :  vector.point1.x;
            for (let i = lower; i <= greater; i++) {
                grid[i][vector.point1.y] += 1;
            }
        }  else {
            if (vector.point1.x < vector.point2.x && vector.point1.y < vector.point2.y) {
                for (let i = 0; i <= vector.point2.x - vector.point1.x; i++) {
                    grid[vector.point1.x + i][vector.point1.y + i] += 1;
                }
            } else if (vector.point1.x < vector.point2.x && vector.point1.y > vector.point2.y) {
                for (let i = 0; i <= vector.point2.x - vector.point1.x; i++) {
                    grid[vector.point1.x + i][vector.point1.y - i] += 1;
                }
            }
            else if (vector.point1.x > vector.point2.x && vector.point1.y > vector.point2.y) {
                for (let i = 0; i <=  vector.point1.x - vector.point2.x; i++) {
                    grid[vector.point1.x - i][vector.point1.y - i] += 1;
                }
            } else if (vector.point1.x > vector.point2.x && vector.point1.y < vector.point2.y) {
                for (let i = 0; i <=  vector.point1.x - vector.point2.x; i++) {
                    grid[vector.point1.x - i][vector.point1.y + i] += 1;
                }
            }
        }
    })

    console.log(grid.flat().filter(x => x > 1).length);

});

const parseVector = (line) => {
    const points = line.split(' -> ');
    const point1 = points[0].split(',');
    const point2 = points[1].split(',');
    return {
        point1: {
            x: Number(point1[0]),
            y: Number(point1[1]),
        },
        point2: {
            x: Number(point2[0]),
            y: Number(point2[1]),
        }
    }
};
