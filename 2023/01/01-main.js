var fs = require('fs'),
    path = require('path'),
    filePath = path.join(__dirname, 'inputs', '1.txt');

function partOne() {
  fs.readFile(filePath, {encoding: 'utf-8'}, function(err,data){
    if (!err) {
        lines = data.split('\n')
        sum = 0

        for (let i = 0; i < lines.length; i++) {
          console.log(lines[i]);
          line = lines[i].replace(/\D/g,'');
          console.log(line)

          if (line.length == 1) {
            console.log(line)
            sum += parseInt(line + line, 10)
          } else if (line.length == 2) {
            console.log(line)
            sum += parseInt(line, 10)
          } else {
            console.log(line[0] + line[line.length-1])
            sum += parseInt(line[0] + line[line.length-1], 10)
          }

          console.log()
        }

        console.log(sum)
    } else {
        console.log(err);
    }
  });
}


partOne()
