import { useEffect, useRef, useState } from 'react'
import SimpleEditor from '../../../@/components/tiptap-templates/simple/simple-editor'
import { PlusOutlined } from '@ant-design/icons';
import './style.css'
import { Button, Input, Select, type SelectProps } from 'antd'
import axiosInstance from '../../api/axios';
import { useAuth } from '../../Auth/useAuth';
import { useAppNotification } from '../../components/Notification/Notification';
import { useNavigate } from 'react-router-dom';
const CreatePost = () =>{
    const {me} = useAuth()
    const editorRef = useRef<any>(null)
    const [categories, setCategories] = useState<{ id: string; name: string }[]>([])
    const [selectCategories, setSelectCategories] = useState<string | string[]>()
    const { notify, contextHolder } = useAppNotification()
    const navigate = useNavigate()

    const handleChange = (value: string | string[]) => {
        setSelectCategories(value)
    };

    const homeNavigate = () =>{
        navigate('/home')
    }
    
    const createPost = async () => {
        try{
            if (!editorRef.current) return

            const titleInput = document.getElementById('Title-input') as HTMLInputElement
            const title = titleInput?.value

            if (!title) {
                notify("Missing the title", "Add a title before posting", "error")
                return
            }
            
            if (selectCategories?.length == 0 || !selectCategories){
                notify("Missing the categories", "Add the categories before posting", "error")
                return
            }

            const content = editorRef.current.getJSON()

            const values = {
                Id_user: me?.sub,
                Title: title,
                Categories: selectCategories,
                Content: content,
            }

            await axiosInstance.post("/devzone-api/v1/posts", values)
            notify("Post created", "Post created successfully", "success")
            setTimeout(() =>{
                homeNavigate()
            },3000)
        }catch(error){
            console.log(error)
        }
    }


    const getCategories = async () =>{
        try{
            const response = await axiosInstance.get("/devzone-api/v1/categories")
            setCategories(response.data.categories)
        }catch(error){
            console.log(error)
        }
    }

    useEffect(() =>{
        getCategories()
    },[])
    
    const options: SelectProps['options'] = categories.map(cat => ({
        value: cat.id,
        label: cat.name,
    }))

    return(
        <>    {contextHolder}

            <div className='sel-inp flex gap-5 justify-between mb-10'>
                <Input id='Title-input' className='!w-[50%] !bg-[#1D1D1D] !text-white !placeholder-[#9e9e9e]' placeholder="Title" />
                <Select className='!w-[40%] !bg-[#1D1D1D] !text-white !placeholder-[#9e9e9e]'
                    mode="multiple"
                    placeholder="Please select"
                    onChange={handleChange}
                    options={options}
                />
            </div>
            <SimpleEditor ref={editorRef} />
            <Button
                type="primary"
                icon={<PlusOutlined />}
                className={`
                !bg-[rgba(0,102,197,1)]
                !h-12
                !mt-10`}
                onClick={() => createPost()}
            >
                Create Post
            </Button>
        </>
    )
}
export default CreatePost