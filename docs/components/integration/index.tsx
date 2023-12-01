import {Cards} from 'nextra/components'
import './index.module.css'
import Image from "next/image";
import killerCoda from "../../public/killercoda.svg";


interface TrainingsProps {
    title: string
    description: string
    items: Training[]
}

interface Training {
    title: string
    href: string
}

export default function Trainings(props: TrainingsProps) {
    return <>
        <h1 className="mt-10 mb-4 text-center text-[2.5rem] font-bold tracking-tight">
            {props.title}
        </h1>

        <p className="mb-16 text-center text-lg text-gray-500 dark:text-gray-400">
            {props.description}
        </p>

        <Integrations>
            {props.items.map((item, index) => {
                return <Item key={index} title={item.title} href={item.href}>
                    <div className="bg-black p-10">
                        <Image src={killerCoda} alt="KillerCoda"/>
                    </div>
                </Item>
            })}

        </Integrations>
    </>
}
export const Integrations = (props) => {
    return <Cards {...props} />
}

export const Item = ({children, ...props}) => {
    const cardProps = {
        ...props, ...{
            title: props.title || '未知',
            href: props.href || '',
            icon: props.icon || '',
            image: true,
            arrow: true,
            target: '_blank',
        }
    }
    return <Cards.Card {...cardProps}>
        <div className="bg-white">{children}</div>
    </Cards.Card>
}
