const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const grid = data
        .split('\n')
        .filter(x => x.length)
        .map(x => x.split('').map(z => Number(z)));

    const result = grid.flatMap((x, i) =>
            x.filter((y, j) => isLowPoint(grid, y, i, j)))
            .reduce((acc, curr) => acc + curr + 1, 0);

    console.log(result);
});

const isLowPoint = (grid, item, i, j) => {
    const upIsLower = i < 1 || item < grid[i - 1][j];
    const downIsLower = i > (grid.length - 2) || item < grid[i + 1][j];
    const leftIsLower = j < 1 || item < grid[i][j - 1];
    const rightIsLower = j > (grid[0].length - 2) || item < grid[i][j + 1];

    return upIsLower && downIsLower && leftIsLower && rightIsLower;
};
