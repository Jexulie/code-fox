import {
    addToast,
    Button, Chip, Code,
    Input,
    Table,
    TableBody,
    TableCell,
    TableColumn,
    TableHeader,
    TableRow,
    useDisclosure
} from "@heroui/react";
import Loading from "../components/Loading";
import {useDispatch, useSelector} from "react-redux";
import {useEffect, useState} from "react";
import {createSnippet, deleteSnippet, getSnippets, setLoading} from "../features/snippet/snippetSlice";
import NewSnippetModal from "../components/NewSnippetModal";
import {useNavigate} from "react-router";

export default function Snippets() {
    const dispatch = useDispatch();
    const navigate = useNavigate();

    useEffect(() => {
        const fetchData = async () => {
            dispatch(setLoading(true));
            let result = await dispatch(getSnippets());
            dispatch(setLoading(false));
            if (!result) {
                addToast({
                    title: "Unable to get snippets",
                    icon: "danger",
                    timeout: 3000,
                    color: "danger",
                    shouldShowTimeoutProgress: true,
                });
            }
        }

        fetchData();
    }, []);


    const {isOpen, onOpen, onOpenChange} = useDisclosure();

    const loading = useSelector((state) => state.snippet.loading);
    const snippets = useSelector((state) => state.snippet.snippets);

    const [search, setSearch] = useState('');

    const handleResetSearch = async () => {
        setSearch('');
        dispatch(setLoading(true));
        let result = await dispatch(getSnippets());
        dispatch(setLoading(false));
        if (!result) {
            addToast({
                title: "Unable to get snippets",
                icon: "danger",
                timeout: 3000,
                color: "danger",
                shouldShowTimeoutProgress: true,
            });
        }
    }

    const handleSearch = async (search) => {
        setSearch(search);
        dispatch(setLoading(true));
        let result = await dispatch(getSnippets(search));
        dispatch(setLoading(false));
        if (!result) {
            addToast({
                title: "Unable to get passwords",
                icon: "danger",
                timeout: 3000,
                color: "danger",
                shouldShowTimeoutProgress: true,
            });
        }
    }

    const addSnippet = async (snippet) => {
        dispatch(setLoading(true));
        let result = await dispatch(createSnippet(snippet));
        if (result) {
            addToast({
                title: "Snippet added!",
                icon: "success",
                timeout: 3000,
                color: "success",
                shouldShowTimeoutProgress: true,
            });
            dispatch(setLoading(false));
            return;
        }

        addToast({
            title: "Could not create snippet",
            icon: "danger",
            timeout: 3000,
            color: "danger",
            shouldShowTimeoutProgress: true,
        });
        dispatch(setLoading(false));
    }

    const removeSnippet = async (id) => {
        dispatch(setLoading(true));
        let result = await dispatch(deleteSnippet(id));
        if (result) {
            addToast({
                title: "Snippet deleted!",
                icon: "success",
                timeout: 3000,
                color: "success",
                shouldShowTimeoutProgress: true,
            });
            dispatch(setLoading(false));
            return;
        }

        addToast({
            title: "Could not delete snippet",
            icon: "danger",
            timeout: 3000,
            color: "danger",
            shouldShowTimeoutProgress: true,
        });
        dispatch(setLoading(false));
    }

    let getTagComponents = (tags) => {
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

    let emptyContent = (
        <TableRow>
            <TableCell className="" colSpan={8}>
                <p className='text-center mt-5 mb-5 text-gray-500'>Nothing found</p>
            </TableCell>
        </TableRow>
    );

    let loadingContent = (
        <div className="flex w-full justify-center items-center mt-50">
            <Loading size='lg' variant='wave'/>
        </div>
    );

    let tableContent = (
        <Table>
            <TableHeader>
                <TableColumn>ID</TableColumn>
                <TableColumn>Title</TableColumn>
                <TableColumn>Language</TableColumn>
                <TableColumn>Description</TableColumn>
                <TableColumn>Tags</TableColumn>
                <TableColumn>Created At</TableColumn>
                <TableColumn>Updated At</TableColumn>
                <TableColumn>Actions</TableColumn>
            </TableHeader>

            <TableBody>
                {snippets.length === 0 ? emptyContent : (snippets.map((item) => (
                    <TableRow key={item.Id}>

                        <TableCell>
                            {item.Id}
                        </TableCell>

                        <TableCell>
                            {item.Title}
                        </TableCell>

                        <TableCell>
                            <Code size='sm' color='primary' className='text-xs'>
                                {item.Language}
                            </Code>
                        </TableCell>

                        <TableCell>
                            {item.Description}
                        </TableCell>

                        <TableCell>
                            {getTagComponents(item.Tags)}
                        </TableCell>

                        <TableCell>
                            {new Date(item.CreatedAt).toLocaleString()}
                        </TableCell>

                        <TableCell>
                            {item.UpdatedAt != null ? new Date(item.UpdatedAt).toLocaleString() : ''}
                        </TableCell>

                        <TableCell className="flex w-full justify-end">
                            <Button size="sm"
                                    color="primary"
                                    className="ml-2"
                                    onPress={async () => navigate(`/snippet/${item.Id}`)}>Details
                            </Button>
                            <Button size="sm"
                                    color="danger"
                                    className="ml-2"
                                    onPress={async () => await removeSnippet(item.Id)}>Delete
                            </Button>
                        </TableCell>
                    </TableRow>
                )))}
            </TableBody>
        </Table>
    );

    return (
        <div>
            <div className="flex w-full justify-end pb-5">
                <Button
                    size="sm"
                    color="primary"
                    onPress={onOpen}>
                    New
                </Button>
            </div>

            <div className="flex">
                <Input
                    isClearable
                    size="sm"
                    color="default"
                    className="mb-5"
                    label="🔍 Search"
                    value={search}
                    onValueChange={handleSearch}
                    onClear={handleResetSearch}/>
            </div>
            {loading ? loadingContent : tableContent}

            <NewSnippetModal isOpen={isOpen} onOpenChange={onOpenChange} onAdd={addSnippet}/>
        </div>
    )
}