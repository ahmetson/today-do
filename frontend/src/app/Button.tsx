"use client";

import {
    useReducer,
    useContext,
    createContext,
    ReactNode,
    Dispatch,
} from 'react'

type ButtonState = boolean
type ButtonAction =
    | {
    type: 'TRUE' | 'FALSE'
}
    | {
    type: 'TOGGLE'
    payload: boolean
}

const ButtonStateContext = createContext<ButtonState>(true)
const ButtonDispatchContext = createContext<Dispatch<ButtonAction>>(
    () => null
)

const reducer = (state: ButtonState, action: ButtonAction) => {
    console.log(`Called a button reducer`);
    switch (action.type) {
        case 'TRUE':
            return true
        case 'FALSE':
            return false
        case 'TOGGLE':
            return !state
        default:
            throw new Error(`Unknown action: ${JSON.stringify(action)}`)
    }
}

type ButtonProviderProps = {
    children: ReactNode
    initialValue?: boolean
}

export const ButtonProvider = ({
                                    children,
                                    initialValue = true,
                                }: ButtonProviderProps) => {
    const [state, dispatch] = useReducer(reducer, initialValue)
    return (
        <ButtonDispatchContext.Provider value={dispatch}>
            <ButtonStateContext.Provider value={state}>
                {children}
            </ButtonStateContext.Provider>
        </ButtonDispatchContext.Provider>
    )
}

export const useButton = () => useContext(ButtonStateContext)
export const useDispatchButton = () => useContext(ButtonDispatchContext)