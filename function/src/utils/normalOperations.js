import {OP_TYPE,FRAME_MESSAGE_TYPE,DATA_TYPE} from './constant';

const MENU_QUERY_URL="/definition/getUserMenus";

const opUpdateMenu={
    type:OP_TYPE.UPDATE_FRAME_DATA,
    params:{
        dataType:DATA_TYPE.QUERY_RESULT
    }
}

const opQueryMenu={
    type:OP_TYPE.REQUEST,
    params:{
        url:MENU_QUERY_URL,
        method:"post"
    },
    input:{},
    description:{key:'gwprj.function.getUserMenu',default:'获取用户菜单'}
}

export function createGetMenuMessage(frameParams){
    opUpdateMenu.params={...opUpdateMenu.params,...frameParams};
    opQueryMenu.successOperation=opUpdateMenu;
    return {
        type:FRAME_MESSAGE_TYPE.DO_OPERATION,
        data:{
            operationItem:opQueryMenu
        }
    };
}

export function createOpenMenuMessage(frameParams,operation){
    return {
        type:FRAME_MESSAGE_TYPE.DO_OPERATION,
        data:{
            operationItem:operation
        }
    };
}