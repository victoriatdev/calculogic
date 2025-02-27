import Markdown from "react-markdown";

const markdown = `
# Custom Ruleset
`;

export default function CustomRuleset() {
  return (
    <div>
      <Markdown>{markdown}</Markdown>
    </div>
  );
}
