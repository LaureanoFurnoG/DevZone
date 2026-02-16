import { useEffect, useState } from "react"
import PostCard from "../../components/PostCard/Card"
import axiosInstance from "../../api/axios"

type Category = {
  id: number
  name: string
}

type Post = {
  id: number
  title: string
  content: any
  created_at: string
  categoriesdata: Category[]
}

const Home = () =>{
    const [posts, setPosts] = useState<Post[]>([])

    useEffect(() =>{
        const getPosts = async () =>{
            try{
                const response = await axiosInstance.get('/devzone-api/v1/posts')
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
                        ?.map((c: any) => c.text)
                        .join("") ?? ""

                    return (
                    <PostCard
                        Id={post.id}
                        Title={post.title}
                        Text={preview}
                        UserName={"UserName"}
                        DateP={post.created_at}
                        Categories={post.categoriesdata}
                        ImageProfile=""
                    />
                    )
                })}
            </div>
        </>
    )
}
export default Home