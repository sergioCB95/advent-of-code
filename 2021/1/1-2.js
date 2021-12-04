const fs = require('fs');

fs.readFile('./data1.txt', 'utf8', (err, data) => {
    const lines = data.split('\n').map(line => Number(line));
    const [, result] = lines.reduce(([last, acc], curr) =>
            hasIncreased(curr, last) ? [updateLast(curr, last), acc + 1] : [updateLast(curr, last), acc],
        [{
            minus1: null,
            minus2: null,
            minus3: null,
        }, 0]
    )
    console.log(result);
});

const hasIncreased = (curr, last) => {
    if (last.minus1 === null || last.minus2 === null || last.minus3 === null) {
        return false
    }
    return (curr + last.minus1 + last.minus2) > (last.minus1 + last.minus2 + last.minus3)
}

const updateLast = (curr, last) => ({
    minus1: curr,
    minus2: last.minus1,
    minus3: last.minus2,
})
