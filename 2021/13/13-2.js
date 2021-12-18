const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    let maxY = null;
    let maxX = null;
    const lines = data.split('\n');
    let points = Object.fromEntries(lines
        .filter(x => /[0-9]+,[0-9]+/.test(x))
        .map(x => [x, x.split(',').map(y => Number(y))]));
    const actions = lines
        .filter(x => /^fold along/.test(x))
        .map(x => x.match(/fold along (?<axis>x|y)=(?<coord>[0-9]+)/).groups);

    actions.forEach(({ axis, coord}) => {
        if (axis === 'x') {
            maxX = Number(coord);
        } else {
            maxY = Number(coord);
        }
        points = fold(points, coord, axis === 'x' ? 0 : 1);
    })

    const grid = Array(maxY).fill(null)
        .map((_, y) => Array(maxX).fill(null).map((_, x) => points[`${x},${y}`] ? '#' : '.'));

    grid.forEach(x => console.log(x.join(' ')))
});

const fold = (points, coord, index) => {
    const entries = Object.entries(points);
    const removedPoints = entries.filter(([, coords]) => coords[index] > coord);
    const newPoints = Object.fromEntries(entries.filter(([, coords]) => coords[index] < coord));
    removedPoints.forEach(([key, coords]) => {
        const newCoords = coords;
        newCoords[index] = 2 * coord - newCoords[index];
        newPoints[newCoords.join(',')] = newCoords;
    });
    return newPoints;
}

