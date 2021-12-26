const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    let scanners = data.split('\n\n').map(scanner => {
        const lines = scanner.split('\n');
        const { name } = lines[0].match(/--- scanner (?<name>[0-9]+) ---/).groups;

        const beacons = lines
            .slice(1)
            .filter(x => x.length)
            .map(line => line.split(',')).map(line => ({
                x: Number(line[0]),
                y: Number(line[1]),
                z: Number(line[2])
            }));
        return {
            name,
            beacons,
        };
    });

    scanners = scanners.map(scanner => ({
        ...scanner,
        beacons: scanner.beacons.map(beacon => ({
            ...beacon,
            distances: scanner.beacons.map(beacon2 => calcDistance(beacon, beacon2)),
        }))
    }))
    const unifiedScanners = [scanners[0]];
    scanners[0].coords = {
        x: 0,
        y: 0,
        z: 0,
    };
    const diffBeacons = [...scanners[0].beacons];
    let leftScanners = scanners.slice(1);

    while (unifiedScanners.length) {
        const foundScanners = [];
        const targetScanner = unifiedScanners.pop();
        for (let i = 0; i < leftScanners.length; i++) {
            const [diffOnes, repeatedOnes] = compareScanners(targetScanner, leftScanners[i]);
            if (repeatedOnes.length >= 12) {
                const targetDiffs =  [
                    repeatedOnes[1][1].x - repeatedOnes[0][1].x,
                    repeatedOnes[1][1].y - repeatedOnes[0][1].y,
                    repeatedOnes[1][1].z - repeatedOnes[0][1].z
                ];
                const newDiffs =  [
                    repeatedOnes[1][0].x - repeatedOnes[0][0].x,
                    repeatedOnes[1][0].y - repeatedOnes[0][0].y,
                    repeatedOnes[1][0].z - repeatedOnes[0][0].z
                ];

                const xRel = findCoordRel(newDiffs, targetDiffs[0]);
                const yRel = findCoordRel(newDiffs, targetDiffs[1]);
                const zRel = findCoordRel(newDiffs, targetDiffs[2]);

                const scannerPosition = {
                    x: repeatedOnes[0][1].x - selectCoord(repeatedOnes[0][0], xRel),
                    y: repeatedOnes[0][1].y - selectCoord(repeatedOnes[0][0], yRel),
                    z: repeatedOnes[0][1].z - selectCoord(repeatedOnes[0][0], zRel),
                }

                leftScanners[i].coords = scannerPosition;

                leftScanners[i].beacons = leftScanners[i].beacons.map(beacon => {
                    const transformedBeacon = {
                        ...beacon,
                        ...normalizeBeacon(beacon, scannerPosition, [xRel, yRel, zRel]),
                    };
                    if (!diffBeacons.some(x => x.x === transformedBeacon.x && x.y === transformedBeacon.y && x.z === transformedBeacon.z)) {
                        diffBeacons.push(transformedBeacon);
                    }
                    return transformedBeacon;
                });

                unifiedScanners.push(leftScanners[i]);
                foundScanners.push(leftScanners[i]);
            }
        }
        foundScanners.forEach(i => leftScanners.splice(leftScanners.indexOf(i), 1));
    }
    const scannerCoords = scanners.map(x => ({ ...x.coords }));

    let max = 0;
    for (let i = 0; i < scannerCoords.length; i++) {
        const currentScanner = scannerCoords[i];
        for (let j = i + 1; j < scannerCoords.length; j++) {
            const iterativeScanner = scannerCoords[j];
            const manhattanDist = Math.abs(currentScanner.x - iterativeScanner.x) + Math.abs(currentScanner.y - iterativeScanner.y) + Math.abs(currentScanner.z - iterativeScanner.z);
            max = manhattanDist > max ? manhattanDist : max;
        }
    }
    console.log(scannerCoords);
    console.log(max);
});

const normalizeBeacon = (beacon, scannerPos, coordsRel) => ({
    x: selectCoord(beacon, coordsRel[0]) + scannerPos.x,
    y: selectCoord(beacon, coordsRel[1]) + scannerPos.y,
    z: selectCoord(beacon, coordsRel[2]) + scannerPos.z,
});

const selectCoord = (beacon, rel) => {
    const switcher = ['x', 'y', 'z']
    const result = beacon[switcher[rel.index]];
    return rel.sameDir ? result : -result;
}

const findCoordRel = (diffs, targetDiff) => {
    const index = diffs.findIndex(x => Math.abs(x) === Math.abs(targetDiff));
    return {
        index,
        sameDir: targetDiff === diffs[index],
    }
}

const compareScanners = (scanner, newScanner) => {
    const diffOnes = [];
    const repeatedOnes = [];
    for (let i = 0; i < newScanner.beacons.length; i++) {
        const currentBeacon = newScanner.beacons[i];
        const foundBeacon = scanner.beacons.find(x => compareDistances(x.distances, currentBeacon.distances));
        if (foundBeacon) {
            repeatedOnes.push([currentBeacon, foundBeacon]);
        } else {
            diffOnes.push(currentBeacon);
        }
    }
    return [diffOnes, repeatedOnes];
}

const calcDistance = (point1, point2) => Math.sqrt(
    Math.pow(point1.x - point2.x, 2)
    + Math.pow(point1.y - point2.y, 2)
    + Math.pow(point1.z - point2.z, 2)).toFixed(2);

const compareDistances = (arr1, arr2) => arr1.filter((value) => arr2.includes(value)).length >= 12
