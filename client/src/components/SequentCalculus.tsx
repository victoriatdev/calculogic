import SelfSolveToggle from "./SelfSolveToggle";
import SubmitProof from "./SubmitProof";
import RenderSC from "./RenderSC";
import { v4 } from "uuid";
import { useState } from "react";
import { ProofNode } from "./GentzenTree";

const t = {
  sequent: {
    Antecedent: "p,q",
    Succedent: "Kpq;1,2",
    InferenceRule: "ki",
  },
  id: v4(),
  proof: [
    {
      sequent: {
        Antecedent: "p",
        Succedent: "p",
        InferenceRule: "a",
      },
      id: v4(),
      proof: [],
    },
    {
      sequent: {
        Antecedent: "q",
        Succedent: "q",
        InferenceRule: "a",
      },
      id: v4(),
      proof: [],
    },
  ],
};

const t2 = {
  sequent: {
    Antecedent: "",
    Succedent: "p->q->p^q",
    InferenceRule: "->I",
  },
  id: v4(),
  proof: [
    {
      sequent: {
        Antecedent: "p",
        Succedent: "q->p^q",
        InferenceRule: "->I",
      },
      id: v4(),
      proof: [
        {
          sequent: {
            Antecedent: "p,q",
            Succedent: "p^q",
            InferenceRule: "^I",
          },
          id: v4(),
          proof: [
            {
              sequent: {
                Antecedent: "p,q",
                Succedent: "p",
                InferenceRule: "a",
              },
              id: v4(),
              proof: [],
            },
            {
              sequent: {
                Antecedent: "p,q",
                Succedent: "q",
                InferenceRule: "a",
              },
              id: v4(),
              proof: [],
            },
          ],
        },
      ],
    },
  ],
};

const SequentCalculus = () => {
  const [tree, setTree] = useState<ProofNode>();

  const handleChildData = (data: any) => {
    setTree(data);
  };

  return (
    // global container
    <div className="w-full flex items-center justify-between px-8 text-(--color-tx-normal)">
      <div>Cached Examples</div>
      <div>
        <SubmitProof proofTree={t2} passToChild={handleChildData} />
        <RenderSC proofTree={t2} />
      </div>
      <SelfSolveToggle />
    </div>
  );
};

export default SequentCalculus;

// sc /> => stores the prooftree datastructure build from "CreateProof />"
// submitproof takes the proof from the ds and submits it
// rendersc takes the fully formed ds and renders it
