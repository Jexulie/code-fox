import {createSlice} from '@reduxjs/toolkit';
import {
    CreateSnippet,
    DeleteSnippet,
    GetAllSnippets,
    GetSnippetById,
    UpdateSnippet
} from "../../../wailsjs/go/app/SnippetManager";

const initialState = {
    loading: false,
    snippets: [],
    current: null,
};

export const getSnippetById = (id) => async (dispatch, getState) => {
    try {

        let response = await GetSnippetById(id);
        if(response != null){
            dispatch(setCurrent(response));
        }

        return true;
    }catch(err){
        console.log(err);
        return false;
    }
}

export const getSnippets = (tag="") => async (dispatch, getState) => {
    try {

        let response = await GetAllSnippets(tag);
        if(response != null){
            dispatch(setSnippets(response));
        }

        return true;
    }catch(err){
        console.log(err);
        return false;
    }
}

export const createSnippet = (snippet) => async (dispatch, getState) => {
    try {

        let response = await CreateSnippet(snippet);
        if(response != null){
            await dispatch(getSnippets());
        }

        return true;
    }catch(err){
        console.log(err);
        return false;
    }
}

export const updateSnippet = (snippet) => async (dispatch, getState) => {
    try {

        let response = await UpdateSnippet(snippet);
        if(response != null){
            await dispatch(getSnippets());
        }

        return true;
    }catch(err){
        console.log(err);
        return false;
    }
}

export const deleteSnippet = (id) => async (dispatch, getState) => {
    try {

        let response = await DeleteSnippet(id);
        if(response != null){
            await dispatch(getSnippets());
        }

        return true;
    }catch(err){
        console.log(err);
        return false;
    }
}

export const snippetSlice = createSlice({
    name: 'snippet',
    initialState,
    reducers: {
        setLoading: (state, action) => {
            state.loading = action.payload;
        },
        setSnippets: (state, action) => {
            state.snippets = action.payload;
        },
        setCurrent: (state, action) => {
            state.current = action.payload;
        }
    },
});

export const {
    setLoading,
    setSnippets,
    setCurrent

} = snippetSlice.actions;
export default snippetSlice.reducer;