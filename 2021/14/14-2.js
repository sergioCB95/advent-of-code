const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const lines = data.split('\n').filter(x => x.length);
    let polymerTemplate = lines[0];
    let counts = {};
    const firstChar = polymerTemplate.slice(0, 1);
    const lastChar = polymerTemplate.slice(polymerTemplate.length - 1);

    const rules = Object.fromEntries(lines
        .filter(x => /[A-Z]+ -> [A-Z]+/.test(x))
        .map(x => x.match(/(?<rule>[A-Z]+) -> (?<result>[A-Z])+/).groups)
        .map((x => {
            const chars = x.rule.split('');
            const results = [`${chars[0]}${x.result}`, `${x.result}${chars[1]}`];

            return [x.rule, results]
        })));

    for (let i = 0; i < polymerTemplate.length - 1; i++) {
        const slice = polymerTemplate.slice(i, i + 2);
        counts[slice] = (counts[slice] || 0) + 1;
    }

    for (let count = 0; count < 40; count++) {
        const newCount = {};
        Object.entries(counts).forEach(([key, value]) => {
            const results = rules[key];
            newCount[results[0]] = (newCount[results[0]] || 0) + value;
            newCount[results[1]] = (newCount[results[1]] || 0) + value;
        })
        counts = newCount;
    }

    let finalCount = {};

    Object.entries(counts).forEach(([key, value]) => {
        const chars = key.split('');
        finalCount[chars[0]] = (finalCount[chars[0]] || 0) + value;
        finalCount[chars[1]] = (finalCount[chars[1]] || 0) + value;
    });

    finalCount[firstChar] += 1;
    finalCount[lastChar] += 1;

    finalCount = Object.fromEntries(Object.entries(finalCount).map(([key, value]) => [key, value / 2]));

    console.log(finalCount);
});
