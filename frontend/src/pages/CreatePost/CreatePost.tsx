import { useRef, useState } from 'react'
import SimpleEditor from '../../../@/components/tiptap-templates/simple/simple-editor'
import { PlusOutlined } from '@ant-design/icons';
import './style.css'
import { Button, Input, Select, type SelectProps } from 'antd'
const CreatePost = () =>{
    const editorRef = useRef<any>(null)
    const options: SelectProps['options'] = [];
    const [json, setPostJson] = useState()
    const [categories, setCategories] = useState<string | string[]>(null)

    for (let i = 10; i < 36; i++) {
        options.push({
            value: i.toString(36) + i,
            label: i.toString(36) + i,
        });
    }

    const handleChange = (value: string | string[]) => {
        setCategories(value)
    };

    const getJSON = () =>{
        if (editorRef.current) {
            const json = editorRef.current.getJSON()
            setPostJson(json)
        }
    }

    const createPost = async () =>{
        getJSON()
    }
    
    return(
        <>
            <div className='sel-inp flex gap-5 justify-between mb-10'>
                <Input id='Title-input' className='!w-[50%] !bg-[#1D1D1D] !text-white !placeholder-[#9e9e9e]' placeholder="Title" />
                <Select className='!w-[20%] !bg-[#1D1D1D] !text-white !placeholder-[#9e9e9e]'
                    mode="multiple"
                    placeholder="Please select"
                    onChange={handleChange}
                    style={{ width: '100%' }}
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