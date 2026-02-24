import { useEffect, useState } from "react"
import axiosInstance from "../../api/axios"
import { useParams } from "react-router";
import { TiptapRenderer } from "../../components/PostRender/Render";

type TiptapNode = {
  type: string
  content?: TiptapNode[]
  text?: string
}

type TiptapDocument = {
  type: "doc"
  content: TiptapNode[]
}

type Category = {
  id: number
  name: string
}

type Post = {
  id: number
  title: string
  content: TiptapDocument
  username: string
  profile_image: string
  created_at: string
  categoriesdata: Category[]
}

const PostView = () =>{
    const [post, setPost] = useState<Post>({
        id: 0,
        title: '',
        content: { type: 'doc', content: [] },
        username: '',
        profile_image: '',
        created_at: '',
        categoriesdata: [],
    })
    const params = useParams();

    useEffect(() =>{
        const getPost = async () =>{
            try{
                const response = await axiosInstance.get(`/devzone-api/v1/posts/publishedpost/${params.postId}`)
                setPost(response.data.post)
                console.log(response.data)
            }catch(error){
                console.log(error)
            }
        }

        getPost()
    },[])

    return(
        <>
            <h1 className="text-2xl font-bold">{post?.title}</h1>
            <div className="gap-5 mt-[3rm]">
                <TiptapRenderer content={post?.content} categories={post.categoriesdata} author={post?.username} profileImage={post?.profile_image} DatePublished={post.created_at} ></TiptapRenderer>
            </div>
        </>
    )
}
export default PostView