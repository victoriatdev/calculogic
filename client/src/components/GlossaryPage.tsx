import Markdown from "react-markdown";
import RenderSC from "./RenderSC";
import { ProofNode } from "./GentzenTree";
import { v4 } from "uuid";
import { formatInput } from "../lib/utils";

const markdown = `First-order logic can be solved via two competing systems: *Natural Deduction* and *Sequent Calculus*. This page will act as a brief explaination of how to solve a proof using these systems,
and how the sets of rules differ to allow the server to verify your formulae you input.

**Sequent Calculus**

The Sequent Calclus is a particular style of first-order logic, where sequents are proven directly as consequence of smaller, simpler sequents. These sequents are sets of terms and operators,
comprised in particular of two parts that surround a turnstile (|-) - the antecedent (or hypothesis) on the left, and the succedent (or conclusion) on the right. The turnstile can be read as "Left implies Right".

> Axioms

`;

const assumption = {
  sequent: {
    Antecedent: formatInput("X, A ^ B"),
    Succedent: "Y",
    InferenceRule: "A"
  },
  id: v4(),
  proof: [],
}

const leftConjunction = {
  sequent: {
    Antecedent: "",
    Succedent: formatInput("X, A ^ B"),
    InferenceRule: formatInput("^L"),
  },
  id: v4(),
  proof: [
    {
      sequent: {
        Antecedent: formatInput("X, A, B"),
        Succedent: formatInput("Y"),
        // InferenceRule: "",
      },
      id: v4(),
      proof: [],
    },
  ],
};
const rightConjunction = {
  sequent: {
    Antecedent: formatInput("X1, X2"),
    Succedent: formatInput("A ^ B, Y1, Y2"),
    InferenceRule: formatInput("^R"),
  },
  id: v4(),
  proof: [
    {
      sequent: {
        Antecedent: formatInput("X1"),
        Succedent: formatInput("A, Y1"),
        // InferenceRule: "",
      },
      id: v4(),
      proof: [],
    },
    {
      sequent: {
        Antecedent: formatInput("X2"),
        Succedent: formatInput("B, Y2"),
        // InferenceRule: "",
      },
      id: v4(),
      proof: [],
    },
  ],
};

const leftDisjunction = {
  sequent: {
    Antecedent: formatInput("X1, X2, A v B"),
    Succedent: formatInput("Y1, Y2"),
    InferenceRule: formatInput("vL"),
  },
  id: v4(),
  proof: [
    {
      sequent: {
        Antecedent: formatInput("X1, A"),
        Succedent: formatInput("Y1"),
        // InferenceRule: "",
      },
      id: v4(),
      proof: [],
    },
    {
      sequent: {
       Antecedent: formatInput("X2, B"),
       Succedent: formatInput("Y2"), 
      },
      id: v4(),
      proof: [],
    }
  ],
};

const rightDisjunction = {
  sequent: {
    Antecedent: formatInput("X"),
    Succedent: formatInput("A v B, Y"),
    InferenceRule: formatInput("vR"),
  },
  id: v4(),
  proof: [
    {
      sequent: {
        Antecedent: formatInput("X"),
        Succedent: formatInput("A, B, Y"),
        // InferenceRule: "",
      },
      id: v4(),
      proof: [],
    },
  ],
};

const leftNegation = {
  sequent: {
    Antecedent: formatInput("X"),
    Succedent: formatInput("A, Y"),
    InferenceRule: formatInput("¬L"),
  },
  id: v4(),
  proof: [
    {
      sequent: {
        Antecedent: formatInput("X, ¬A"),
        Succedent: formatInput("Y"),
        // InferenceRule: "",
      },
      id: v4(),
      proof: [],
    },
  ],
};
const rightNegation = {
  sequent: {
    Antecedent: formatInput("X"),
    Succedent: formatInput("¬A, Y"),
    InferenceRule: formatInput("¬L"),
  },
  id: v4(),
  proof: [
    {
      sequent: {
        Antecedent: formatInput("X, A"),
        Succedent: formatInput("Y"),
        // InferenceRule: "",
      },
      id: v4(),
      proof: [],
    },
  ],
};

const rightImplication = {
  sequent: {
    Antecedent: formatInput("X"),
    Succedent: formatInput("A -> B, Y"),
    InferenceRule: formatInput("->R"),
  },
  id: v4(),
  proof: [
    {
      sequent: {
        Antecedent: formatInput("X, A"),
        Succedent: formatInput("B, Y"),
        // InferenceRule: "",
      },
      id: v4(),
      proof: [],
    },
  ],
};
const leftImplication= {
  sequent: {
    Antecedent: formatInput("X1, X2, A -> B"),
    Succedent: formatInput("Y1, Y2"),
    InferenceRule: formatInput("->L"),
  },
  id: v4(),
  proof: [
    {
      sequent: {
        Antecedent: formatInput("X1"),
        Succedent: formatInput("A, Y1"),
        // InferenceRule: "",
      },
      id: v4(),
      proof: [],
    },
    {
      sequent: {
        Antecedent: formatInput("X2, B"),
        Succedent: formatInput("Y2"),
        // InferenceRule: "",
      },
      id: v4(),
      proof: [],
    },
  ],
};

const GlossaryPage = () => {
  return (
    <div className="flex items-center flex-col justify-between gap-10 text-(--color-tx-normal)">
      <div className="flex">Glossary</div>
      <div className="items-center text-wrap flex flex-col gap-5 justify-between">
        {/* <Markdown>{markdown}</Markdown> */}
        <p>First-order logic can be solved via two competing systems: Natural Deduction and Sequent Calculus.</p>
        <p className="font-bold">Input Codes for Operators</p>
        <div>
          <p>Implication Operator (→): "\to", -{`>`}, {`=>`} </p>
          <p>Negation Operator (¬): "\neg", {`¬`}, {`~`}, {`!`} </p>
          <p>Conjunction Operator (∧): "\land", {`&`}, {`^`} </p>
          <p>Disjunction Operator (∨): "\lor", {`||`}, lowercase {`v`}, uppercase {`V`} </p>
          <p>Turnstile (⊢): {`|-`}</p>
        </div>
        <h1 className="font-bold">Sequent Calculus Rules:</h1>
        <RenderSC  proofTree={assumption} />
        <div className="flex gap-10 w-full">
          <RenderSC  proofTree={leftConjunction} />
          <RenderSC  proofTree={rightConjunction} />
        </div>
         <div className="flex gap-10 w-full">
          <RenderSC  proofTree={leftDisjunction} />
          <RenderSC  proofTree={rightDisjunction} />
        </div>
         <div className="flex gap-10 w-full">
          <RenderSC  proofTree={leftNegation} />
          <RenderSC  proofTree={rightNegation} />
        </div>
         <div className="flex gap-10 w-full">
          <RenderSC  proofTree={leftImplication} />
          <RenderSC  proofTree={rightImplication} />
        </div>
       
      </div>
    </div>
  );
};

export default GlossaryPage;
