const fs = require("fs");


const ELF_PLAY = {
  ROCK: "A",
  PAPER: "B",
  SCISSOR: "C",
};

const MY_PLAY = {
  ROCK: "X",
  PAPER: "Y",
  SCISSOR: "Z",
};

const GAME_POINTS = {
  WON: 6,
  DRAW: 3,
  LOST: 0,
};

const PLAYED_POINTS = {
  [MY_PLAY["ROCK"]]: 1,
  [MY_PLAY["PAPER"]]: 2,
  [MY_PLAY["SCISSOR"]]: 3,
};

const checkRound = {
  [`${ELF_PLAY.ROCK} ${MY_PLAY.ROCK}`]: GAME_POINTS.DRAW,
  [`${ELF_PLAY.ROCK} ${MY_PLAY.PAPER}`]: GAME_POINTS.WON,
  [`${ELF_PLAY.ROCK} ${MY_PLAY.SCISSOR}`]: GAME_POINTS.LOST,
  [`${ELF_PLAY.PAPER} ${MY_PLAY.ROCK}`]: GAME_POINTS.LOST,
  [`${ELF_PLAY.PAPER} ${MY_PLAY.PAPER}`]: GAME_POINTS.DRAW,
  [`${ELF_PLAY.PAPER} ${MY_PLAY.SCISSOR}`]: GAME_POINTS.WON,
  [`${ELF_PLAY.SCISSOR} ${MY_PLAY.ROCK}`]: GAME_POINTS.WON,
  [`${ELF_PLAY.SCISSOR} ${MY_PLAY.PAPER}`]: GAME_POINTS.LOST,
  [`${ELF_PLAY.SCISSOR} ${MY_PLAY.SCISSOR}`]: GAME_POINTS.DRAW,
};

function mySolutions(input) {
  return input
    .split("\n")
    .reduce(
      (acc, play) => acc + checkRound[play] + PLAYED_POINTS[play.split(" ")[1]],
      0
    );
}

console.time("MY  SOLUTION");

const inputTxt = fs.readFileSync("../input.txt");

const totalPoints = mySolutions(inputTxt.toString());

console.log(`[MY SOLUTION IN JS] -The total points ${totalPoints}.`);

console.timeEnd("MY  SOLUTION");
