interface CardProps {
    title: string
    children: JSX.Element,
    type?: "future" | "complete" | "beta"
}

export const Cards = function(props: {children: JSX.Element[]}) {
    return <div className="grid grid-cols-3 gap-4">
        {props.children}
    </div>
}

export const Card = function(props: CardProps) {
    return <div className="card w-96 bg-base-100 border mt-4 shadow-sm">
        <div className="card-body">
            <h2 className="card-title">
                {props.title}
                {props.type == "future" && <span className="badge badge-lg badge-primary">FUTURE</span>}
                {props.type == "complete" && <span className="badge badge-lg badge-success">COMPLETE</span>}
                {props.type == "beta" && <span className="badge badge-lg badge-accent">BETA</span>}
            </h2>
            {props.children}
        </div>
    </div>
}
