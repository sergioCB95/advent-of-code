const initTimes = () => {
    const times = Array(10).fill(0);
    for (let i = 1; i < 4; i++) {
        for (let j = 1; j < 4; j++) {
            for (let k = 1; k < 4; k++) {
                const value = i + j + k;
                times[value] += 1;
            }
        }
    }
    return times;
}

const movePlayer = (_positions, _times) => {
    const newPositions = Array(9).fill().map(() => ({
        count: 0,
        scores: Array(22).fill(0),
    }));
    for (let i = 0; i < _positions.length; i++) {
        _times.forEach((x, j) => {
            const nextPos = (i + j) % 9;
            newPositions[nextPos].count += (x * _positions[i].count);

            _positions[i].scores.forEach((score, scoreIndex) => {
                let nextScore = scoreIndex + nextPos + 1;
                nextScore = nextScore > 21 ? 21 : nextScore;
                newPositions[nextPos].scores[nextScore] += x * score;
            });
        });
    }
    return newPositions;
}


const times = initTimes();
let positions = Array(9).fill().map(() => ({
    count: 0,
    scores: Array(22).fill(0),
}));
positions[3].count = 1;
positions[3].scores[0] = 1;

for (let i = 0; i < 2; i++) {
    positions = movePlayer(positions, times);
}

console.log(positions);







