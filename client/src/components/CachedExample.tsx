import { formatInput } from "../lib/utils"
import Badge from "./Badge"

export type CachedExample = {
    expression: string,
    variant: string
}

const CachedExample = ({ expression, variant, loadExpression } : any) => {

    return (
        <div onClick={() => loadExpression(expression)} className="block max-w-sm p-3 border border-(--color-ui-normal) rounded-sm hover:bg-(--color-bg-hover) cursor-pointer active:bg-(--color-selection)">
            {formatInput(expression)}
            <Badge variant={variant}/>   
        </div>
    )
}


export default CachedExample