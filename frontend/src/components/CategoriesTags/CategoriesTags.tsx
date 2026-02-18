type Props = {
    Title: string
}
const CategoriesTags = ({Title}: Props) =>{
    const color = (() => {
        switch (Title) {
        case "Auth":
            return "#5073FF"
        case "Framework":
            return "#50CAFF"
        case "Libraries":
            return "#f0354e"
        case "Dependencies":
            return "#7800c9"
        case "Backend":
            return "#009e42"
        case "ReleaseNotes":
            return "#dbba00"
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