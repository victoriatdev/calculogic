import Markdown from "react-markdown";

const markdown = `First-order logic can be solved via two competing systems: *Natural Deduction* and *Sequent Calculus*. This page will act as a brief explaination of how to solve a proof using these systems,
and how the sets of rules differ to allow the server to verify your formulae you input.`;

const GlossaryPage = () => {
  return (
    <div className="flex items-center flex-col justify-between gap-20 text-(--color-tx-normal)">
      <div className="flex">Glossary</div>
      <div className="size-1/2 text-wrap">
        <Markdown>{markdown}</Markdown>
      </div>
    </div>
  );
};

export default GlossaryPage;
