import { NavLink } from "react-router";

export default function Sidebar() {
    return (
        <div className="w-64 bg-gray-800 text-white p-4">
            <h2 className="text-xl font-bold mb-6">Menu</h2>
            <nav>
                <NavLink to="/" className="block py-2 px-4 hover:bg-gray-700 rounded">
                    💾 Snippets
                </NavLink>
                <NavLink to="/tags" className="block py-2 px-4 hover:bg-gray-700 rounded">
                    🎫 Tags
                </NavLink>
                {/*<NavLink to="/settings" className="block py-2 px-4 hover:bg-gray-700 rounded">*/}
                {/*    ⚙️ Settings*/}
                {/*</NavLink>*/}
            </nav>
        </div>
    );
}