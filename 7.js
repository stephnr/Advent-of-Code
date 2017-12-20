// ########################################################################################################################################################

var input = require('fs').readFileSync(`./data/${__filename.split('/').pop().substring(0,1)}`).toString();

var towers = input.split('\n');
var links = new Map();
var candidates = [];

for(var tower of towers) {
    var tokens = (new RegExp(/^([A-Za-z]+)\s\((\d+)\)\s?\-?\>?\s?(.*)/g)).exec(tower);
    // console.log(tokens);
    links.set(tokens[1], { weight: parseInt(tokens[2]), links: tokens[3].split(', ').length == 1 ? null : tokens[3].split(', ') });
    if (tokens[3].length > 0) candidates.push(tokens[1]);
}

// Candidates for bottom are towers with links

// Part 1 Answer = find candidate with no other candidate referencing it
var start = '';

candidates.forEach(name => {
    var found = true;
    candidates.forEach(towerName => {
        if(links.get(towerName).links.indexOf(name) >= 0) found = false;
    });
    if (found) start = name;
});

start;

// Part 2: Find the program with bad weights
var found = false;
var i = 0;

// Recursively calculates the weight of a tower
function calcWeights(weights) {
    if (weights == null) return 0;
    else if (Array.isArray(weights)) return weights.reduce((sum, weight) => sum + calcWeights(weight), 0);
    else if(typeof(weights) == 'string') {
        return links.get(weights).weight + calcWeights(links.get(weights).links);
    }
    else throw new Error('FAILURE');
}

var start = 'drjmjug';

while(!found && i < 1) {
    var weights = [];

    links.get(start).links.forEach(weight => {
        weight;
        weights.push(calcWeights(weight));
    });

    weights;

    i++;
}

/**
 * eugwuhl
 * smaygo
 * hmgrlpj
 * drjmjug <- needs to be 8 less
**/
