const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const grid = data
        .split('\n')
        .filter(x => x.length)
        .map(x => x.split('').map(z => Number(z)));

    const result = grid.map((x, i) =>
            x.map((y, j) =>
                isLowPoint(grid, i, j)
                    ? calcBasin(grid.map(x => [...x]), i, j)
                    : 0));

    console.log(result.flatMap(x => x).sort((a, b) => b - a).slice(0, 3).reduce((acc, curr) => acc * curr, 1));
});

const calcBasin = (grid, i, j) => {
    if (i < 0 || i > (grid.length - 1)) return 0;
    if (j < 0 || j > (grid[0].length - 1)) return 0;
    if (grid[i][j] > 8) return 0;
    grid[i][j] = 9;

    return 1 + calcBasin(grid, i + 1, j)
        + calcBasin(grid, i - 1, j)
        + calcBasin(grid, i, j - 1)
        + calcBasin(grid, i, j + 1);
}

const isLowPoint = (grid, i, j) => {
    const upIsLower = i < 1 || grid[i][j] < grid[i - 1][j];
    const downIsLower = i > (grid.length - 2) || grid[i][j] < grid[i + 1][j];
    const leftIsLower = j < 1 || grid[i][j] < grid[i][j - 1];
    const rightIsLower = j > (grid[0].length - 2) || grid[i][j] < grid[i][j + 1];

    return upIsLower && downIsLower && leftIsLower && rightIsLower;
};
