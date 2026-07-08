// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

/**
 * Calculates the total bird count.
 *
 * @param {number[]} birdsPerDay
 * @returns {number} total bird count
 */
export function totalBirdCount(birdsPerDay) {
  let s = 0;
  for(let i = 0; i < birdsPerDay.length; i++) {
    s += birdsPerDay[i];
  }
  return s;
}

/**
 * Calculates the total number of birds seen in a specific week.
 *
 * @param {number[]} birdsPerDay
 * @param {number} week
 * @returns {number} birds counted in the given week
 */
export function birdsInWeek(birdsPerDay, week) {
  const start = (week-1)*7;
  const end = start + 7;
  let s = 0;
  for (let i = start; i < end; i++) {
    s += birdsPerDay[i];
  }
  return s;
}

/**
 * Fixes the counting mistake by increasing the bird count
 * by one for every second day.
 *
 * @param {number[]} birdsPerDay
 * @returns {void} should not return anything
 */
export function fixBirdCountLog(birdsPerDay) {
  for (let i = 0; i < birdsPerDay.length; i++){
    if ((i&1) === 0){
      birdsPerDay[i]+=1;
    }
  }
  return birdsPerDay;
}
