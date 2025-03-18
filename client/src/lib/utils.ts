import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const formatInput = (input: string) => {
  var formattedInput = input;
  console.log("trigger");

  if (input.includes("->") || input.includes("=>")) {
    formattedInput = formattedInput.replaceAll("->", String.fromCodePoint(8594));
  }
  if (input.includes("^") || input.includes("&")) {
    formattedInput = formattedInput.replaceAll(
      /[\^&]/g,
      String.fromCodePoint(8743)
    );
  }
  if (input.includes("v") || input.includes("V") || input.includes("||")) {
    formattedInput = formattedInput.replaceAll(
      /[vV||]/g,
      String.fromCodePoint(8744)
    );
  }
  if (input.includes("!") || input.includes("¬") || input.includes("~")) {
    formattedInput = formattedInput.replaceAll(
      /[!~¬]/g,
      String.fromCodePoint(172)
    );
  }
  if (input.includes("|-")) {
    formattedInput = formattedInput.replace("|-", String.fromCodePoint(8866));
  }
  if (input.includes("forall")) {
    formattedInput = formattedInput.replace(
      "forall",
      String.fromCodePoint(8704)
    );
  }
  if (input.includes("exists")) {
    formattedInput = formattedInput.replace(
      "exists",
      String.fromCodePoint(8707)
    );
  }

  return formattedInput;
};
