const fs = require('fs');

const getCharValue = (char) => {
    return char === char.toUpperCase()
        ? char.charCodeAt(0) - 'A'.charCodeAt(0) + 27
        : char.charCodeAt(0) - 'a'.charCodeAt(0) + 1;
}

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const groups = data
        .split('\n')
        .filter(line => line.length)
        .reduce((acc, curr) => {
            if (acc[0].length < 3) {
                acc[0].push(curr);
                return acc;
            } else {
                return [[curr],...acc]
            }
        }, [[]]);
    const result = groups.reduce((acc, curr) => {
        const first = curr[0].split('');
        const second = curr[1].split('');
        const third = curr[2].split('');
        const commonChars = first
            .filter(char => second.includes(char))
            .filter(char => third.includes(char))
        return acc + getCharValue(commonChars[0]);
    }, 0);
    console.log(result);
});
