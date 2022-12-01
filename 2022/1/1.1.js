const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const calories = data
        .split('\n\n')
        .map(elf => elf.split('\n').reduce((acc, curr) => acc + Number(curr), 0));
    console.log(Math.max(...calories));
});
