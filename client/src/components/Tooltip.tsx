import { useState } from "react";
import { HiQuestionMarkCircle } from "react-icons/hi";


const Tooltip = ({ text }) => {
    const [hover, setHover] = useState(false);

    const toggleHover = () => {
        setHover(!hover);
    }    
    
    return (
        <>
            <div
                    className="ml-5"
                    onMouseEnter={toggleHover}
                    onMouseLeave={toggleHover}
                >
                    <HiQuestionMarkCircle data-tooltip-target="tooltip" />
                    { hover && <div id="tooltip" role="tooltip" className="z-10 transform -translate-x-1/3 w-1/6 text-sm absolute border border-(--color-ui-active) rtl:text-right shadow-lg rounded-lg tootip p-2 bg-(--color-ui-normal)">
                        {text}
                       </div>
                        }
            </div>
        </>
    )
}

export default Tooltip;