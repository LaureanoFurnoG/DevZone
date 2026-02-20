import CategoriesTags from "../CategoriesTags/CategoriesTags"
import moment from 'moment';

type Props = {
    Id: number,
    Title: string,
    Text: string,
    UserName: string,
    DateP: string,
    Categories: {id: number, name: string}[], //array
    ImageProfile: string
}
const PostCard = ({Id, Title, Text, Categories, ImageProfile, UserName, DateP}: Props) =>{
    const TruncateText = (text: string) =>{
        return text.slice(0, 500) + '...'
    }

    const DateAgo = (DatePosted: string) =>{
        const datePosted = moment(DatePosted).fromNow();
        
        return datePosted
    }
    return(
        <>
            <div className="bg-[#1D1D1D] p-10 cursor-pointer rounded-[7px]" id={Id.toString()}>
                <div className="flex justify-between">
                    <div className="flex gap-5"> 
                        <img className="border-white border-1 w-[50px] h-[50px] rounded-[100%]" src={ImageProfile} alt="profileImage" />
                        <div>
                            <h2 className="text-xl font-bold">{UserName}</h2>
                            <p>Posted {DateAgo(DateP)}</p>
                        </div>
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
                <div className="mt-10 ml-18">
                    <h1 className="text-2xl font-bold">{Title}</h1>
                    <p className="mt-3">{TruncateText(Text)}</p>
                    <div className="flex w-[100%] gap-5 mt-10">
                       {Categories.map((element) => (
                        <CategoriesTags Title={element.name} />
                       ))}
                    </div>
                </div>

            </div>
        </>
    )
}

export default PostCard