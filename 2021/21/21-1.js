const dice = {
    value: 1,
    times: 0,
}

const player1 = {
    pos: 7,
    score: 0,
};

const player2 = {
    pos: 4,
    score: 0,
};

const players = [
    player1,
    player2,
]


const rollDice = (_dice) => {
    const value = _dice.value;
    _dice.value += 1;
    if (_dice.value > 100) {
        _dice.value = 1;
    }
    _dice.times += 1;
    return value;
}

let winner = null;

while (!players.filter(player => player.score >= 1000).length) {
    for (let player of players) {
        const value = rollDice(dice) + rollDice(dice) + rollDice(dice);
        player.pos += value;
        if (player.pos > 9) {
            player.pos = player.pos % 10;
        }
        player.score += player.pos + 1;

        if (player.score >= 1000) {
            winner = player;
            break;
        }
    }
}

const loser = players.find(player => player !== winner);

console.log(loser.score);
console.log(dice.times);
console.log(loser.score * dice.times);

