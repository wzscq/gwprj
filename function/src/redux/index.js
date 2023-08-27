import { configureStore } from '@reduxjs/toolkit'

import frameReducer from './frameSlice';
import i18nReducer from './i18nSlice';
import menuReducer from './menuSlice';

export default configureStore({
  reducer: {
    frame:frameReducer,
    i18n:i18nReducer,
    menu:menuReducer
  }
});