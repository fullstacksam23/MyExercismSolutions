
export function hey(message: string): string {
  let whitespace = true;
  let question = message.trimEnd().endsWith("?");
  let yelling = true;
  let hasLetter = false;

  for (const chr of message) {
    if (chr.trim() !== "") {
      whitespace = false;
    }

    if (chr.toLowerCase() !== chr.toUpperCase()) {
      hasLetter = true;

      if (chr !== chr.toUpperCase()) {
        yelling = false;
      }
    }
  }

  if (whitespace) {
    return "Fine. Be that way!";
  }

  yelling = yelling && hasLetter;

  if (question && yelling) {
    return "Calm down, I know what I'm doing!";
  } else if (question) {
    return "Sure.";
  } else if (yelling) {
    return "Whoa, chill out!";
  } else {
    return "Whatever.";
  }
}
