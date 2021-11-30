// ########################################################################################################################################################

var input = require('fs').readFileSync(`./data/${__filename.split('/').pop().substring(0,1)}`).toString();

var instructions = input.split('\n').map(x => parseInt(x));

var steps = 0;
var pos = 0;

try {

    do {
        var jump = instructions[pos];
        if (jump >= 3) instructions[pos]--; else instructions[pos]++;
        steps++;
        pos += jump;
    } while(pos >= 0 && pos < instructions.length);

} catch(e) {
    
} finally {
    console.log(steps);
}

// Part 1

// steps = 0;
// pos = 0;

// try {

//     do {
//         var jump = instructions[pos];
//         instructions[pos]++;        
//         steps++;
//         pos += jump;
//     } while(pos >= 0 && pos < instructions.length);

// } catch(e) {
    
// } finally {
//     steps;
//     pos;
// }
