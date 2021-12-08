const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const lines = data.split('\n').filter(x => x.length);
    const output = lines.flatMap(line => line.split('|')[1].trim().split(' '));
    const result = output
        .map(x => x.trim())
        .filter(x => x.length === 2 || x.length === 3 || x.length === 4 || x.length === 7);
    console.log(result.length);
});
