import Markdown from "react-markdown";

const markdown = `# Hello!

Welcome to *Calculogic* - a tool designed for people interested in logic to prove logical formulae in an easy and understanding way.

`;

export default function Home() {
  return (
    <div className="flex items-center flex-col justify-between gap-20 text-(--color-tx-normal)">
      <div className="flex">
        Calculogic
      </div>
      <div>
        <Markdown>{markdown}</Markdown>
      </div>
    </div>
  );
}
