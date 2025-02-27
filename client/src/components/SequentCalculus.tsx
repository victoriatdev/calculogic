import SelfSolveToggle from "./SelfSolveToggle";

const SequentCalculus = () => {
  return (
    // global container
    <div title="Sequent Calculus" className="w-full flex items-center justify-between px-8 text-(--color-tx-normal)">
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

export default SequentCalculus;
