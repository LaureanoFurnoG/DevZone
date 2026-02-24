import { useEffect, useState } from "react"
import PostCard from "../../components/PostCard/Card"
import axiosInstance from "../../api/axios"
import { Link } from "react-router-dom"
import { useAuth } from "../../Auth/useAuth"

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
  id_user: string
  username: string
  profile_image: string
  created_at: string
  categoriesdata: Category[]
}

const ReleaseNotes = () =>{
    const [posts, setPosts] = useState<Post[]>([])
    const {me} = useAuth()
    useEffect(() =>{
        const getPosts = async () =>{
            try{
                const response = await axiosInstance.get(`/devzone-api/v1/posts/${4}`)
                setPosts(response.data.posts)
            }catch(error){
                console.log(error)
            }
        }

        getPosts()
    },[])

    const handleDelete = (id: number) => {
        setPosts(prev => prev.filter(post => post.id !== id))
    }

    return(
        <>
            <h1 className="text-2xl font-bold">Last Posts</h1>
            <div className="flex flex-col gap-5 mt-5">
                {posts.map((post) => {
                    const preview =
                    post.content?.content?.[0]?.content
                        ?.map((c: TiptapNode) => c.text)
                        .join("") ?? ""

                    return (
                        <Link to={`/post/${post.id}`}  className="block !text-white"> 
                            <PostCard
                                key={post.id}
                                Id={post.id}
                                AuthorID={post.id_user}
                                Title={post.title}
                                Text={preview}
                                UserName={post.username}
                                DateP={post.created_at}
                                Categories={post.categoriesdata}
                                ImageProfile={post.profile_image} 
                                userSessionID={me?.sub}
                                onDelete={handleDelete}
                            />
                        </Link>
                    )
                })}
            </div>
        </>
    )
}
export default ReleaseNotes