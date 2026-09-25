export function format(name: string, number: number): string {
  let suff = "th";
  const last = number%10;
  let secondLast = -1;
  if(number > 9){
    secondLast = (Math.floor(number/10))%10;  
  }
  if(last === 1 && secondLast !== 1){
    suff = "st";
  }else if(last === 2 && secondLast !== 1){
    suff = "nd";
  }else if(last === 3 && secondLast !== 1){
    suff = "rd";
  }
  return `${name}, you are the ${number}${suff} customer we serve today. Thank you!`;
}
