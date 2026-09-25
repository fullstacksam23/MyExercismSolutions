export function toRna(dna: string): string {
  const mapping: Record<string, string> = {
    "G": "C",
    "C": "G",
    "T": "A",
    "A": "U"
  };
  const chars: string[] = [];
    for(const c of dna){
      if(!(c in mapping)){
        throw new Error("Invalid input DNA.");
      }
      chars.push(mapping[c]);
    }
    return chars.join("");
}