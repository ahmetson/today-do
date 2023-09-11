import type { AppProps } from 'next/app'
import { ButtonProvider } from './Button'


export default function MyApp({ Component, pageProps }: AppProps) {
    alert(`app loaded!`);
    return (
        <ButtonProvider>
            <Component {...pageProps} />
        </ButtonProvider>
    )
}