'use client';

import React from 'react';
import Modal from 'react-modal';
import { useButton, useDispatchButton } from './Button'
import { XCircleIcon } from '@heroicons/react/24/outline'
import {SubmitHandler, useForm} from "react-hook-form";
import {Task} from "@/data/task";
import { useRouter } from 'next/navigation';

// // Make sure to bind modal to your appElement (https://reactcommunity.org/react-modal/accessibility/)
Modal.setAppElement('body');

export default () => {
    const router = useRouter();
    const modalIsOpen = useButton()
    const dispatch = useDispatchButton()
    const { register, handleSubmit } = useForm<Task>();

    const onSubmit: SubmitHandler<Task> = async (data: Task) => {
        await fetch('api/task', {
            method: 'POST',
            body: JSON.stringify(data),
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            next: {
                revalidate: 0
            }
        });
        closeModal();
        router.refresh();
    }

    function afterOpenModal() {
        // references are now sync and can be accessed.
    }

    function closeModal() {
        console.log(`Dispatching false`);
        dispatch({
            type: 'FALSE',
        })
    }

    return <Modal
        isOpen={modalIsOpen}
        onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        contentLabel="New Task"
        className="new-task-modal"
        // overlayClassName="new-task.ts-modal-overlay"
    >
        <h2>New Task</h2>
        <button className="close-button" onClick={closeModal}><XCircleIcon /></button>
        <form className="mb-4" onSubmit={handleSubmit(onSubmit)}>
            <div className="mb-4">
                <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="title">
                    Title
                </label>
                <input
                    className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    type="text" placeholder="Title" {...register("title")} />
            </div>
            <div className="mb-6">
                <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="description">
                    Description
                </label>
                <textarea rows={4}
                          className="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                          placeholder="Write your task here..."
                          {...register("description")}
                ></textarea>
            </div>
            <div className="flex items-center justify-between">
                <input
                    className="m-auto bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                    type="submit" value="Add" />
            </div>
        </form>
    </Modal>
}