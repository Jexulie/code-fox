import Navbar from "../components/Navbar";
import Sidebar from "../components/Sidebar";

import { Outlet } from "react-router";

export default function Main() {
    return (
        <div className="flex h-screen bg-gray-50">
            <Sidebar />
            <div className="flex-1 overflow-auto">
                <Navbar />
                <div className="p-6">
                    <Outlet /> {/* Renders nested routes */}
                </div>
            </div>
        </div>
    );
}