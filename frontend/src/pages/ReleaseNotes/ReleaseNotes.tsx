import { useAuth } from "../../Auth/useAuth"

const CreatePost = () =>{
    const { me } = useAuth()
    return(
        <>
            <p>{me?.name}</p>
        </>
    )
}
export default CreatePost