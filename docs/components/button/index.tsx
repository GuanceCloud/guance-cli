interface ButtonProps {
    href: string
    title: string
    type: string
}

export default function Button(props: ButtonProps) {
    return <div className="mt-6">
        <a href={props.href} target="_blank">
            <button className={`btn btn-${props.type}`}>{props.title}</button>
        </a>
    </div>
}
