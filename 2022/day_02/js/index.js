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

const checkRoundResult = {
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

const MY_CHOICE = {
  LOSE: "X",
  DRAW: "Y",
  WIN: "Z",
};

const checkRoundChoice = {
  [`${ELF_PLAY.ROCK} ${MY_CHOICE.LOSE}`]: PLAYED_POINTS[MY_PLAY["SCISSOR"]],
  [`${ELF_PLAY.ROCK} ${MY_CHOICE.DRAW}`]: PLAYED_POINTS[MY_PLAY["ROCK"]],
  [`${ELF_PLAY.ROCK} ${MY_CHOICE.WIN}`]: PLAYED_POINTS[MY_PLAY["PAPER"]],
  [`${ELF_PLAY.PAPER} ${MY_CHOICE.LOSE}`]: PLAYED_POINTS[MY_PLAY["ROCK"]],
  [`${ELF_PLAY.PAPER} ${MY_CHOICE.DRAW}`]: PLAYED_POINTS[MY_PLAY["PAPER"]],
  [`${ELF_PLAY.PAPER} ${MY_CHOICE.WIN}`]: PLAYED_POINTS[MY_PLAY["SCISSOR"]],
  [`${ELF_PLAY.SCISSOR} ${MY_CHOICE.LOSE}`]: PLAYED_POINTS[MY_PLAY["PAPER"]],
  [`${ELF_PLAY.SCISSOR} ${MY_CHOICE.DRAW}`]: PLAYED_POINTS[MY_PLAY["SCISSOR"]],
  [`${ELF_PLAY.SCISSOR} ${MY_CHOICE.WIN}`]: PLAYED_POINTS[MY_PLAY["ROCK"]],
};

const GAME_POINTS_BY_CHOICE = {
  [MY_CHOICE.WIN]: GAME_POINTS.WON,
  [MY_CHOICE.DRAW]: GAME_POINTS.DRAW,
  [MY_CHOICE.LOSE]: GAME_POINTS.LOST,
};

function getTotalPointsOfRoundResult(input) {
  return input
    .split("\n")
    .reduce(
      (acc, play) =>
        acc + checkRoundResult[play] + PLAYED_POINTS[play.split(" ")[1]],
      0
    );
}

function getTotalPointsOfChoiceResult(input) {
  return input
    .split("\n")
    .reduce(
      (acc, play) =>
        acc +
        checkRoundChoice[play] +
        GAME_POINTS_BY_CHOICE[play.split(" ")[1]],
      0
    );
}

console.time("MY  SOLUTION");

const inputTxt = fs.readFileSync("../input.txt");

const totalPointsOfRoundResult = getTotalPointsOfRoundResult(
  inputTxt.toString()
);
const totalPointsOfChoiceResult = getTotalPointsOfChoiceResult(
  inputTxt.toString()
);

console.log(
  `[MY SOLUTION IN JS] -The total points of round  ${totalPointsOfRoundResult}.`
);
console.log(
  `[MY SOLUTION IN JS] -The total points of choice ${totalPointsOfChoiceResult}.`
);

console.timeEnd("MY  SOLUTION");
