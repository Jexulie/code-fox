import {createSlice} from '@reduxjs/toolkit';
import {CreateTag, DeleteTag, GetAllTags, UpdateTag} from "../../../wailsjs/go/app/TagManager";

const initialState = {
    loading: false,
    tags: [],
};

export const getTags = (tag="") => async (dispatch, getState) => {
    try {

        let response = await GetAllTags(tag);
        if(response != null){
            dispatch(setTags(response));
        }

        return true;
    }catch(err){
        console.log(err);
        return false;
    }
}

export const createTag = (tag) => async (dispatch, getState) => {
    try {

        let response = await CreateTag(tag);
        if(response != null){
            await dispatch(getTags());
        }

        return true;
    }catch(err){
        console.log(err);
        return false;
    }
}

export const updateTag = (tag) => async (dispatch, getState) => {
    try {

        let response = await UpdateTag(tag);
        if(response != null){
            await dispatch(getTags());
        }

        return true;
    }catch(err){
        console.log(err);
        return false;
    }
}

export const deleteTag = (id) => async (dispatch, getState) => {
    try {

        let response = await DeleteTag(id);
        if(response != null){
            await dispatch(getTags());
        }

        return true;
    }catch(err){
        console.log(err);
        return false;
    }
}

export const tagSlice = createSlice({
    name: 'tag',
    initialState,
    reducers: {
        setLoading: (state, action) => {
            state.loading = action.payload;
        },
        setTags: (state, action) => {
            state.tags = action.payload;
        }
    },
});

export const {
    setLoading,
    setTags

} = tagSlice.actions;
export default tagSlice.reducer;