import { FaCheckCircle, FaTimesCircle } from "react-icons/fa";

const Indicator = ({ indicatorStatus } : any) => {    

    return (
        <div>
            {indicatorStatus == "proven" &&  
                <div className="flex items-center w-full max-w-xs p-4 space-x-4 rtl:space-x-reverse bg-(--color-gr-hover) rounded-lg shadow-sm">
                    <FaCheckCircle />
                    <div className="ms-3 text-sm text-(--color-tx-normal)">
                        This formula has been proven to be valid.
                    </div>
            </div>}
            {indicatorStatus == "fail" &&
                <div className="flex items-center w-full max-w-xs p-4 space-x-4 rtl:space-x-reverse bg-(--color-re-hover) rounded-lg shadow-sm">
                    <FaTimesCircle />
                    <div className="ms-3 text-sm text-(--color-tx-normal)">
                        This formula is invalid.
                    </div>
                </div>}
        </div>
    )
}

export default Indicator;