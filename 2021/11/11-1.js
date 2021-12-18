const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    let grid = data.split('\n')
        .filter(x => x.length).map(x => x.split('').map(y => Number(y)));
    let count = 0;

    for (let i = 0; i < 100; i++) {
        grid = grid.map(x => x.map(y => y + 1));

        while (grid.flat().filter(x => x > 9).length) {
            for (let x = 0; x < grid.length; x++) {
                for (let y = 0; y < grid[x].length; y++) {
                    if (grid[x][y] > 9) {
                        count ++;
                        grid[x][y] = 0;
                        flash(grid, x, y);
                    }
                }
            }
        }
    }

    console.log(grid);
    console.log(count);
});


const flash = (grid, x, y) => {
    const indexes = [-1, 0, 1];
    const isFlashingElement = (i, j) => (i === 0 && j === 0);
    const isExhausted = (i, j) => (grid[x + i][y + j] === 0);
    const isNotOutOfRange = (i, j) => grid[x + i] !== undefined && grid[x + i][y + j] !== undefined;
    indexes.forEach(i => {
        indexes.forEach(j => {
            if (isNotOutOfRange(i, j) && !isFlashingElement(i, j) && !isExhausted(i, j)) {
                grid[x + i][y + j] += 1;
            }
        })
    })
}
