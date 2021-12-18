const fs = require('fs');

fs.readFile('./data.txt', 'utf8', (err, data) => {
    const lines = data.split('\n').filter(x => x.length);
    const graph = initGraph(lines);
    console.log(trackPaths('start', graph));
});

const initGraph = (lines) => {
    const graph = {};
    lines.forEach(line => {
        const conn = line.split('-');
        if (graph[conn[0]] === undefined) {
            addNode(graph, conn[0]);
        }
        if (graph[conn[1]] === undefined) {
            addNode(graph, conn[1]);
        }
        graph[conn[0]].conns.push(conn[1]);
        graph[conn[1]].conns.push(conn[0]);
    });

    return graph;
}

const trackPaths = (node, graph, path = [], smallCaveVisited = false) => {
    const currentPath = [...path];
    // console.log('Node: ', node, ' -- current path: ', currentPath);
    currentPath.push(node);
    if (node === 'end') {
        // console.log('Path found: ', currentPath);
        return 1;
    }
    return graph[node].conns
        .filter(x => !(graph[x].isSmall && currentPath.includes(x)))
        .reduce((acc, curr) => acc + trackPaths(curr, graph, currentPath) , 0)
}


const addNode = (graph, nodeName) => {
    graph[nodeName] = { isSmall: !/[A-Z]+/.test(nodeName), conns: [] };
}
