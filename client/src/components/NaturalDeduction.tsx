import SelfSolveToggle from "./SelfSolveToggle";
import SubmitProof from "./SubmitProof";
import RenderND from "./RenderND.tsx";

const NaturalDeduction = () => {
  return (
    // global container
    <div className="w-full flex items-center justify-between px-8 text-(--color-tx-normal)">
      <div>Cached Examples</div>
      <div>
        <SubmitProof proofType="natural-deduction" />
        <RenderND />
      </div>
      <SelfSolveToggle />
    </div>
  );
};

export default NaturalDeduction;
