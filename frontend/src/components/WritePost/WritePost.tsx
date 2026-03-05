import { useRef, useState } from 'react'
import SimpleEditor from '../../../@/components/tiptap-templates/simple/simple-editor'
import { PlusOutlined } from '@ant-design/icons';
import './style.css'
import { Button } from 'antd'
import axiosInstance from '../../api/axios';
import { useAuth } from '../../Auth/useAuth';
import { useAppNotification } from '../../components/Notification/Notification';

type Props = {
    PostId: string | undefined
    setReload: (value: boolean) => void
    reload: boolean

}

const WriteComment = ({ PostId, setReload, reload }: Props) =>{
    const {me} = useAuth()
    const editorRef = useRef<any>(null)
    const { notify, contextHolder } = useAppNotification()
    const [loading, setLoading] = useState(false)

    const writeComment = async () => {
        try{
            if (!editorRef.current) return
            setLoading(true)

            const content = editorRef.current.getJSON()

            const values = {
                Id_user: me?.sub,
                Content: content,
            }
            
            await axiosInstance.post(`/devzone-api/v1/posts/comment/${PostId}`, values)
            notify("Commented created", "Commented successfully", "success")
            setLoading(false)
            setReload(!reload)
            editorRef.current.clearContent(true)
        }catch(error: any){
            setLoading(false)
            console.log(error)
            notify("Unauthorized", error.response.data, "error")
        }finally {
            setLoading(false) 
        }
    }

    return(
        <>    
        <div className='createComment-Container'>
            {contextHolder}
            <SimpleEditor ref={editorRef} />
            <Button
                type="primary"
                icon={<PlusOutlined />}
                className={`
                !bg-[rgba(0,102,197,1)]
                !h-12
                !mt-5 !mb-10`}
                onClick={() => writeComment()}
                loading={loading}
            >
                Create Comment
            </Button>
        </div>
        </>
    )
}
export default WriteComment