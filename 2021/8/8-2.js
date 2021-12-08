const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const lines = data.split('\n').filter(x => x.length);
    const parsedLines = lines.map(line => line.split('|').map(x => x.trim().split(' ')));
    const result = parsedLines
        .map(([input, output]) => [guessNumbers(input) ,output])
        .map(([input, output]) => {
            const nums = output.map(x => {
                const key = Object.keys(input)
                    .filter(y => [...y].sort().join('') === [...x].sort().join(''))[0];
                return input[key];
            })
            return Number(nums.join(''));
        }).reduce((acc, curr) => acc + curr, 0);
    console.log(result);
});

const guessNumbers = (numbers) => {
    const one = numbers.filter(x => x.length === 2)[0];
    const seven = numbers.filter(x => x.length === 3)[0];
    const four = numbers.filter(x => x.length === 4)[0];
    const eight = numbers.filter(x => x.length === 7)[0];

    const result = {
        [one]: '1',
        [seven]: '7',
        [four]: '4',
        [eight]: '8',
    };
    const fourSevenUnion = [...new Set([...four, ...seven])];
    const nine = numbers.filter(x => x.length === 6 && fourSevenUnion.every(y => x.includes(y)))[0];
    result[nine] = '9';

    const three = numbers.filter(x => x.length === 5 && [...one].every(y => x.includes(y)))[0];
    result[three] = '3';

    const zero = numbers.filter(x => x.length === 6 && x !== nine && [...one].every(y => x.includes(y)))[0];
    result[zero] = '0';

    const six = numbers.filter(x => x.length === 6 && x !== nine && x !== zero)[0];
    result[six] = '6';

    const diffEightNine = [...eight].filter(x => !nine.includes(x))[0];
    const two = numbers.filter(x => x.length === 5 && x !== three && x.includes(diffEightNine))[0];
    result[two] = '2';

    const five = numbers.filter(x => x.length === 5 && x !== three && x !== two)[0];
    result[five] = '5';

    return result;
}
