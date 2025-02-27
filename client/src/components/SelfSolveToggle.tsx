import React, { useEffect, useState } from "react";
import { FaCheckCircle } from "react-icons/fa";

const SelfSolveToggle = () => {

    const [selfSolve, setSelfSolve] = useState(() => {
        const saved = localStorage.getItem("self-solve-enabled") || "false";
        const initialValue = JSON.parse(saved);
        return initialValue || false;
    });

    useEffect(() => {
        localStorage.setItem("self-solve-enabled", JSON.stringify(selfSolve));
    });

    return (
        <div className="flex flex-col items-center"> 
            Enable Self-Solve?
            <button aria-checked={selfSolve} className="text-(--color-ui-normal) aria-checked:text-(--color-gr) hover:text-(--color-action)" onClick={() => {
                setSelfSolve(!selfSolve);
            }}>
                <FaCheckCircle />
            </button>
        </div>
    )
}

export default SelfSolveToggle;