import Markdown from "react-markdown";

const markdown = `# Hello!

Welcome to *NAME* - a tool designed for people interested in logic to prove logical formulae in an easy and understanding way.

`;

export default function SequentCalculus() {
  return (
    <div>
      <Markdown>{markdown}</Markdown>
    </div>
  );
}
