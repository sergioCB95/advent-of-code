const fs = require('fs');

const Action = {
    Rock: 'rock',
    Paper: 'paper',
    Scissor: 'scissor',
}

const OpponentAction = {
    A: Action.Rock,
    B: Action.Paper,
    C: Action.Scissor
}

const Result = {
    Draw: 'draw',
    Player1: 'player1',
    Player2: 'player2',
}

const MyResult = {
    X: Result.Player1,
    Y: Result.Draw,
    Z: Result.Player2
}

const ActionScore = {
    [Action.Rock]: 1,
    [Action.Paper]: 2,
    [Action.Scissor]: 3,
}

const ResultScore = {
    [Result.Draw]: 3,
    [Result.Player1]: 0,
    [Result.Player2]: 6,
}

const getPlayerWinner = (player1, player2) => {
    if (player1 === player2) {
        return Result.Draw;
    } else if (player1 === Action.Rock && player2 === Action.Scissor
        || player1 === Action.Paper && player2 === Action.Rock
        || player1 === Action.Scissor && player2 === Action.Paper) {
        return Result.Player1;
    }
    return Result.Player2;
}

const gessMyAction = (player1, result) => {
    if (result === Result.Draw) {
        return player1;
    }
    if (player1 === Action.Rock) {
        if (result === Result.Player1) {
            return Action.Scissor
        }
        return Action.Paper
    }
    if (player1 === Action.Paper) {
        if (result === Result.Player1) {
            return Action.Rock
        }
        return Action.Scissor
    }
    if (player1 === Action.Scissor) {
        if (result === Result.Player1) {
            return Action.Paper
        }
        return Action.Rock
    }
}

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const games = data.split('\n').filter(line => line.length).map(line => line.split(' '));
    const result = games.reduce((acc, [opponent, me]) => {
        const myAction = gessMyAction(OpponentAction[opponent], MyResult[me]);
        const score = ActionScore[myAction] + ResultScore[MyResult[me]];
        return acc + score;
    },0);
    console.log(result);
});
