const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const seedGrid = data
        .split('\n')
        .filter(x => x.length)
        .map(x => x.split('')
            .map(y => ({
                weight: Number(y),
                dist: Infinity,
                prev: null,
            })));

    const yLength = seedGrid.length;
    const xLength = seedGrid[0].length;

    const grid = Array(yLength * 5).fill()
        .map((line, i)=> Array(xLength * 5).fill().map((_, j) => {
            const weight = (seedGrid[i % yLength][j % xLength].weight + Math.floor(j / xLength) + Math.floor(i / yLength));
            return {
                weight: weight < 10 ? weight : (((Math.floor(weight / 10) + weight) % 10)),
                dist: Infinity,
                prev: null,
            }
        }))

    let verts = [];
    for (let i = 0; i < grid.length; i++) {
        for (let j = 0; j < grid[i].length; j++) {
            grid[i][j].adj = [
                grid[i][j + 1],
                grid[i][j - 1],
                grid[i + 1] && grid[i + 1][j],
                grid[i - 1] && grid[i - 1][j],
            ];
        }
        verts = [...verts, ...grid[i]];
    }

    let currentVert = verts.shift();
    currentVert.dist = 0;

    const end = verts[verts.length - 1];

    while (currentVert !== end) {
        for (let adj of currentVert.adj) {
            if (adj) {
                const alt = currentVert.dist + adj.weight;
                if (alt < adj.dist) {
                    adj.dist = alt;
                    adj.prev = currentVert;
                }
            }
        }
        currentVert = verts.reduce((a, b) => a.dist < b.dist ? a : b);
        const index = verts.indexOf(currentVert);
        verts.splice(index, 1);
    }

    console.log(currentVert);
});
