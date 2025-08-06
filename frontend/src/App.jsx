import './App.css';
import {HeroUIProvider} from "@heroui/react";
import AppRoutes from "./routes";
import {store} from './store/store';
import {Provider} from "react-redux";
import {HashRouter} from "react-router";
import {ToastProvider} from "@heroui/react";

function App() {
    return (
        <Provider store={store}>
            <HeroUIProvider>
                <main className="light text-foreground bg-background">
                <ToastProvider />
                <HashRouter>
                    <AppRoutes/>
                </HashRouter>
                    </main>
            </HeroUIProvider>
        </Provider>
    )
}

export default App
