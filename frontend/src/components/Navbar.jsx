import { useDispatch } from 'react-redux';
import { useNavigate } from 'react-router';

export default function Navbar() {
    const dispatch = useDispatch();
    const navigate = useNavigate();

    return (
        <div className="bg-white shadow-sm p-4 flex justify-between items-center">
            <h3 className="font-semibold text-gray-500">🦊 Code Fox</h3>
        </div>
    );
}