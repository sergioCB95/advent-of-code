const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    let fishState = data
        .split(',')
        .map((x => Number(x)))
        .reduce((acc, curr) => {
            acc[curr] += 1;
            return acc;
        }, Array(9).fill(0));

    for (let i = 0; i < 256; i++) {
        const fishStateCopy = [...fishState];
        const bornFish = fishStateCopy[0];
        fishState = fishState.map((state, i) => {
            if (i < fishState.length - 1 && i !== 6) {
                return fishStateCopy[i + 1];
            } else if (i === 6) {
                return fishStateCopy[i + 1] + bornFish;
            }
            return bornFish;
        });
    }

    console.log(fishState.reduce((acc, curr) => acc + curr , 0));
});
