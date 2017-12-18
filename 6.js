// ########################################################################################################################################################

var input = require('fs').readFileSync(`./data/${__filename.split('/').pop().substring(0,1)}`).toString();

var banks = input.split('\t').map(x => parseInt(x));

var snapshots = new Set();
var snapshotsLength = snapshots.size;

var cycles = 0;

while(!snapshots.has(banks.join(''))) {
    // console.log(`A: ${banks}`);

    snapshots.add(banks.join(''));
    // 1. Find largest bank
    var max = Math.max(...banks);
    var bankIdx = banks.indexOf(max);
    var idx = bankIdx;

    banks[idx] = 0;

    // 2. Reallocate looping to the right
    for(var i = 0; i <= max; i++) {
        if (idx+1 > banks.length) idx = 0;
        banks[idx]++;
        idx++;
    }

    // console.log(banks.reduce((sum, x) => sum + x, 0));

    banks[bankIdx] = 0;
    cycles++;

    // console.log(`B: ${banks}`);
}

console.log(cycles); // Part 1

var x = Array.from(snapshots);
console.log(cycles - x.findIndex(x => x == banks.join(''))); // Part 2
