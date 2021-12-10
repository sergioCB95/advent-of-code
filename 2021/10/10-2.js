const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const lines = data.split('\n').filter(x => x.length);
    const result = lines.map(( line) => {
        const openedClosures = [];
        const chars = line.split('');
        for (let char of chars) {
            if (isOpenClosure(char)) {
                openedClosures.push(char);
            } else {
                if (!matchOpenClosure(openedClosures.pop(), char)) {
                    return [];
                }
            }
        }
        return openedClosures
            .reduceRight((acc, curr) => (acc * 5) + getValue(curr), 0);
    })
        .filter(x => x.length === undefined)
        .sort((b, a) => b - a)
    
    console.log(result[Math.floor(result.length / 2)]);
});

const isOpenClosure = (c) => ['(', '{', '[', '<'].includes(c);
const matchOpenClosure = (openC, closeC) => (openC === '(' && closeC === ')')
    || (openC === '{' && closeC === '}')
    || (openC === '[' && closeC === ']')
    || (openC === '<' && closeC === '>')

const getValue = (c) => {
    const match = {
        '(': 1,
        '[': 2,
        '{': 3,
        '<': 4,
    }
    return match[c];
}
