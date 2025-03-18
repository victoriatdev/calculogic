type BadgeProps = {
    variant: string
}

const Badge = (props : BadgeProps) => {
    
    return (
        <div>
            {props.variant == "SC" && <span className="shadow-sm bg-blue-100 text-blue-800 text-xs font-medium me-2 px-2.5 py-0.5 rounded-sm dark:bg-blue-900 dark:text-blue-300">Sequent Calculus</span>}
            {props.variant == "ND" && <span className="shadow-sm bg-green-100 text-green-800 text-xs font-medium me-2 px-2.5 py-0.5 rounded-sm dark:bg-green-900 dark:text-green-300">Natural Deduction</span>}
        </div>
    )
}

export default Badge