// ########################################################################################################################################################

var input = require('fs').readFileSync(`./data/${__filename.split('/').pop().substring(0,1)}`).toString();

// Part 1

var instructions = new Map();
var matchOps = new Set();

var max = Number.MIN_SAFE_INTEGER;

input.split('\n').forEach(row => {
    var tokens = row.split(' ');
    var [key, op, val, cond, matchKey, matchOp, matchValue] = tokens;

    var currVal = instructions.get(key) || 0;
    var shouldOp = false;

    var checkInstructionVal = instructions.get(matchKey) || 0;

    switch (matchOp) {
        case '>=':
            shouldOp = (checkInstructionVal >= parseInt(matchValue));
            break;
        case '<=':
            shouldOp = (checkInstructionVal <= parseInt(matchValue));
            break;
        case '>':
            shouldOp = (checkInstructionVal > parseInt(matchValue));
            break;
        case '==':
            shouldOp = (checkInstructionVal == parseInt(matchValue));
            break;
        case '!=':
            shouldOp = (checkInstructionVal != parseInt(matchValue));
            break;
        case '<':
            shouldOp = (checkInstructionVal < parseInt(matchValue));
            break;
    }

    if (shouldOp) {
        switch(op) {
            case 'inc':
                currVal += parseInt(val);
                break;
            case 'dec':
                currVal -= parseInt(val);
                break;
        }
    }

    max = Math.max(max, currVal);

    instructions.set(key, currVal);
});

console.log(Math.max(...instructions.values()));

max;

// ########################################################################################################################################################
