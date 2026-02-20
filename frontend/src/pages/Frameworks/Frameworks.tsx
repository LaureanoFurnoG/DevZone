import { useEffect, useState } from "react"
import PostCard from "../../components/PostCard/Card"
import axiosInstance from "../../api/axios"

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


const Frameworks = () =>{
    const [posts, setPosts] = useState<Post[]>([])

    useEffect(() =>{
        const getPosts = async () =>{
            try{
                const response = await axiosInstance.get(`/devzone-api/v1/posts/${5}`)
                setPosts(response.data.posts)
            }catch(error){
                console.log(error)
            }
        }

        getPosts()
    },[])
    return(
        <>
            <h1 className="text-2xl font-bold">Last Posts</h1>
            <div className="grid gap-5 mt-5">
                {posts.slice(0, 30).map((post) => {
                    const preview =
                    post.content?.content?.[0]?.content
                        ?.map((c: TiptapNode) => c.text)
                        .join("") ?? ""

                    return (
                    <PostCard
                        Id={post.id}
                        Title={post.title}
                        Text={preview}
                        UserName={post.username}
                        DateP={post.created_at}
                        Categories={post.categoriesdata}
                        ImageProfile={post.profile_image}
                    />
                    )
                })}
            </div>
        </>
    )
}
export default Frameworks