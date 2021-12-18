const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const lines = data.split('\n').filter(x => x.length);
    let polymerTemplate = lines[0].split('');

    const rules = lines
        .filter(x => /[A-Z]+ -> [A-Z]+/.test(x))
        .map(x => x.match(/(?<rule>[A-Z]+) -> (?<result>[A-Z])+/).groups);

    for (let count = 0; count < 10; count++) {
        const result = [];
        for (let i = 0; i < polymerTemplate.length; i++) {
            const currentChar = polymerTemplate[i];
            result.push(currentChar);
            if ( i + 1 <= polymerTemplate.length - 1) {
                const resultedChar = rules
                    .filter((x) => x.rule === `${currentChar}${polymerTemplate[i+1]}`)[0].result;
                result.push(resultedChar);
            }
        }
        polymerTemplate = result;
    }

    const counts = {};
    polymerTemplate.forEach(x => {
        counts[x] = (counts[x] || 0) + 1;
    });

    console.log(counts);
});
