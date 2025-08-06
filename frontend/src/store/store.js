import { configureStore } from '@reduxjs/toolkit';
import snippetReducer from "../features/snippet/snippetSlice";
import tagReducer from "../features/tag/tagSlice";

export const store = configureStore({
    reducer: {
        snippet: snippetReducer,
        tag: tagReducer,
    },
});