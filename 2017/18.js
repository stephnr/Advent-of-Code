// ########################################################################################################################################################

var input = require('fs').readFileSync(`./data/${__filename.split('/').pop().substring(0,2)}`).toString();

// Series of registers
// Each start with 0

// snd X => plays freq of X
// set X Y => X = Y
// add X Y => X += Y
// mul X Y => X *= Y
// mod X Y => X % Y
// rcv X => X = last played freq if register X != 0
// jgz X Y => jumps with offset Y if X > 0 (can go backwards)

var instr = input.split('\n');

var registers = new Map();

var done = false;
var cycles = 0;
var curr = 0;
var freq = 0;

var times = 0;

while(cycles < 1374 && !done) {
    if (curr > instr.length || curr < 0) {
        cycles = Number.MAX_SAFE_INTEGER;
        return;
    }

    var [ op, a, b, c ] = instr[curr].split(' ');

    switch(op) {
        case 'add':
            registers.set(a, ((registers.get(a) || 0) + parseInt(b)));
            curr++;
            break;
        case 'set':
            if (isNaN(parseInt(b))) registers.set(a, (registers.get(b) || 0));
            else registers.set(a, parseInt(b));
            curr++;
            break;
        case 'mul':
            registers.set(a, ((registers.get(a) || 0) * parseInt(b)));
            curr++;
            break;
        case 'jgz':
            if ((registers.get(a) || 0) > 0) {
                if (isNaN(parseInt(b))) curr += (registers.get(a) || 0);
                else curr += parseInt(b);
            } else curr++;
            break;
        case 'mod':
            if (isNaN(parseInt(b))) registers.set(a, (registers.get(a) || 0) % (registers.get(b) || 0));
            else registers.set(a, (registers.get(a) || 0) % parseInt(b));
            curr++;
            break;
        case 'snd':
            freq = (registers.get(a) || 0);
            curr++;
            break;
        case 'rcv':
            freq; // look here for answer to Part 1
            if ((registers.get(a) || 0) != 0) registers.set(a, freq);
            curr++;
            break;
    }

    cycles++;
}

registers;

// Part 2

var queueA = [];
var queueB = [];

var registersA = new Map();
var registersB = new Map();

var currA = 0;
var currB = 0;

var useA = true;

var cycles = 0;

var doneA = false;
var doneB = false;

function handleOp(op, a, b) {
    if (useA) {
        switch(op) {
            case 'add':
                if (isNaN(parseInt(b))) registersA.set(a, ((registersA.get(a) || 0) + (registersA.get(b) || 0)));
                else registersA.set(a, ((registersA.get(a) || 0) + parseInt(b)));
                currA++;
                break;
            case 'set':
                if (isNaN(parseInt(b))) registersA.set(a, (registersA.get(b) || 0));
                else registersA.set(a, parseInt(b));
                currA++;
                break;
            case 'mul':
                registersA.set(a, ((registersA.get(a) || 0) * parseInt(b)));
                currA++;
                break;
            case 'jgz':
                if ((registersA.get(a) || 0) > 0) {
                    if (isNaN(parseInt(b))) currA += (registersA.get(a) || 0);
                    else currA += parseInt(b);
                } else currA++;
                break;
            case 'mod':
                if (isNaN(parseInt(b))) registersA.set(a, (registersA.get(a) || 0) % (registersA.get(b) || 0));
                else registersA.set(a, (registersA.get(a) || 0) % parseInt(b));
                currA++;
                break;
            case 'snd':
                if (isNaN(parseInt(b))) queueB.push(registersA.get(b));
                else queueB.push(parseInt(b));
                currA++;
                break;
            case 'rcv':
                var val = queueB.shift();
                val;
                if (typeof val == 'undefined') useA = false;
                else {
                    registersA.set(a, val);
                }
                break;
        }
    } else {
        switch(op) {
            case 'add':
                registersB.set(a, ((registersB.get(a) || 0) + parseInt(b)));
                currB++;
                break;
            case 'set':
                if (isNaN(parseInt(b))) registersB.set(a, (registersB.get(b) || 0));
                else registersB.set(a, parseInt(b));
                currB++;
                break;
            case 'mul':
                registersB.set(a, ((registersB.get(a) || 0) * parseInt(b)));
                currB++;
                break;
            case 'jgz':
                if ((registersB.get(a) || 0) > 0) {
                    if (isNaN(parseInt(b))) currB += (registersB.get(a) || 0);
                    else currB += parseInt(b);
                } else currB++;
                break;
            case 'mod':
                if (isNaN(parseInt(b))) registersB.set(a, (registersB.get(a) || 0) % (registersB.get(b) || 0));
                else registersB.set(a, (registersB.get(a) || 0) % parseInt(b));
                currB++;
                break;
            case 'snd':
                times++;
                if (isNaN(parseInt(b))) queueA.push(registersB.get(b));
                else queueA.push(parseInt(b));
                currB++;
                break;
            case 'rcv':
                var val = queueA.shift();
                if (val == undefined) useA = true;
                else {
                    registersB.set(a, val);
                }
                break;
        }
    }
}

while (cycles < 5000 && !doneA && !doneB) {
    var curr = useA ? currA : currB;
    var [ op, a, b ] = instr[curr].split(' ');

    handleOp(op, a, b);

    cycles++;
}

console.log(`A:`, registersA);
console.log(`B:`, registersB);
currA;
currB
