const WIN_POINTS = 21;

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

const movePlayer = (_positions, _times, _winners) => {
    const newWinners = [..._winners];
    newWinners.push(0);
    const newPositions = Array(9).fill().map(() => ({
        count: 0,
        scores: Array(22).fill(0),
    }));
    for (let i = 0; i < _positions.length; i++) {
        _times.forEach((x, j) => {
            const nextPos = (i + j) % 9;
            newPositions[nextPos].count += (x * _positions[i].count);

            _positions[i].scores.forEach((score, scoreIndex) => {
                if (scoreIndex >= WIN_POINTS) {
                    return;
                }
                let nextScore = scoreIndex + nextPos + 1;
                nextScore = nextScore > WIN_POINTS ? WIN_POINTS : nextScore;
                const scoreMove = x * score;
                newPositions[nextPos].scores[nextScore] += scoreMove;

                if (nextScore === WIN_POINTS) {
                    newWinners[newWinners.length - 1] += scoreMove;
                }
            });
        });
    }
    return [newPositions, newWinners];
}

const IsPlaying = (_positions) => !!_positions
    .flatMap(position => position.scores.filter((x, i) => i !== WIN_POINTS && x)).length

const calcPlayer = (startPosition, _times) => {
    let winners = [];
    let positions = Array(9).fill().map(() => ({
        count: 0,
        scores: Array(22).fill(0),
    }));
    positions[startPosition].count = 1;
    positions[startPosition].scores[0] = 1;

    while (IsPlaying(positions)) {
        [positions, winners] = movePlayer(positions, _times, winners);
    }

    const sum = positions.reduce((acc, curr) => acc + curr.count , 0);
    console.log(sum);
    console.log(winners);
    return winners;
}


const times = initTimes();
const possibilities = times.reduce((acc, curr) => acc + curr, 0);
const winner1 = calcPlayer(3, times);
const winner2 = calcPlayer(7, times);

const length = winner1.length > winner2.length ? winner1.length : winner2.length;

let wins1 = 0;
let wins2 = 0;
let permutations1 = possibilities;
let permutations2 = possibilities;
for (let i = 0; i < length; i++) {
    permutations1 = permutations1 - winner1[i];
    wins1 += winner1[i] * permutations2;
    wins2 += winner2[i] ? winner2[i] * permutations1 : 0;
    permutations1 = (permutations1 * possibilities);
    permutations2 = (permutations2 * possibilities) - winner2[i];
}

console.log(wins1);
console.log(wins2);






