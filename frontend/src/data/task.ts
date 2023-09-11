
export type Task = {
    number: number;
    title: string;
    description: string;
}

type Request = {
    command: string;
    parameters: any;
}

type Reply = {
    status: "OK" | "fail";
    message: string;
    parameters: any;
}

export const url = "http://127.0.0.1:2626"

export const FetchList = async function() {
    const params: Request = {
        command: "list",
        parameters: {}
    }

    const reply = await request(params);
    // Wait for the playlists

    let tasks: Array<Task>;

    if (reply.status === 'OK') {
        tasks = reply.parameters.list as Array<Task>;
    } else {
        console.error(`replied from server: ${reply.message}`);
        tasks = [];
    }

    return tasks;
}

export const FetchAdd = async function(task: Task): Promise<Task> {
    const params: Request = {
        command: "add",
        parameters: {
            title: task.title,
            description: task.description,
        }
    }

    console.log(`request`, params);

    const reply = await request(params);
    // Wait for the playlists

    if (reply.status === 'OK') {
        task.number = reply.parameters.number as number;
    } else {
        console.error(`replied from server: ${reply.message}`);
        task.number = 0;
    }

    return task;
}

export const FetchDone = async function(number: number): Promise<boolean> {
    const params: Request = {
        command: "done",
        parameters: {
            number,
        }
    }

    const reply = await request(params);

    if (reply.status === 'OK') {
        return true;
    } else {
        console.error(`replied from server: ${reply.message}`);
    }

    return false;
}


export const request = async function(request: Request): Promise<Reply> {
    // Wait for the playlists
    const response = await fetch(url, {
        method: "POST",
        body: JSON.stringify(request),
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        next: {
            revalidate: 0
        }
    });
    return await response.json();
}

