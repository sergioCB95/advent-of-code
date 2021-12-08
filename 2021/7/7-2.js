const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    let index = 0;
    let value = null;
    let lastValue = null;
    const numbers = data.split(',').map(x => Number(x));

    while (lastValue === null || value <= lastValue) {
        lastValue = value;
        value = numbers.map(x => {
            const n = Math.abs(x - index);
            return (Math.pow(n, 2) + n) / 2
        }).reduce((acc, curr) => acc + curr, 0);
        index++;
    }

    console.log('Final');
    console.log(lastValue);
    console.log(value);
    console.log(index);
});
