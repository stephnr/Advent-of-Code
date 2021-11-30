var input = require('fs').readFileSync('./data/2').toString();

// Part 1
var party = input.split('\n')
    .reduce((sum, row) => {
        row = row.split('\t').map(x => parseInt(x));
        sum += (Math.max(...row) - Math.min(...row));
        return sum;
    }, 0);

party;

// Part 2
var harty = input.split('\n')
    .reduce((sum, row) => {
        row = row.split('\t').map(x => parseInt(x));

        for(var i = 0; i <= row.length-1; i++) {
            for(var j = 0; j <= row.length-1; j++) {
                if (i == j) continue;
                else if ((row[i] % row[j]) === 0) {
                    sum += ((Math.max(row[i], row[j])) / (Math.min(row[i], row[j])));
                }
            }
        }

        return sum;
    }, 0);

harty;
