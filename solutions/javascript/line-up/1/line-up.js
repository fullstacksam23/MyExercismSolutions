export const format = (name, number) => {
  let suffix = "";

  const lastTwo = number % 100;
  const lastOne = number % 10;

  if (lastTwo >= 11 && lastTwo <= 13) {
    suffix = "th";
  } else {
    switch (lastOne) {
      case 1:
        suffix = "st";
        break;
      case 2:
        suffix = "nd";
        break;
      case 3:
        suffix = "rd";
        break;
      default:
        suffix = "th";
    }
  }

  return name + ", you are the " + number + suffix +
    " customer we serve today. Thank you!";
};