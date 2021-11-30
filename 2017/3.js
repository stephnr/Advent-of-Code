// ########################################################################################################################################################

var input = require('fs').readFileSync('./data/3').toString();

input;

var gridSize = 0;
var size = 0;

var i = 0;

// Determine grid size needed for puzzle input
do {
    if ((i * i) >= input) {
        gridSize = (i*i);
        size = i;
    } else i++;
} while(gridSize == 0);

var i = 1;
while((i * i) <= input) { i+=2; }

// Dimension size
var dimensionSize = i;
// Actual grid size
var totalSize = (dimensionSize * dimensionSize);
// Distance from corner
var distance = (i*i) - input;

var dimensionSegment = (dimensionSize-1);
var nearestCorner = 0;

// Determine which part of the grid ring the number is
for (var i = 1; i <= 4; i++) {
    if (input >= (totalSize - dimensionSegment*i)) nearestCorner = (totalSize - dimensionSegment*i);
}

var distanceToCorner = input - nearestCorner;
var distanceToMiddleOfSegment = Math.max(Math.round(dimensionSize/2), distanceToCorner) - Math.min(Math.round(dimensionSize/2), distanceToCorner);

// Part 1 Answer
// console.log(distanceToMiddleOfSegment + Math.round(dimensionSize/2));

// ########################################################################################################################################################

// Part 2

var gridSize = 9;

var grid = [
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
]

var x = Math.floor(gridSize/2);
var y = Math.floor(gridSize/2);

grid[x][y] = 1;

y++;

var i = 2;

do {
    var value = 0;

    var positionsAdded = new Array();

    // Directions (up, right, left, down)
    if (x+1 <= (gridSize-1)) { if (grid[x+1][y] > 0) { value += grid[x+1][y]; positionsAdded.push(1); } }
    if (y+1 <= (gridSize-1)) { if (grid[x][y+1] > 0) { value += grid[x][y+1]; positionsAdded.push(2); } }
    if (x-1 >= 0) { if (grid[x-1][y] > 0) { value += grid[x-1][y]; positionsAdded.push(3); } }
    if (y-1 >= 0) { if (grid[x][y-1] > 0) { value += grid[x][y-1]; positionsAdded.push(4); } }

    // Corners
    if (x-1 >= 0 && y-1 >= 0) { if (grid[x-1][y-1] > 0) { value += grid[x-1][y-1]; positionsAdded.push(5); } }
    if (x+1 <= (gridSize-1) && y-1 >= 0) { if (grid[x+1][y-1] > 0) { value += grid[x+1][y-1]; positionsAdded.push(6); } }
    if (x-1 >= 0 && y+1 <= (gridSize-1)) { if (grid[x-1][y+1] > 0) { value += grid[x-1][y+1]; positionsAdded.push(7); } }
    if (x+1 <= (gridSize-1) && y+1 <= (gridSize-1)) { if (grid[x+1][y+1] > 0) { value += grid[x+1][y+1]; positionsAdded.push(8); } }

    grid[x][y] = value;

    if (value >= input) value;

    // Switch the direction + coordinates
    var pa = positionsAdded.join('');

    if (pa == '4' || pa == '45' || pa == '1456' || pa == '146') x--;
    else if (pa == '16' || pa == '128' || pa == '16' || pa == '1268' || pa == '128') y--;
    else if (pa == '28' || pa == '237' || pa == '28' || pa == '2378' || pa == '237') x++;
    else if (pa == '37' || pa == '3457' || pa == '345' || pa == '37' || pa == '3457') y++;
    else console.log(pa);

    i++;
} while(i <= (gridSize*gridSize));

grid;

// ########################################################################################################################################################
