import {useNavigate, useParams} from "react-router";
import {useEffect} from "react";
import {addToast, Button, Chip, Code} from "@heroui/react";
import {useDispatch, useSelector} from "react-redux";
import {getSnippetById, setLoading} from "../features/snippet/snippetSlice.js";
import CodeEditor from "../components/CodeEditor.jsx";

export default function Snippet() {
    const dispatch = useDispatch();
    const navigate = useNavigate();

    let {id} = useParams();

    useEffect(() => {
        const fetchData = async () => {
            let snippetId = parseInt(id);
            dispatch(setLoading(true));
            let result = await dispatch(getSnippetById(snippetId));
            dispatch(setLoading(false));
            if (!result) {
                addToast({
                    title: "Unable to get snippet",
                    icon: "danger",
                    timeout: 3000,
                    color: "danger",
                    shouldShowTimeoutProgress: true,
                });
            }
        }

        fetchData();

    }, [id]);

    const loading = useSelector((state) => state.snippet.loading);
    const current = useSelector((state) => state.snippet.current);

    let getTagComponents = (tags) => {
        if (tags == null) return null;
        if (tags.length === 0) return null;

        let tagList = tags.split(",")

        return (
            <>
                {tagList.map(tag => (
                    <Chip key={tag} className='m-1 text-xs' color='secondary'>
                        {tag}
                    </Chip>
                ))}
            </>
        )
    }

    const emptyContent = (
        <div>
            <p>Snippet not found</p>
        </div>
    );

    const snippetContent = () => (
        <div>
            <p className='text-md subpixel-antialiased'>{current?.Id}</p>
            <p className='text-md subpixel-antialiased'>{current?.Title}</p>
            <p className='text-md subpixel-antialiased'>{current?.Description}</p>
            <p className='text-md subpixel-antialiased'>
                <Code size='sm' className='text-xs' color='primary'>
                    {current?.Language}
                </Code>
            </p>
            <p className='text-lg font-bold subpixel-antialiased'>
                {getTagComponents(current?.Tags)}
            </p>
            <p className='text-md subpixel-antialiased'>{new Date(current?.CreatedAt).toLocaleString()}</p>
            <p className='text-md subpixel-antialiased'>{current?.UpdatedAt != null ? new Date(current?.UpdatedAt).toLocaleString() : ''}</p>

            <CodeEditor
                loading={loading}
                snippetCode={current?.Code}
                snippetLanguage={current?.Language}
            />
        </div>
    );

    return (
        <div>
            <div className='w-full justify-start items-center'>
                <Button size="sm"
                        color="primary"
                        className="ml-2"
                        onPress={async () => navigate('/')}>Back
                </Button>
            </div>

            {current != null ? snippetContent() : emptyContent}
        </div>
    )
}