"use client";

import { Task } from '@/data/task';
import {JSX, useState} from "react";
import {useRouter} from "next/navigation";

export default function List({tasks}: {tasks: Array<Task>}): JSX.Element {
    const router = useRouter();
    const [clicked, setClicked] = useState(false);

    const onClick = async function(number: number) {
        await fetch('api/task?number='+number, {
            method: 'GET',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            next: {
                revalidate: 0
            }
        });
        router.refresh();
    }

    return (
        <ul className="grid grid-cols-4 gap-4">
            {tasks.map((task: Task) => (
                <li key={task.number}
                    className="task group rounded-lg border border-transparent px-5 py-4 transition-colors hover:border-gray-300 hover:bg-gray-100 hover:dark:border-neutral-700 hover:dark:bg-neutral-800/30"
                >
                    <h2 className={`mb-3 text-2xl font-semibold`}>
                        {task.title}
                        <span className="inline-block transition-transform group-hover:translate-x-1 motion-reduce:transform-none">
                     -&gt;
                   </span>
                    </h2>
                    <p className={`m-0 max-w-[30ch] text-sm opacity-50`}>
                        {task.description}
                    </p>
                    <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
                        onClick={() => onClick(task.number)}
                    >Done</button>
                </li>
            ))}
        </ul>
    )
}

