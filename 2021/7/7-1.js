const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    //bruteForce(data);
    usingMedian(data);
});

const usingMedian = (data) => {
    let result = null;
    const numbers = data.split(',').map(x => Number(x)).sort((a,b) => a - b );
    if (numbers.length % 2 === 0 ) {
        result = (numbers[(numbers.length / 2) - 1] + numbers[numbers.length / 2]) / 2;
    } else {
        result = numbers[numbers.length / 2];
    }
    console.log(result)
    console.log(numbers.map(x => Math.abs(x - result)).reduce((acc, curr) => acc + curr, 0));
}

const bruteForce = (data) => {
    let index = 0;
    let value = null;
    let lastValue = null;
    const numbers = data.split(',').map(x => Number(x));

    while (lastValue === null || value <= lastValue) {
        lastValue = value;
        value = numbers.map(x => Math.abs(x - index)).reduce((acc, curr) => acc + curr, 0);
        index++;
    }

    console.log('Final');
    console.log(lastValue);
    console.log(value);
    console.log(index);
}
