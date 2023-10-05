const fs = require("fs");

function gtpSolution(elvesCalories) {
  // Split the input into lines and initialize variables
  const lines = elvesCalories.split("\n");
  let currentElfCalories = 0;
  let maxElfCalories = 0;

  // Iterate through the lines
  for (const line of lines) {
    const calories = parseInt(line);

    // Check if the line is empty or not a number (e.g., blank lines)
    if (!isNaN(calories)) {
      // Add the Calories to the current Elf's total
      currentElfCalories += calories;
    } else {
      // Update the maxElfCalories if the currentElfCalories are greater
      if (currentElfCalories > maxElfCalories) {
        maxElfCalories = currentElfCalories;
      }
      // Reset the currentElfCalories for the next Elf
      currentElfCalories = 0;
    }
  }

  // Check one more time in case the last Elf had the most Calories
  if (currentElfCalories > maxElfCalories) {
    maxElfCalories = currentElfCalories;
  }

  return maxElfCalories;
}

function mySolutions(input) {
  const res = input
    .split("\n\n")
    .map((el) => el.split("\n").reduce((acc, el) => acc + parseInt(el, 10), 0))
    .sort((A, B) => (A <= B ? 1 : -1))[0];

  return res;
}

console.time("MY  SOLUTION");
console.time("GTP SOLUTION");

const inputTxt = fs.readFileSync("./input.txt");

const maxCaloriesB = gtpSolution(inputTxt.toString());
const maxCaloriesA = mySolutions(inputTxt.toString());

console.log(
  `[MY SOLUTION IN JS] -The Elf carrying the most Calories has ${maxCaloriesA} Calories.`
);
console.timeEnd("MY  SOLUTION");

console.log(
  `[GTP SOLUTION IN JS]-The Elf carrying the most Calories has ${maxCaloriesB} Calories.`
);
console.timeEnd("GTP SOLUTION");
