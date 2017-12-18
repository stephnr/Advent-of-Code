// ########################################################################################################################################################

var input = require('fs').readFileSync(`./data/${__filename.split('/').pop().substring(0,1)}`).toString();

var passwords = input.split('\n');

var validPasswords = 0;
var remainingPasswords = [];

for(var idx in passwords) {
    var password = passwords[idx];
    var phrases = password.split(' ');
    if ((new Set(phrases).size == phrases.length)) {
        validPasswords++;
        remainingPasswords.push(password);
    };
}

validPasswords;

// Part 2

var validPasswords = 0;

for(var password of remainingPasswords) {
    var set = new Set();
    var phrases = password.split(' ');

    for(var phrase of phrases) {
        set.add(phrase.toLowerCase().split('').sort().join(''));
    }

    if (set.size == phrases.length) validPasswords++;
}

validPasswords;

// ########################################################################################################################################################
