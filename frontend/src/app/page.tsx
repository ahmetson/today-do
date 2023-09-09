import Image from 'next/image'

type Task = {
  number: number;
  title: string;
  description: string;
}

async function List() {
  const url = "http://127.0.0.1:2626"
  const params = {
    command: "list",
    parameters: {}
  }
  // Wait for the playlists
  const response = await fetch(url, {
    method: "POST",
    body: JSON.stringify(params),
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    next: {
      revalidate: 0
    }
  });
  const reply = await response.json();

  let tasks: Array<Task>;

  if (reply.status === 'OK') {
    tasks = reply.parameters.list as Array<Task>;
  } else {
    console.error(`replied from server: ${reply.message}`);
    tasks = [];
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
                  <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">Done</button>
             </li>
         ))}
       </ul>
   )
}

export default async function Home() {
  const list = await List();

  return (
      <div>
        <header className="fixed left-0 top-0 flex flex-col w-full border-b border-gray-300 bg-gradient-to-b from-zinc-200 pb-8 pt-4 pl-12 pr-12 backdrop-blur-2xl dark:border-neutral-800 dark:bg-zinc-800/30 dark:from-inherit lg:static lg:w-auto  lg:rounded-xl lg:border lg:bg-gray-200 lg:dark:bg-zinc-800/30">
          <h1>My Short task list <strong>Today</strong> to <strong>Do</strong></h1>
          <sub>Test the SDS Framework <a href="https://github.com/ahmetson/service-lib/tree/service-v1">service-lib@service-v1</a></sub>
          <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">New</button>
        </header>
        <main className="flex min-h-screen flex-col items-center justify-between p-12 pt-24">
          {list}
        </main>
        <footer className="z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex">
          <div className="fixed bottom-0 left-0 flex h-48 w-full items-end justify-center bg-gradient-to-t from-white via-white dark:from-black dark:via-black lg:static lg:h-auto lg:w-auto lg:bg-none">
            <a
                className="pointer-events-none flex place-items-center gap-2 p-8 lg:pointer-events-auto lg:p-0"
                href="https://github.com/ahmetson"
                target="_blank"
                rel="noopener noreferrer"
            >
              By Medet Ahmetson
            </a>
          </div>
        </footer>
      </div>
  )
}
