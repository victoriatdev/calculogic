import { Sequent,  } from "./GentzenTree";

type RootNodeProps = {
  proofTree
};

type NodeProps = {
    
};

const RootNode = (props: RootNodeProps) => {
  const data = props.data;

  return (
    <div className="flex justify-center">
      <span id="left" className="hover:bg-(--color-bl-hover)/25 cursor-pointer">
        {data?.Antecedent}
      </span>
      <span
        id="turnstile"
        className="hover:bg-(--color-re-hover)/25 cursor-pointer"
      >
        {data && String.fromCodePoint(8866)}
      </span>
      <span
        id="right"
        className="hover:bg-(--color-bl-hover)/25 cursor-pointer"
      >
        {data?.Succedent}
      </span>
      <span
        id="rule"
        className="hover:bg-(--color-gr-hover)/25 cursor-pointer text-xs"
      >
        [{data?.InferenceRule}]
      </span>
    </div>
  );
};

const Node = (props: NodeProps) => {
  return <Node />;
};

export default Node;
