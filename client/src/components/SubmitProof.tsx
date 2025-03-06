import { useState } from "react";
import { GentzenTree, ProofNode, type Sequent } from "./GentzenTree";
import { getEnumKeys, InferenceRule } from "../types/LogicalRules";
import RenderSC from "./RenderSC";
import { v4 } from "uuid";
import { formatInput } from "../lib/utils";

const SubmitProof = ({ proofTree, passToChild }: any) => {
  const [inputtedProof, setInputtedProof] = useState("");
  // const [tree, setTree] = useState<ProofNode>();
  const [rule, setRule] = useState<InferenceRule>(InferenceRule.ASSUMPTION);

  const handleSubmit = (e: any) => {
    e.preventDefault();

    fetch(`http://localhost:1323/sequent-calculus`, {
      method: "POST",
      body: JSON.stringify(proofTree),
      headers: {
        "Content-Type": "application/json",
      },
    });
  };

  const handleKeyDown = (event: any) => {
    if (event.key === "Enter") {
      handleAdd();
    }
  };

  const handleInput = (input) => {
    setInputtedProof(formatInput(input.target.value));
  };

  const handleAdd = () => {
    const splitProof = inputtedProof.split(String.fromCodePoint(8866));
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

    setInputtedProof("");
  };

  return (
    <div>
      <label htmlFor="proof" className="block text-(--color-tx-normal)">
        Proof to Solve:
        {String.fromCodePoint(8594)}
      </label>
      <div className="mt-2">
        <AddItem
          inputtedProof={inputtedProof}
          onChange={(input) => handleInput(input)}
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
    <div className="flex items-center rounded-md bg-white pl-3 outline-1 -outline-offset-1 outline-gray-300 has-[input:focus-within]:outline-2 has-[input:focus-within]:-outline-offset-2 has-[input:focus-within]:outline-indigo-600">
      <input
        id="proof"
        name="proof"
        type="text"
        value={inputtedProof}
        onChange={onChange}
        onKeyDown={handleKeyDown}
        className="block min-w-0 grow py-1.5 pr-3 pl-1 text-base text-(--color-tx-normal) focus:outline-none"
      ></input>
      <div className="grid shrink-0 grid-cols-1 focus-within:relative">
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
      </div>
      <button type="button" onClick={handleSubmit}>
        Submit
      </button>
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
