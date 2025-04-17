import SelfSolveToggle from "./SelfSolveToggle";
import SubmitProof from "./SubmitProof";
import RenderSC from "./RenderSC";
import { v4 } from "uuid";
import { useEffect, useState } from "react";
import { ProofNode, Sequent } from "./GentzenTree";
import Indicator from "./Indicator";
import CachedExample, {
  type CachedExample as TCachedExample,
} from "./CachedExample";
import { formatInput } from "../lib/utils";
import Tooltip from "./Tooltip";
import { sleep } from "bun";
import { NavLink } from "react-router-dom";

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
    Succedent: formatInput("P->Q->P^Q"),
    InferenceRule: "->I",
  },
  id: v4(),
  proof: [
    {
      sequent: {
        Antecedent: "P",
        Succedent: formatInput("Q->P^Q"),
        InferenceRule: "->I",
      },
      id: v4(),
      proof: [
        {
          sequent: {
            Antecedent: "P,Q",
            Succedent: formatInput("P^Q"),
            InferenceRule: "^I",
          },
          id: v4(),
          proof: [
            {
              sequent: {
                Antecedent: "P,Q",
                Succedent: "P",
                InferenceRule: "a",
              },
              id: v4(),
              proof: [],
            },
            {
              sequent: {
                Antecedent: "P,Q",
                Succedent: "Q",
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

const buildTree = (serverTree: any) => {
  console.log(serverTree);
  console.log(serverTree.Children.length)

  const sequent: Sequent = {
    Antecedent: serverTree.Data[0],
    Succedent: serverTree.Data[1],
    InferenceRule: serverTree.Rule,
  };
  
  if (serverTree.Children.length == 0) {
    const builtLeaf : ProofNode = new ProofNode({
      id: serverTree.id,
      sequent: sequent,
      proof: []
    })
    return builtLeaf
  }

  const proof = []

  serverTree.Children.forEach((child) => {
    const childNode = buildTree(child);
    proof.push(childNode)
  })

  console.log(proof)

  const builtTree: ProofNode = new ProofNode({
    id: serverTree.Id,
    sequent: sequent,
    proof: proof
  });

  console.log(builtTree)

  return builtTree;
};

const SequentCalculus = () => {
  const [tree, setTree] = useState<ProofNode>();
  const [requestStatus, setRequestStatus] = useState("");
  const [cachedExamples, setCachedExamples] = useState<TCachedExample[]>([]);
  const [expressionToLoad, setExpressionToLoad] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const handleChildData = (data: any) => {
    // console.log(data);
    const tree = buildTree(data);
    // const tree = new ProofNode(data)
    console.log(tree);

    setTree(tree);
  };

  // const examples : string[] = ["P->Q->P^Q", "¬AvA"];

  const examples: TCachedExample[] = [
    {
      expression: "P->Q->P^Q",
      variant: "SC",
    },
    {
      expression: "¬AvA",
      variant: "ND",
    },
  ];

  const cachedExampleText: string =
    "These cached examples can be used as examples of how to format proofs, and they will always work. They are labelled with the system they work for. Simply click on an example to load it into the submission box.";

  const refreshCachedExamples = async () => {
    setIsLoading(true);
    await new Promise((f) => setTimeout(f, 500));
    setIsLoading(false);
    setCachedExamples(examples);
  };

  const handleRequestUpdate = (status: string) => {
    setRequestStatus(status);
  };

  const handleLoadExpression = (expression: string) => {
    // console.log(expression)
    handleRequestUpdate("");
    setExpressionToLoad(formatInput(expression));
  };

  return (
    // global container
    <div className="w-full flex items-center justify-between px-8 text-(--color-tx-normal)">
      <div className="space-y-2 h-[calc(100vh-80px)] border-r w-1/3 mr-6">
        <div className="flex justify-between p-2">
          <div className="flex align-center items-center">
            <span>Cached Examples</span>
            <Tooltip text={cachedExampleText} />
          </div>
          <button
            type="button"
            onClick={refreshCachedExamples}
            className="p-2 hover:bg-(--color-ui-hover) active:bg-(--color-ui-active) cursor-pointer rounded-sm bg-(--color-ui-normal)"
          >
            Refresh
          </button>
        </div>
        <div className="space-y-2">
          {isLoading
            ? "Loading..."
            : cachedExamples.map((example, id) => (
                <div>
                  <CachedExample
                    expression={example.expression}
                    variant={example.variant}
                    loadExpression={handleLoadExpression}
                    key={id}
                  />
                </div>
              ))}
        </div>
      </div>
      <div className="space-y-4 w-1/3 min-h-[calc(100vh-80px)] pt-10 flex flex-col items-center">
        <div>
          Input your Sequent Calculus formula you wish to be solved below. Please see the <NavLink className="text-(--color-bl)" to="/glossary">Glossary </NavLink>
           on more information on how to format your input and how the Sequent Calculus works;
           or refresh the examples on the left to see some common examples you can click on to load into the solver. 
        </div>
        <SubmitProof
          proofTree={tree}
          passToChild={handleChildData}
          handleLoadExpression={handleLoadExpression}
          expressionToLoad={expressionToLoad}
          setRequestStatus={handleRequestUpdate}
        />
        {requestStatus == "proven" &&
            <RenderSC proofTree={tree} />
          }
        {requestStatus && <Indicator indicatorStatus={requestStatus} />}
      </div>
      <SelfSolveToggle />
    </div>
  );
};

export default SequentCalculus;

// sc /> => stores the prooftree datastructure build from "CreateProof />"
// submitproof takes the proof from the ds and submits it
// rendersc takes the fully formed ds and renders it
