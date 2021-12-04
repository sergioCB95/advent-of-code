const fs = require('fs');

fs.readFile('./data1.txt', 'utf8', (err, data) => {
    const lines = data.split('\n');
    const [, h, d] = lines.reduce(([a, h, d], curr) => {
        const [action, num] = curr.split(' ');
        const parsedNum = Number(num);

        const actionMatch = {
            forward: () => [a, h + parsedNum, d + (a * parsedNum)],
            down: () => [a + parsedNum, h, d],
            up: () => [a - parsedNum, h, d],
            default: () => [a, h, d],
        };
        return actionMatch[action] ? actionMatch[action]() : actionMatch.default();
    }, [0, 0,0]);
    console.log(h * d);
});
