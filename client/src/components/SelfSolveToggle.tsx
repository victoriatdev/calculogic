import { useEffect, useState } from "react";
import { FaCheckCircle } from "react-icons/fa";
import Tooltip from "./Tooltip";

const SelfSolveToggle = () => {

    const [selfSolve, setSelfSolve] = useState(() => {
        const saved = localStorage.getItem("self-solve-enabled") || "false";
        const initialValue = JSON.parse(saved);
        return initialValue || false;
    });

    const tooltipText : string = "This toggle will toggle functionality to send only the topmost level of sequents over to the server to prove and 'step-through' the proof."

    useEffect(() => {
        localStorage.setItem("self-solve-enabled", JSON.stringify(selfSolve));
    });

    return (
        <div className="flex w-1/3 flex-col items-center"> 
            <div className="flex justify-between">
                <span>Enable Self-Solve?</span>
                <Tooltip text={tooltipText}/>
            </div>
            <button aria-checked={selfSolve} className="text-(--color-ui-normal) aria-checked:text-(--color-gr) hover:text-(--color-action)" onClick={() => {
                setSelfSolve(!selfSolve);
            }}>
                <FaCheckCircle />
            </button>
            
        </div>
    )
}

export default SelfSolveToggle;