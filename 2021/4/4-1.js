const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const lines = data.split('\n\n');
    let finish = false;

    const bingoNumbers = lines[0].split(',');
    let boards = lines
        .slice(1)
        .map(board => board
            .split('\n')
            .filter(line => line.length)
            .map(line => line
                .split(/ {1,2}/)
                .map(num => ({ num, marked: false }))));

    for (let bingo of bingoNumbers) {
        boards = boards.map(board => {
            const newBoard = markNumber(bingo, board);
            if (matchBingo(newBoard)) {
                const sum = newBoard
                    .flat()
                    .filter(number => !number.marked)
                    .reduce((acc, curr) => acc + Number(curr.num), 0);
                console.log(sum * Number(bingo));
                finish = true;
            }
            return newBoard;
        })
        if (finish) {
            break;
        }
    }
});

const markNumber = (markedNumber, board) => board.map(line =>
    line.map(number => number.num === markedNumber
        ? {...number, marked: true }
        : number))

const matchBingo = (board) => {
    for (let line of board) {
        if (line.filter(number => number.marked).length === 5) {
            return true;
        }
    }
    for (let col = 0; col < 5; col++) {
        if (board.map(line => line[col]).filter(number => number.marked === true).length === 5) {
            return true;
        }
    }
    return false;
}
