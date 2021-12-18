const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    let version, type, lengthType, subPacketsLength, subPacketCount;
    let versionCount = 0;

    let bits = data
        .split('')
        .filter(x => x !== '\n')
        .map(x => hexToBinarySwitch[x])
        .join('')
        .split('');

    while (bits.length > 4) {
        ([version, type, bits] = getPacketHeader(bits));
        if (isLiteralPacket(type.join(''))) {
            bits = removeLiteralValue(bits);
        } else {
            ([lengthType, bits] = getLengthType(bits));
            if (lengthType.join('') === '0') {
                ([subPacketsLength, bits] = getSubPacketLength(bits));
            } else {
                ([subPacketCount, bits] = getSubPacketCount(bits));
            }
        }
        versionCount += parseInt(version.join(''), 2);
    }
    console.log(versionCount);
});

const getSubPacketCount = (bits) => {
    let count = bits.splice(0, 11);
    return [count, bits];
};

const isLiteralPacket = (type) => type === '100';

const getSubPacketLength = (bits) => {
    let length = bits.splice(0, 15);
    return [length, bits];
}

const removeLiteralValue = (bits) => {
    let literal = bits.splice(0, 5);
     while (literal[0] === '1') {
         literal = bits.splice(0, 5);
     }
     return bits;
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
