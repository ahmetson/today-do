"use client";

import {useButton, useDispatchButton} from "@/app/Button";

export default function Header() {
    const modalIsOpen = useButton()
    const dispatch = useDispatchButton()

    function openModal() {
        console.log(`dispatching true: ${modalIsOpen}`);
        dispatch({
            type: 'TRUE',
        })
    }

    const button = <button
        onClick={openModal}
        data-modal-target="popup-modal" data-modal-toggle="popup-modal"
        className="new-button bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
        New
    </button>

    return (
        <header className="fixed left-0 top-0 flex flex-col w-full border-b border-gray-300 bg-gradient-to-b from-zinc-200 pb-8 pt-4 pl-12 pr-12 backdrop-blur-2xl dark:border-neutral-800 dark:bg-zinc-800/30 dark:from-inherit lg:static lg:w-auto  lg:rounded-xl lg:border lg:bg-gray-200 lg:dark:bg-zinc-800/30">
            <h1>My Short task list <strong>Today</strong> to <strong>Do</strong></h1>
            <sub>Test the SDS Framework <a href="https://github.com/ahmetson/service-lib/tree/service-v1">service-lib@service-v1</a></sub>
            {button}
        </header>
    )
}
