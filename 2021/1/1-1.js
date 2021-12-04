const fs = require('fs');

fs.readFile('./data1.txt', 'utf8', (err, data) => {
    const lines = data.split('\n').map(line => Number(line));
    const [, result] = lines.reduce(([last, acc], curr) =>
            (last !== null && curr > last) ? [curr, acc + 1] : [curr, acc],
        [null, 0]
    )
    console.log(result);
});
