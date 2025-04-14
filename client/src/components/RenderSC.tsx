import React, { useState } from "react";
import { ProofNode } from "./GentzenTree";

type RenderSCProps = {
  tree: ProofNode;
};

const Leaf = ({ sequent, width }: any) => {
  const tree = sequent;
  // no more children hence just render the "part"
  // Leaf props: width (passed from parent Node calculated from width / arr items )

  return (
    <div className="pt-1">
      <div className="flex justify-center">
        <span
          id="left"
          className="hover:bg-(--color-bl-hover)/25 cursor-pointer"
        >
          {tree && tree.Antecedent}
        </span>
        <span
          id="turnstile"
          className="hover:bg-(--color-re-hover)/25 cursor-pointer"
        >
          {tree && String.fromCodePoint(8866)}
        </span>
        <span
          id="right"
          className="hover:bg-(--color-bl-hover)/25 cursor-pointer"
        >
          {tree && tree.Succedent}
        </span>
        <span
          id="rule"
          className="hover:bg-(--color-gr-hover)/25 cursor-pointer text-xs"
        >
          [{tree && tree.InferenceRule}]
        </span>
      </div>
    </div>
  );
};

const calcWidth = (p, width) => {
  return width / p.length;
};

const TreeNode = ({ proofTree, width }: any) => {
  let children = null;

  if (proofTree.proof && proofTree.proof.length) {
    const w = calcWidth(proofTree.proof, width);
    children = (
      <div className="flex">
        {proofTree.proof.map((p) => (
          <TreeNode proofTree={p} key={p.id} width={w} />
        ))}
      </div>
    );
  }

  return (
    <div className="flex-wrap w-full">
      {children}
      {proofTree.sequent && <Leaf sequent={proofTree.sequent} width={width} />}
    </div>
  );
};

const RenderSC = ({ proofTree }: any) => {
  const tree = proofTree;

  // are there any elements in tree.proof? if so, we need to build another "node"
  // if not, then we build a leaf

  // build top level node last
  console.log(tree);

  return (
    <div className="border rounded-sm w-full h-full">
      {tree && <TreeNode proofTree={tree} key={tree.Id} width={100} />}
    </div>
  );
};

export default RenderSC;
