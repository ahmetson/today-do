import Modal from './Modal'
import Header from './Header';
import { Task, FetchList } from '@/data/task';
import List from './List';

export default async function Home() {
    const tasks = await FetchList();

  return (
      <div>
        <Header></Header>
        <main className="flex min-h-screen flex-col items-center justify-between p-12 pt-24">
          <List tasks={tasks}></List>
          <Modal></Modal>
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
