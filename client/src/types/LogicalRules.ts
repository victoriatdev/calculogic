enum NDLogicalRules {
  AXIOM = "Axiom",
  LEFT_NEGATION = "Left Negation",
  LEFT_CONJUNCTION = "Left Conjunction",
  LEFT_DISJUNCTION = "Left Disjunction",
  LEFT_IMPLICATION = "Left Implication",
  LEFT_UNIVERSAL_QUANTIFICATION = "Left Universal Quantification",
  LEFT_EXISTENTIAL_QUANTIFICATION = "Left Existential Quantification",
  RIGHT_NEGATION = "Right Negation",
  RIGHT_CONJUNCTION = "Right Conjunction",
  RIGHT_DISJUNCTION = "Right Disjunction",
  RIGHT_IMPLICATION = "Right Implication",
  RIGHT_UNIVERSAL_QUANTIFICATION = "Right Universal Quantification",
  RIGHT_EXISTENTIAL_QUANTIFICATION = "Right Existential Quantification",
}

enum InferenceRule {
  ASSUMPTION = "A",
  NEGATION_ELIMINATION = "¬E",
  NEGATION_INTRODUCTION = "¬I",
  DISJUNCTION_ELIMINATION = "∨E",
  DISJUNCTION_INTRODUCTION = "∨I",
  CONJUNCTION_ELIMINATION = "∧E",
  CONJUNCTION_INTRODUCTION = "∧I",
  CONDITIONAL_ELIMINATION = "→E",
  CONDITIONAL_INTRODUCTION = "→I",
}

function getEnumKeys<
  T extends string,
  TEnumValue extends string | number
>(enumVariable: {
  [key in T]: TEnumValue;
}) {
  return Object.keys(enumVariable) as Array<T>;
}

export { getEnumKeys, NDLogicalRules, InferenceRule };
