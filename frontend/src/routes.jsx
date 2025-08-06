import { Routes, Route } from "react-router";

import Main from "./pages/Main";

import Snippets from "./pages/Snippets";
import Tags from "./pages/Tags";
import Snippet from "./pages/Snippet";

export default function AppRoutes() {
    return (
        <Routes>
            <Route path="/" element={<Main />}>
                <Route index element={<Snippets />} />
                <Route path="snippet/:id" element={<Snippet />} />
                <Route path="tags" element={<Tags />} />
            </Route>
        </Routes>
    );
}