const colorAmt: Record<string, number> = {
  "black": 0,
  "brown": 1,
  "red": 2,
  "orange": 3,
  "yellow": 4,
  "green": 5,
  "blue": 6,
  "violet": 7,
  "grey": 8,
  "white": 9
};
const metricPrefix: Record<number, string> = {
  0: "",
  1: "kilo",
  2: "mega",
  3: "giga"
};
export function decodedResistorValue(colors: string[]): string {
  let value: number = (colorAmt[colors[0]]*10 + colorAmt[colors[1]]) * (10**colorAmt[colors[2]]);
  let i = 0;
  while(value >= 1000 && value%1000 === 0){
    value /= 1000;
    i++;
  }
  return `${value} ${metricPrefix[i]}ohms`;
}
