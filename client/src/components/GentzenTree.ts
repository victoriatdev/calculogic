import { v4 } from "uuid";

type Sequent = {
  Antecedent: string;
  Succedent: string;
  InferenceRule: string;
};

type ProofNodeProps = {
  id: string;
  sequent: Sequent;
  proof: ProofNode[];
};

class ProofNode {
  sequent: Sequent;
  id: string;
  parentNode: ProofNode | null;
  proof: ProofNode[];

  constructor(data?: ProofNodeProps) {
    this.sequent = <Sequent>{};
    this.id = v4();
    this.parentNode = null;
    this.proof = [];
    if (data) {
      this.sequent = data.sequent;
      this.id = data.id;
      this.proof = data.proof.map((i) => this.addChild(i));
    }
  }

  public setSequent(data: Sequent) {
    this.sequent = data;
  }

  public addChild(child: ProofNode) {
    let newNode = new ProofNode(child);
    newNode.parentNode = this;
    this.proof.push(newNode);
    return child;
  }

  public addParent() {
    let parent = new ProofNode();
    parent.addChild(this);
    this.replace(parent);
  }

  private replace(node: ProofNode) {
    let length = this.proof.length;
    for (let i = 0; i < length; i++) {
      this.proof[0].remove();
    }
    node.proof.map((n) => this.addChild(n));
    this.sequent = node.sequent;
  }

  private remove() {
    if (this.parentNode) {
      this.parentNode.proof.splice(this.parentNode.proof.indexOf(this), 1);
      if (this.parentNode.proof.length == 0) {
        // rethink?
      }
    }
  }

  public toJSON() {
    return {
      sequent: this.sequent,
      id: this.id,
      proof: this.proof,
    };
  }
}

export { ProofNode, type Sequent };
