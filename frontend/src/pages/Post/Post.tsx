import { useEffect, useState } from "react"
import axiosInstance from "../../api/axios"
import { useParams } from "react-router";
import { TiptapRenderer } from "../../components/PostRender/Render";
import { TiptapRendererComment } from "../../components/CommentRender/CommentRender";
import WriteComment from "../../components/WritePost/WritePost";

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

type Comment = {
  id: number
  content: TiptapDocument
  username: string
  profile_image: string
  created_at: string
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
    
    const [comments, setComments] = useState<Comment[]>([{
        id: 0,
        content: { type: 'doc', content: [] },
        username: '',
        profile_image: '',
        created_at: '',
    }])

    const [reload, setReload] = useState(false)
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

        const getComments = async () =>{
            try{
                const response = await axiosInstance.get(`/devzone-api/v1/posts/comment/${params.postId}`)
                setComments(response.data.comments)
            }catch(error){
                console.log(error)
            }
        }

        getComments()
        getPost()
    },[reload])

    return(
        <>
            <h1 className="text-2xl font-bold">{post?.title}</h1>
            <div className="gap-5 mt-[3rm]">
                <TiptapRenderer content={post?.content} categories={post.categoriesdata} author={post?.username} profileImage={post?.profile_image} DatePublished={post.created_at} ></TiptapRenderer>
            </div>
            <h1 className="text-2xl font-bold mt-20 mb-10">Comments</h1>
            <WriteComment PostId={params.postId} setReload={setReload} reload={reload}/>
            <div className="gap-5 mt-[3rm]">
                {comments.length > 0 ? (comments.map((comment) => (
                    <TiptapRendererComment
                        key={comment.id}
                        content={comment?.content}
                        author={comment?.username}
                        profileImage={comment?.profile_image}
                        DatePublished={comment?.created_at}
                    />
                ))) : ("The post don't have comments")}
            </div>
        </>
    )
}
export default PostView