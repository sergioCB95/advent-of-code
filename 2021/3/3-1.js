const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const lines = data.split('\n');
    const acc = lines.reduce((acc, curr) => {
        const numbers = curr.split('').map(num => Number(num));
        return numbers.length === 12
            ? acc.map((num, i) => num + numbers[i])
            : acc;
    }, new Array(12).fill(0))

    const gamma = acc.map(num => num > Math.floor(lines.length / 2) ? 1 : 0).join('');
    const epsilon = acc.map(num => num > Math.floor(lines.length / 2) ? 0 : 1).join('');

    console.log(parseInt(gamma, 2) * parseInt(epsilon, 2));
});
