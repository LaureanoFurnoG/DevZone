import { useEffect, useState } from "react"
import PostCard from "../../components/PostCard/Card"
import axiosInstance from "../../api/axios"
import { useParams } from "react-router";

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
    const [post, setPosts] = useState<Post>()
    const params = useParams();

    useEffect(() =>{
        const getPost = async () =>{
            try{
                const response = await axiosInstance.get(`/devzone-api/v1/posts/${params.postId}`)
                setPosts(response.data.posts)
                console.log(response.data.posts)
            }catch(error){
                console.log(error)
            }
        }

        getPost()
    },[])
    return(
        <>
            <h1 className="text-2xl font-bold">Last Posts</h1>
            <div className="grid gap-5 mt-5">
            {post?.title}
            </div>
        </>
    )
}
export default PostView