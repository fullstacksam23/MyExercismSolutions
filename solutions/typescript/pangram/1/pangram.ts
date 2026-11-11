export function isPangram(s:string):boolean {
  const a: number[] = Array(26).fill(0);
  s = s.toLowerCase();
  for(const char of s){
    const idx = char.charCodeAt(0) - 97;
    if (idx >= 0 && idx < 26) a[idx] += 1;
  }
  for(let i = 0; i < 26; i++){
    if(a[i] === 0){
      return false;
    }
  }
  return true;
}
