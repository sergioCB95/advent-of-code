const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const lines = data.split('\n').filter(x => x.length);
    const result = lines.reduce((acc, line) => {
        const openedClosures = [];
        const chars = line.split('');
        for (let char of chars) {
            if (isOpenClosure(char)) {
                openedClosures.push(char);
            } else {
                if (!matchOpenClosure(openedClosures.pop(), char)) {
                    return acc + getValue(char);
                }
            }
        }
        return acc;
    }, 0);

    console.log(result);
});

const isOpenClosure = (c) => ['(', '{', '[', '<'].includes(c);
const matchOpenClosure = (openC, closeC) => (openC === '(' && closeC === ')')
    || (openC === '{' && closeC === '}')
    || (openC === '[' && closeC === ']')
    || (openC === '<' && closeC === '>')

const getValue = (c) => {
    const match = {
        ')': 3,
        ']': 57,
        '}': 1197,
        '>': 25137,
    }
    return match[c];
}
