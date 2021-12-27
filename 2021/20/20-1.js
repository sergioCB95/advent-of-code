const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const [imageAlg, input] = data.split('\n\n');
    const imageAlgList = imageAlg.split('');
    let fillingChar = '.';
    let matrix = input.split('\n').filter(x => x.length).map(x => x.split(''));
    for(let i = 0; i < 50; i++) {
        [matrix, fillingChar] = enhanceImage(matrix, imageAlgList, fillingChar);
    }
    console.log(matrix.map(x => x.join('')).join('\n'));
    console.log(matrix.flatMap(x => x).filter(x => x === '#').length);
});

const coverWithFilling = (matrix, fillingChar) => {
    const newMatrix = matrix.map(row => {
        newRow = [...row];
        newRow.push(fillingChar);
        newRow.unshift(fillingChar);
        return newRow;
    });

    newMatrix.push(Array(newMatrix[0].length).fill(fillingChar));
    newMatrix.unshift(Array(newMatrix[0].length).fill(fillingChar));

    return newMatrix;
}

const enhanceElement = (surroundedElem, imageAlgList) => {
    const binaryNum = surroundedElem.map(x => x === '#' ? '1' : '0').join('');
    const num = parseInt(binaryNum, 2);
    return imageAlgList[num];
}

const enhanceImage = (_matrix, imageAlgList, fillingChar) => {
    let matrix = coverWithFilling(_matrix, fillingChar);
    matrix = matrix.map((row, i) => row.map((_, j) => {
        const getSafeMatrixElem = (_matrix, _i, _j) => (_matrix[_i] && _matrix[_i][_j]) || fillingChar
        const surroundedElem = [
            getSafeMatrixElem(matrix, i - 1, j - 1),
            getSafeMatrixElem(matrix, i - 1, j),
            getSafeMatrixElem(matrix, i - 1, j + 1),
            getSafeMatrixElem(matrix, i, j - 1),
            getSafeMatrixElem(matrix, i, j),
            getSafeMatrixElem(matrix, i, j + 1),
            getSafeMatrixElem(matrix, i + 1, j - 1),
            getSafeMatrixElem(matrix, i + 1, j),
            getSafeMatrixElem(matrix, i + 1, j + 1),
        ];

        return enhanceElement(surroundedElem, imageAlgList);
    }));
    return [
        matrix,
        enhanceElement(Array(9).fill(fillingChar), imageAlgList)
    ];
}
