const fs = require('fs');

fs.readFile('./data1.txt', 'utf8', (err, data) => {
    const lines = data.split('\n');
    const [h, d] = lines.reduce(([h, d], curr) => {
        const [action, num] = curr.split(' ');
        const parsedNum = Number(num);

        const actionMatch = {
            forward: () => [h + parsedNum, d],
            down: () => [h, d + parsedNum],
            up: () => [h, d - parsedNum],
            default: () => [h, d],
        };
        return actionMatch[action] ? actionMatch[action]() : actionMatch.default();
    }, [0,0]);
    console.log(h * d);
});
