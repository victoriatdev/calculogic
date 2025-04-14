import { useState } from "react";
import { ProofNode } from "./GentzenTree";
import { getEnumKeys, InferenceRule } from "../types/LogicalRules";
import { formatInput } from "../lib/utils";

const SubmitProof = ({
  proofTree,
  passToChild,
  setRequestStatus,
  handleLoadExpression,
  expressionToLoad,
}: any) => {
  // const [inputtedProof, setInputtedProof] = useState("");
  // const [tree, setTree] = useState<ProofNode>();
  const [rule, setRule] = useState<InferenceRule>(InferenceRule.ASSUMPTION);

  const handleSubmit = async () => {
    const res = await fetch(`http://localhost:1323/proof/sequent-calculus`, {
      method: "POST",
      body: JSON.stringify({ formula: expressionToLoad }),
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!res.ok) {
      setRequestStatus("fail");
    } else if (res.status == 400) {
      setRequestStatus("fail");
    } else {
      const body = await res.json();
      // console.log(body);
      passToChild(body);
      setRequestStatus("proven");
    }
  };

  const handleKeyDown = (event: any) => {
    if (event.key === "Enter") {
      handleAdd();
    }
  };

  const handleInput = (input) => {
    handleLoadExpression(input);
  };

  const handleAdd = () => {
    const splitProof = expressionToLoad.split(String.fromCodePoint(8866));
    if (!splitProof[1] || !splitProof[0]) {
      return;
    }
    const newTree = new ProofNode({
      sequent: {
        Antecedent: splitProof[0],
        Succedent: splitProof[1],
        InferenceRule: rule,
      },
      proof: [],
      id: null,
    });

    // console.log(proofTree);

    if (proofTree) {
      proofTree.addChild(newTree);
      passToChild(proofTree);
      console.log(proofTree);
    } else {
      passToChild(newTree);
      console.log(newTree);
    }

    // console.log(proofTree);

    handleLoadExpression("");
  };

  return (
    <div className="w-full">
      <label htmlFor="proof" className="block text-(--color-tx-normal)">
        Proof to Solve:
      </label>
      <div className="mt-2">
        <AddItem
          inputtedProof={expressionToLoad}
          onChange={(input) => handleInput(input.target.value)}
          onAdd={handleAdd}
          handleKeyDown={handleKeyDown}
          setRule={setRule}
          handleSubmit={handleSubmit}
        />
      </div>
    </div>
  );
};

const AddItem = ({
  onChange,
  onAdd,
  inputtedProof,
  handleKeyDown,
  setRule,
  handleSubmit,
}) => {
  return (
    <div className="flex flex-col items-center">
      <div className="w-full rounded-md bg-white pl-3 outline-1 -outline-offset-1 outline-gray-300 has-[input:focus-within]:outline-2 has-[input:focus-within]:-outline-offset-2 has-[input:focus-within]:outline-indigo-600">
        <input
          id="proof"
          name="proof"
          type="text"
          value={inputtedProof}
          onChange={onChange}
          onKeyDown={handleKeyDown}
          className="block min-w-0 w-full py-1.5 pr-3 pl-1 text-base text-(--flexoki-black) focus:outline-none"
        ></input>
        {/* <div className="grid shrink-0 grid-cols-1 focus-within:relative">
        <select
          id="rule"
          name="rule"
          aria-label="Rule"
          onChange={(input) => {
            setRule(input.target.value);
          }}
          className="col-start-1 row-start-1 w-full appearance-none rounded-md py-1.5 pr-7 pl-3 text-base text-gray-500 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6"
        >
          {getEnumKeys(InferenceRule).map((key, index) => {
            return (
              <option key={index} value={InferenceRule[key]}>
                {InferenceRule[key]}
              </option>
            );
          })}
        </select>
      </div> */}
      </div>
      <div
        onClick={handleSubmit}
        className="w-1/2 border rounded-sm mt-5 p-2 text-center hover:bg-(--color-ui-hover) active:bg-(--color-ui-active) cursor-pointer"
      >
        Submit
      </div>
    </div>
  );
};

// const List = ({ list }: any) => {
//   return (
//     <div>
//       {list.map((item: listElement) => {
//         return <li key={item.id}>{item.content}</li>;
//       })}
//     </div>
//   );
// };

export default SubmitProof;
