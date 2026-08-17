// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

/**
 * Determines how long it takes to prepare a certain juice.
 *
 * @param {string} name
 * @returns {number} time in minutes
 */
const juiceTime = new Map(
  [
    ['Pure Strawberry Joy', 0.5],
    ['Energizer', 1.5],
    ['Green Garden', 1.5],
    ['Tropical Island', 3],
    ['All or Nothing', 5]
  ]
);
export function timeToMixJuice(name) {
  if (juiceTime.has(name)) {
    return juiceTime.get(name);
  } 
  return 2.5;
}

/**
 * Calculates the number of limes that need to be cut
 * to reach a certain supply.
 *
 * @param {number} wedgesNeeded
 * @param {string[]} limes
 * @returns {number} number of limes cut
 */
export function limesToCut(wedgesNeeded, limes) {
  let idx = 0;
  let wedges = 0;
  while (idx < limes.length && wedges < wedgesNeeded) {
    switch(limes[idx]) {
      case 'small':
        wedges += 6;
        break;
      case 'medium':
        wedges += 8;
        break;
      case 'large':
        wedges += 10;
        break;
    }
    idx+=1;
  }
  return idx;
}

/**
 * Determines which juices still need to be prepared after the end of the shift.
 *
 * @param {number} timeLeft
 * @param {string[]} orders
 * @returns {string[]} remaining orders after the time is up
 */
export function remainingOrders(timeLeft, orders) {
  let idx = 0;
  while (timeLeft > 0) {
    timeLeft -= timeToMixJuice(orders[idx]);
    idx+=1
  }
  return orders.slice(idx, orders.length);
}
