const fs = require('fs');

const gamma = '000010111110';
const epsilon = '111101000001';

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const lines = data.split('\n').filter(x => x.length > 0);
    let o2 = [...lines];
    let co2 = [...lines];
    for(let i = 0; i < 12; i++) {
        if (o2.length > 1) {
            const sum = o2.map(line => Number(line[i])).reduce((acc, curr) => acc + curr, 0);
            const commonBit = sum >= o2.length / 2 ? 1 : 0;
            o2 = o2.filter(line => Number(line[i]) === commonBit);
        }
        if(co2.length > 1) {
            const sum = co2.map(line => Number(line[i])).reduce((acc, curr) => acc + curr, 0);
            const commonBit = sum >= co2.length / 2 ? 0 : 1;
            co2 = co2.filter(line => Number(line[i]) === commonBit);
        }
    }
    console.log(parseInt(o2[0], 2) * parseInt(co2[0], 2));
});
