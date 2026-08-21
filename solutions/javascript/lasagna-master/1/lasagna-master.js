/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Implement the functions needed to solve the exercise here.
 * Do not forget to export them so they are available for the
 * tests. Here an example of the syntax as reminder:
 *
 * export function yourFunction(...) {
 *   ...
 * }
 */
export function cookingStatus(remainingTime) {
  if (remainingTime === null || remainingTime === undefined) {
    return 'You forgot to set the timer.';
  }
  if (remainingTime === 0){
    return "Lasagna is done.";
  }
  return "Not done, please wait.";
}
export function preparationTime(layers, avgPrepTime) {
  const t = avgPrepTime ? avgPrepTime : 2;
  return layers.length*t;
}
export function quantities(layers) {
  let noodles = 0;
  let sauce = 0;
  for(let i = 0; i < layers.length; i++){
    if (layers[i] === 'sauce'){
      sauce += 0.2;
    }
    else if(layers[i] === 'noodles'){
      noodles += 50;
    }
  }
  return {
    noodles: noodles,
    sauce: sauce
  };
}
export function addSecretIngredient(firendsList, myList) {
  myList.push(firendsList[firendsList.length-1])
}
export function scaleRecipe(ogRecipe, portions) {
  const recipe = {...ogRecipe};
  for(const k in recipe){
    recipe[k] = (recipe[k]/2)*portions;
  }
  return recipe;
}