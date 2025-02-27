import SelfSolveToggle from "./SelfSolveToggle";

const NaturalDeduction = () => {
  
  return (
    // global container
    <div className="w-full flex items-center justify-between px-8 text-(--color-tx-normal)">
      <div>
        Cached Examples 
      </div>
      <div>
        Proof to Solve
      </div>
      <SelfSolveToggle />
    </div>
  );
}

export default NaturalDeduction;