const fs = require('fs');

const getCharValue = (char) => {
    return char === char.toUpperCase()
        ? char.charCodeAt(0) - 'A'.charCodeAt(0) + 27
        : char.charCodeAt(0) - 'a'.charCodeAt(0) + 1;
}

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const lines = data.split('\n').filter(line => line.length);
    const result = lines.reduce((acc, curr) => {
        const half = Math.ceil(curr.length / 2);
        const firstHalf = curr.slice(0, half).split('');
        const secondHalf = curr.slice(half).split('');
        const itemsInBoth = firstHalf.filter(char => secondHalf.includes(char));
        return acc + getCharValue(itemsInBoth[0]);
    }, 0);
    console.log(result);
});
