const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const result = data
        .split('\n')
        .filter(line => line.length)
        .map(line => JSON.parse(line))
        .reduce((acc, curr) => {
            const result = [acc, curr];
            let explosiveNumber = findExplosiveNumber(result);
            while (explosiveNumber) {
                console.log(JSON.stringify(result, null, 0));
                shareExplodedNumber(result, ...explosiveNumber);
                explosiveNumber = findExplosiveNumber(result);
            }
            return result;
        });
    console.log(JSON.stringify(result, null, 0));

});

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
