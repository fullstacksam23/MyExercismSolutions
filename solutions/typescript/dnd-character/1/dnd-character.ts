
export class DnDCharacter {
  public strength: number;
  public dexterity: number;
  public constitution: number;
  public intelligence: number;
  public wisdom: number;
  public charisma: number;
  public hitpoints: number;

  constructor() {
    this.strength = DnDCharacter.generateAbilityScore();
    this.dexterity = DnDCharacter.generateAbilityScore();
    this.constitution = DnDCharacter.generateAbilityScore();
    this.intelligence = DnDCharacter.generateAbilityScore();
    this.wisdom = DnDCharacter.generateAbilityScore();
    this.charisma = DnDCharacter.generateAbilityScore();

    this.hitpoints = 10 + DnDCharacter.getModifierFor(this.constitution);
  }

  public static generateAbilityScore(): number {
    let dice: number[] = [];

    for (let i = 0; i < 4; i++) {
      dice.push(Math.ceil(Math.random() * 6));
    }

    let min_val = Number.POSITIVE_INFINITY;
    let sum = 0;

    for (let i = 0; i < dice.length; i++) {
      sum += dice[i];

      if (dice[i] < min_val) {
        min_val = dice[i];
      }
    }

    return sum - min_val;
  }

  public static getModifierFor(abilityValue: number): number {
    return Math.floor((abilityValue - 10) / 2);
  }
}
