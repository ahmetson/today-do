import { NextRequest, NextResponse } from 'next/server'
import {Task, FetchAdd, FetchDone} from '../../../data/task';

// POST method is used to revalidate the parameters
export async function POST(request: NextRequest) {
    if (request.body == null) {
        return NextResponse.json({message: 'Missing body'}, {status: 401})
    }
    let task = await request.json() as Task;

    await FetchAdd(task);

    return NextResponse.json(task , { status: 200 })
}

export async function GET(request: NextRequest) {
    const { searchParams } = new URL(request.url)
    const numberStr = searchParams.get('number');
    if (numberStr == null) {
        return NextResponse.json({message: 'Missing number parameter'}, {status: 401})
    }
    const number = parseInt(numberStr);
    if (isNaN(number)) {
        return NextResponse.json({message: 'Invalid number'}, {status: 401})
    }

    let response = await FetchDone(number);

    return NextResponse.json({done: true, number}, { status: 200 })
}