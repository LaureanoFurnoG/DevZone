import CategoriesTags from "../CategoriesTags/CategoriesTags"
import moment from 'moment';
import { DeleteOutlined } from '@ant-design/icons';
import axiosInstance from "../../api/axios";

type Props = {
    Id: number,
    Title: string,
    Text: string,
    UserName: string,
    DateP: string,
    Categories: {id: number, name: string}[], //array
    ImageProfile: string
    AuthorID: string
    userSessionID?: string
    onDelete?: (id: number) => void
}
const PostCard = ({Id, Title, AuthorID, userSessionID, Text, Categories, ImageProfile, UserName, DateP, onDelete}: Props) =>{
    const TruncateText = (text: string) =>{
        return text.slice(0, 500) + '...'
    }
    //in the future i will do a custom hook with this
    const DateAgo = (DatePosted: string) =>{
        const datePosted = moment(DatePosted).fromNow();
        
        return datePosted
    }

    const DeletePost = async (e: React.MouseEvent) =>{
        e.preventDefault()
        try{
            await axiosInstance.delete(`/devzone-api/v1/posts/${Id}`)
            onDelete?.(Id)
        }catch(err: any){
            console.log(err)
        }
    }
    const stringToColor = (str: string) => {
        let hash = 0;
        for (let i = 0; i < str.length; i++) {
            hash = str.charCodeAt(i) + ((hash << 5) - hash);
        }
        return `#${((hash >> 16) & 0xff).toString(16).padStart(2,'0')}${((hash >> 8) & 0xff).toString(16).padStart(2,'0')}${(hash & 0xff).toString(16).padStart(2,'0')}`;
    };
    return(
        <>
            <div className="bg-[#1D1D1D] p-10 cursor-pointer rounded-[7px]" id={Id.toString()}>
                <div className="flex justify-between">
                    <div className="flex gap-5 w-[100%] !justify-between"> 
                        {ImageProfile ? <img className="border-white border-1 w-[50px] h-[50px] rounded-[100%]" src={ImageProfile} alt="profileImage" /> : <p 
                            style={{ backgroundColor: stringToColor(UserName ?? "?") }} 
                            className="border-white p-4 border-1 w-[50px] h-[50px] rounded-[100%] text-white flex justify-center items-center"
                        >
                            {UserName?.[0] ?? ""}
                        </p>}
                        <div className="w-full">
                            <h2 className="text-xl font-bold">{UserName}</h2>
                            <p>Posted {DateAgo(DateP)}</p>
                        </div>
                        {(AuthorID === userSessionID) && <DeleteOutlined onClick={(e) => DeletePost(e)} className="text-2xl !text-[red] relative z-20 hover:bg-[white] pl-6 pr-6 rounded"/>}
                    </div>
                    {
                        /* in the future maybe implement this, but with a conditional if you are the creator... maybe with the userID in me.
                        <div className="flex gap-5 items-center justify-center">
                            <DeleteOutlined className="text-[23px] !text-red-600 cursor-pointer"/>   
                            <EditOutlined   className="text-[23px] !text-blue-600 cursor-pointer"/>
                        </div>   
                        */
                    }
                </div>
                <div className="mt-10 sm:ml-18">
                    <h1 className="text-2xl font-bold">{Title}</h1>
                    <p className="mt-3">{TruncateText(Text)}</p>
                    <div className="flex w-[100%] gap-5 mt-10 overflow-auto">
                       {Categories.map((element) => (
                        <CategoriesTags key={element.id} Title={element.name} />
                       ))}
                    </div>
                </div>

            </div>
        </>
    )
}

export default PostCard