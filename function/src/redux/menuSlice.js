import { createSlice } from '@reduxjs/toolkit';

// Define the initial state using that type
const initialState = {
    menus:[],
    loaded:false,
    pending:false,
    inlineCollapsed:false,
    errorCode:0
}

export const menuSlice = createSlice({
  name: 'menu',
  initialState,
  reducers: {
    setInlineCollapsed:(state,action) => {
        state.inlineCollapsed=action.payload;
    },
    resetMenu:(state,action)=>{
      state.menus=initialState.menus;
      state.loaded=initialState.loaded;
      state.pending=initialState.pending;
      state.inlineCollapsed=initialState.inlineCollapsed;
      state.errorCode=initialState.errorCode;
    },
    setMenu:(state,action)=>{
      state.menus=action.payload;
      state.errorCode=0;
      state.loaded=true;
    }
  },
  extraReducers: (builder) => {
    
  },
});

// Action creators are generated for each case reducer function
export const {
  setInlineCollapsed,
  resetMenu,
  setMenu
} = menuSlice.actions;

export default menuSlice.reducer;