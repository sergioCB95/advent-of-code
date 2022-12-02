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

const MyAction = {
    X: Action.Rock,
    Y: Action.Paper,
    Z: Action.Scissor
}

const ActionScore = {
    [Action.Rock]: 1,
    [Action.Paper]: 2,
    [Action.Scissor]: 3,
}

const Result = {
    Draw: 'draw',
    Player1: 'player1',
    Player2: 'player2',
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

getResultScore = (result) => {

}

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const games = data.split('\n').filter(line => line.length).map(line => line.split(' '));
    const result = games.reduce((acc, [opponent, me]) => {
        const winner = getPlayerWinner(OpponentAction[opponent], MyAction[me]);
        const score = ActionScore[MyAction[me]] + ResultScore[winner];
        return acc + score;
    },0);
    console.log(result);
});
