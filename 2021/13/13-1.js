const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const lines = data.split('\n');
    let points = Object.fromEntries(lines
        .filter(x => /[0-9]+,[0-9]+/.test(x))
        .map(x => [x, x.split(',').map(y => Number(y))]));
    const actions = lines
        .filter(x => /^fold along/.test(x))
        .map(x => x.match(/fold along (?<axis>x|y)=(?<coord>[0-9]+)/).groups);

    const index = actions[0].axis === 'x' ? 0 : 1;
    points = fold(points, actions[0].coord, index);
    console.log(Object.entries(points).length);
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

