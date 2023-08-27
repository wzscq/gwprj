import {useCallback, useEffect, useMemo} from 'react';
import {useSelector} from 'react-redux';
import { Collapse,Card,Space } from 'antd';

import useI18n from '../../hook/useI18n';
import useFrame from "../../hook/useFrame";
import {createGetMenuMessage,createOpenMenuMessage} from '../../utils/normalOperations';
import PageLoading from './PageLoading';

import './index.css';

export default function Function(){
  const {locale,getLocaleLabel}=useI18n();
  const {menus,loaded}=useSelector(state=>state.menu);

  const sendMessageToParent=useFrame();
  const {origin,item}=useSelector(state=>state.frame);

  const frameParams=useMemo(()=>{
    if(origin&&item){
        return ({
            frameType:item.frameType,
            frameID:item.params.key,
            origin:origin
        });
    }
    return null;
    },
  [origin,item]);

  const openMenu=useCallback((menu)=>{
    sendMessageToParent(createOpenMenuMessage(frameParams,menu));
  },[frameParams,sendMessageToParent]);

  //加载配置
  useEffect(()=>{
    console.log(frameParams,loaded);
    if(frameParams!==null){
        if(loaded===false){
            console.log("sendMessageToParent",frameParams);
            sendMessageToParent(createGetMenuMessage(frameParams));
        }
    }
  },[loaded,frameParams,sendMessageToParent]);

  if(loaded===false){
    return (<PageLoading/>);
  }

  const defaultActiveKey=[];
  const menuGroups=menus.map(item=>{
    const menuItems=item.children.map(child=>{
      return (
        <Card
          size='small'
          title={child.name}
          extra={<a onClick={()=>openMenu(child)} href="#">打开</a>}
          style={{width: 200,}}
        >
          <div>{child.description}</div>
        </Card>);
    });

    const menuGroup=(
      <div className='function-group' style={{width:"100%",display:'flex'}}>
        <Space size={16}>
          {menuItems}
        </Space>
      </div>
    )

    defaultActiveKey.push(item.id);

    return ({
      key: item.id,
      label:item.name,
      children: menuGroup,
    });
  });

  return <Collapse defaultActiveKey={defaultActiveKey} bordered={false} size='small' items={menuGroups}/>;
}