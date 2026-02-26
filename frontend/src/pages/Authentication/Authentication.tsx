import PostCard from "../../components/PostCard/Card"
import { Link } from "react-router-dom"
import { useAuth } from "../../Auth/useAuth"
import { usePosts } from "../../customHooks/usePosts/usePosts"

type TiptapNode = {
  type: string
  content?: TiptapNode[]
  text?: string
}

const Authentication = () =>{
    const {posts, setPosts} = usePosts(2)
    const {me} = useAuth()
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
export default Authentication