type Props = {
    Title: string
}
const CategoriesTags = ({Title}: Props) =>{
    const color = (() => {
        switch (Title) {
        case "Auth":
            return "#5073FF"
        case "React":
            return "#50CAFF"
        default:
            return "#999999"
        }
    })()
    return(
        <>
            <div style={{ backgroundColor: color }} className="text-white border-white border-1 w-30 rounded-[60px]">
                <p className="text-center p-2">{Title}</p>
            </div>
        </>
    )
}

export default CategoriesTags