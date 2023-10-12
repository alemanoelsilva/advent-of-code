const fs = require("fs");

function countContainingPairs(input) {
  const pairs = input.trim().split("\n");
  let count = 0;

  for (const pair of pairs) {
    const [range1, range2] = pair
      .split(",")
      .map((r) => r.split("-").map(Number));

    // Sort the ranges to make comparison easier
    range1.sort((a, b) => a - b);
    range2.sort((a, b) => a - b);

    if (
      (range1[0] <= range2[0] && range1[1] >= range2[1]) || // range1 fully contains range2
      (range2[0] <= range1[0] && range2[1] >= range1[1]) // range2 fully contains range1
    ) {
      count++;
    }
  }

  return count;
}

const inputTxt = fs.readFileSync("../input.txt");
const result = countContainingPairs(inputTxt.toString());
console.log(
  `The number of assignment pairs where one range fully contains the other is: ${result}`
);
