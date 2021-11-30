var input = require('fs').readFileSync('./data/1').toString();

var matches = input.match(/(\d)\1+/g);

// Part 1
var sum = matches.reduce((sum, x) => sum + (parseInt(x.charAt(0))*(x.length-1)), (input[0] == input[input.length-1]) ? parseInt(input[0]) : 0);

sum;

// Part 2
var digits = input.split('');
var sum = 0;

for(var i = 0; i <= digits.length-1; i++)
{
    var left = parseInt(digits[i]);
    var right = parseInt(digits[(i+(digits.length/2))%digits.length]);
    if (left == right) sum += left;
}

sum;
