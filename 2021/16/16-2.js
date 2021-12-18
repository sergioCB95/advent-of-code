const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    let value = 0;

    let bits = data
        .split('')
        .filter(x => x !== '\n')
        .map(x => hexToBinarySwitch[x])
        .join('')
        .split('');

    ([value, bits] = processPacket(bits));

    console.log(value);
});

const processPacket = (_bits) => {
    let bits = [..._bits];
    let type;
    ([, type, bits] = getPacketHeader(bits));
    return isLiteralPacket(type.join(''))
        ? processLiteralPacket(bits)
        : processOperatorPacket(bits, type.join(''));
}

const processLiteralPacket = (bits) => {
    let literal;
    [literal, bits] = getLiteralValue(bits);
    return [getBitsValue(literal), bits];
}

const processOperatorPacket = (bits, type) => {
    let literalValue, lengthType, subPacketsLength, subPacketCount, literal;
    const literals = [];
    ([lengthType, bits] = getLengthType(bits));
    if (lengthType.join('') === '0') {
        ([subPacketsLength, bits] = getSubPacketLength(bits));
        let subBits = bits.splice(0, getBitsValue(subPacketsLength));
        while (subBits.length) {
            ([literal, subBits] = processPacket(subBits));
            literals.push(literal);
        }
    } else {
        ([subPacketCount, bits] = getSubPacketCount(bits));
        for (let i = 0; i < getBitsValue(subPacketCount); i++) {
            ([literal, bits] = processPacket(bits));
            literals.push(literal);
        }
    }
    literalValue = operatorTypeAction(type, literals);
    return [literalValue, bits];
}

const operatorTypeAction = (type, literals) => {
    const operatorSwitch = {
        '000': () => literals.reduce((acc, curr) => acc + curr, 0),
        '001': () => literals.reduce((acc, curr) => acc * curr, 1),
        '010': () => literals.reduce((acc, curr) => acc < curr ? acc : curr),
        '011': () => literals.reduce((acc, curr) => acc > curr ? acc : curr),
        '101': () => literals[0] > literals[1] ? 1 : 0,
        '110': () => literals[0] < literals[1] ? 1 : 0,
        '111': () => literals[0] === literals[1] ? 1 : 0,
    }
    return operatorSwitch[type]();
}

const getBitsValue = (bits) => parseInt(bits.join(''), 2);

const getSubPacketCount = (bits) => {
    let count = bits.splice(0, 11);
    return [count, bits];
};

const isLiteralPacket = (type) => type === '100';

const getSubPacketLength = (bits) => {
    let length = bits.splice(0, 15);
    return [length, bits];
}

const getLiteralValue = (bits) => {
    let [flag, ...literal] = bits.splice(0, 5);
     while (flag === '1') {
         [flag, ...newLiteral] = bits.splice(0, 5);
         literal = [...literal, ...newLiteral];
     }
     return [literal, bits];
}

const getPacketHeader = (bits) => {
    const version = bits.splice(0, 3);
    const type = bits.splice(0, 3);
    return [version, type, bits];
}

const getLengthType = (bits) => {
    const lengthType = bits.splice(0, 1);
    return [lengthType, bits];
}

const hexToBinarySwitch = {
    '0': '0000',
    '1': '0001',
    '2': '0010',
    '3': '0011',
    '4': '0100',
    '5': '0101',
    '6': '0110',
    '7': '0111',
    '8': '1000',
    '9': '1001',
    'A': '1010',
    'B': '1011',
    'C': '1100',
    'D': '1101',
    'E': '1110',
    'F': '1111',
}
