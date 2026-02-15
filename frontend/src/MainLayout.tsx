import React, { useEffect, useState } from 'react';
import {
AppstoreOutlined,
BookOutlined,
HomeOutlined,
LoginOutlined,
LogoutOutlined,
MenuFoldOutlined,
MenuUnfoldOutlined,
PlusOutlined,
} from '@ant-design/icons';
import { Button, Layout, Menu, theme, Input } from 'antd';
import type { GetProps } from 'antd';

import logo from './assets/Dev.svg'
import "./App.css"
import { BsBook } from "react-icons/bs";
import { LuLibraryBig } from "react-icons/lu";
import { GoPackageDependencies } from "react-icons/go";
import { FaShieldAlt } from "react-icons/fa";
import { CiServer } from "react-icons/ci";
import { Outlet, useLocation } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';
import { useAuth } from './Auth/useAuth';

const { Header, Sider, Content } = Layout;
type SearchProps = GetProps<typeof Input.Search>;

const MainLayout: React.FC = () => {
    const { Search } = Input;
    const { pathname } = useLocation();
    const [collapsed, setCollapsed] = useState(true);
    const {
        token: { colorBgContainer },
    } = theme.useToken();

    const navigate = useNavigate()
    const {me, login, logout, isAuthenticated} = useAuth()
    const [currentPage, setCurrentPage] = useState<string>('1')
    const createPost = () =>{
        navigate('/createpost')
    }
    const releaseNotes = () =>{
        navigate('/releasenotes')
    }

    const navigateFrameworks = () =>{
        navigate('/frameworks')
    }

    const navigateBackend = () =>{
        navigate('/backend')
    }

    const navigateDependencies = () =>{
        navigate('/dependencies')
    }

    const navigateLibraries = () =>{
        navigate('/libraries')
    }

    const navigateAuthentication = () =>{
        navigate('/authentication')
    }

    const navigateHome = () =>{
        navigate('/home')
    }

    useEffect(() => {
        // eslint-disable-next-line react-hooks/set-state-in-effect
        if (pathname.includes("home")) setCurrentPage('1');
        else if (pathname.includes("releasenotes")) setCurrentPage('3');
        else if (pathname.includes("frameworks")) setCurrentPage('8');
        else if (pathname.includes("libraries")) setCurrentPage('9');
        else if (pathname.includes("dependencies")) setCurrentPage('10');
        else if (pathname.includes("authentication")) setCurrentPage('11');
        else if (pathname.includes("backend")) setCurrentPage('12');
        else if (pathname.includes("createpost")) setCurrentPage('0.1');
    }, [pathname]);


    const onSearch: SearchProps['onSearch'] = (value, _e, info) => console.log(info?.source, value);

    return (
        <Layout className='!h-screen'>
        <Sider className={`!bg-[#1D1D1D] !border-[#ffffff] !border-2 !border-l-0 !border-b-0 !border-t-0 ${collapsed ? "hidden sm:block !border-0" : "!w-full"}`} trigger={null} collapsible collapsed={collapsed} width={250}>
            <Menu className='!bg-[#1D1D1D]'
            theme="dark"
            mode="inline"
            selectedKeys={[currentPage]}
            items={[
                {
                key: '0',
                icon: <img src={logo} className="w-10 rounded-lg" />,
                label: <span className='text-white'>Zone</span>,
                disabled: true,
                },
                {
                key: '1',
                icon: <HomeOutlined />,
                label: 'Home',
                onClick: () => navigateHome()
                },
                {
                key: '2',
                icon: <AppstoreOutlined/>,
                label: 'Categories',
                children: [
                    { key: '8', icon: <BookOutlined/>, label: 'Frameworks', onClick: () => navigateFrameworks() },
                    { key: '9', icon: <LuLibraryBig />, label: 'Libraries', onClick: () => navigateLibraries()},
                    { key: '10', icon: <GoPackageDependencies />, label: 'Dependencies', onClick: () => navigateDependencies() },
                    { key: '11', icon: <FaShieldAlt />, label: 'Authentication', onClick: () => navigateAuthentication()},
                    { key: '12', icon: <CiServer />, label: 'Backend', onClick: () => navigateBackend()},
                ],
                },
                {
                key: '3',
                icon: <BsBook />,
                label: 'Release Notes',
                onClick: () => releaseNotes()
                },
            ]}
            />
            <div>
                <Button
                    type="primary"
                    icon={<PlusOutlined />}
                    className={`
                    !bg-[linear-gradient(155deg,rgba(55,22,71,1)_1%,rgba(0,102,197,1)_50%,rgba(55,22,71,1)_100%)]
                    !border-0
                    !h-12
                    ${collapsed ? '!w-[89%] !mx-1' : '!w-[94%] !mx-1'}
                    `}
                    onClick={() => createPost()}
                >
                    {!collapsed && 'Create Post'}
                </Button>
                {isAuthenticated ? (
                    <div className='flex justify-center items-center p-5 border-t-2 mt-8 border-white w-full gap-3'>
                        <img className='rounded-full w-10 h-10 !bg-red-500' src={me?.lastname} alt="" />
                        {!collapsed && 
                            <div className='!w-[80%]'>
                            <h2 className='text-white text-xl font-semibold'>{me?.name}</h2>
                            <p className='text-gray-500 break-all'>User</p>
                            </div>
                        }
                    </div>
                ) : (
                    <div className='border-t-2 mt-8 border-white'>
                        <Button
                            type="primary"
                            icon={<LoginOutlined />}
                            className={`
                            !bg-[linear-gradient(155deg,rgba(55,22,71,1)_1%,rgba(0,102,197,1)_50%,rgba(55,22,71,1)_100%)]
                            !border-0
                            !h-12 m-8
                            ${collapsed ? '!w-[89%] !mx-1' : '!w-[94%] !mx-1'}
                            `}
                            onClick={() => login()}
                        >
                            {!collapsed && 'Login'}
                        </Button>
                    </div>
                )}
            </div>
        </Sider>
        <Layout className='!bg-[#303030]'>
            <Header className={`!flex !w-full !justify-between !bg-[#1D1D1D] !border !border-[#ffffff] !border-2 !border-t-0 !border-r-0 ${collapsed ? "!border-l-0" : ""}`} style={{ padding: 0, background: colorBgContainer }}>
                <div className='!w-[100%] flex items-center'>
                    <Button
                        className='!text-white'
                        type="text"
                        icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
                        onClick={() => setCollapsed(!collapsed)}
                        style={{
                        fontSize: '16px',
                        width: 64,
                        height: 64,
                        }}
                    />
                    <Search placeholder="input search text" onSearch={onSearch} className={`search-Header  ${collapsed ? "!w-[75%]" : "!hidden sm:!flex"}`} style={{ width: '60%' }} />
                </div>
                {isAuthenticated &&
                    <Button
                        className='!text-white '
                        type="text"
                        icon={<LogoutOutlined />}
                        onClick={() => logout()}
                        style={{
                        fontSize: '16px',
                        width: 64,
                        height: 64,
                        }}
                    />
                }
            </Header>
            <Content    
            style={{
                margin: '24px 16px',
                padding: 24,
                minHeight: 280,
            }}
            >
            <Outlet />
            </Content>
        </Layout>
        </Layout>
    );
};

export default MainLayout;