const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    let max = 0;
    const result = data
        .split('\n')
        .filter(line => line.length)
        .map(line => JSON.parse(line))

    result.forEach((x, i) => {
        result.forEach((y, j) => {
            const xCopy = deepCopy(x);
            const yCopy = deepCopy(y);
            if (i !== j) {
                const res = sum(xCopy, yCopy);
                const magnitude = getMagnitude(res);
                console.log(JSON.stringify(res, null, 0));
                console.log(magnitude);
                max = Math.max(magnitude, max);
            }
        });
    });

    console.log(max);
});

const sum = (x, y) => {
    const result = [x, y];
    let explosiveNumber = findExplosiveNumber(result);
    let numberToSplit = false;
    if (!explosiveNumber) {
        numberToSplit = findSplitNumber(result);
    }

    while (explosiveNumber || numberToSplit) {
        if (explosiveNumber) {
            shareExplodedNumber(result, ...explosiveNumber);
        } else if (numberToSplit) {
            splitNumber(...numberToSplit);
        }
        explosiveNumber = findExplosiveNumber(result);
        if (!explosiveNumber) {
            numberToSplit = findSplitNumber(result);
        }
    }
    return result;
}

const findAndSum = (closeNum, addedNum, path, prev) => {
    if (closeNum[path] !== undefined) {
        if (!isNaN(closeNum[path])) {
            closeNum[path] += addedNum;
            return true;
        } else {
            return findAndSum(closeNum[path], addedNum, prev ? closeNum[path].length - 1 : 0, prev);
        }
    }
    return false;
}


const shareExplodedNumber = (num, explodedNum, path) => {
    if (path.length > 1) {
        const currentPath = [...path];
        let [prevFound, nextFound] = shareExplodedNumber(num[path.shift()], explodedNum, path);
        if (!prevFound) {
            prevFound = findAndSum(num, explodedNum[0], currentPath[0] - 1, true);
        }
        if (!nextFound) {
            nextFound = findAndSum(num, explodedNum[1], currentPath[0] + 1, false);
        }
        return [prevFound, nextFound];
    }
    num[path[0]] = 0;
    const prevFound = findAndSum(num, explodedNum[0], path[0] - 1, true);
    const nextFound = findAndSum(num, explodedNum[1], path[0] + 1, false);

    return [prevFound, nextFound];
}

const isSnailFishNumber = (num) => num?.length === 2 && !isNaN(num[0]) && !isNaN(num[1]);

const findExplosiveNumber = (num, path = [], depth = 0) => {
    let result = false;
    if (isSnailFishNumber(num)) {
        if (depth > 3) {
            return [num, path];
        } else {
            return false;
        }
    }
    for (let i = 0; i < num.length; i++) {
        const newPath = [...path];
        newPath.push(i);
        result = result || findExplosiveNumber(num[i], newPath, depth + 1);
    }
    return result;
}

const findSplitNumber = (num) => {
    let result = false;
    if (!num.length) {
        return false;
    }
    for (let i = 0; i < num.length; i++) {
        if (!isNaN(num[i]) && num[i] > 9) {
            result = [num, i];
            return result;
        }
        result = result || findSplitNumber(num[i]);
        if (result) {
            return result;
        }
    }
    return result;
}

const splitNumber = (num, index) => {
    const value = num[index];
    num[index] = [Math.floor(value / 2), Math.ceil(value / 2)]
}

const getMagnitude = (num) => {
    if (!isNaN(num)) {
        return num;
    }
    return (3 * getMagnitude(num[0])) + (2 * getMagnitude(num[1]));
}


const deepCopy = (num) => num.map(item => Array.isArray(item) ? deepCopy(item) : item);
