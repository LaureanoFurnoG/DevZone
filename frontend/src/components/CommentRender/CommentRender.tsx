import React from 'react'
import { generateHTML } from '@tiptap/core'
import StarterKit from '@tiptap/starter-kit'
import '../../../@/components/tiptap-templates/simple/simple-editor.scss'
import "../../../@/components/tiptap-node/blockquote-node/blockquote-node.scss"
import "../../../@/components/tiptap-node/code-block-node/code-block-node.scss"
import "../../../@/components/tiptap-node/horizontal-rule-node/horizontal-rule-node.scss"
import "../../../@/components/tiptap-node/list-node/list-node.scss"
import "../../../@/components/tiptap-node/image-node/image-node.scss"
import "../../../@/components/tiptap-node/heading-node/heading-node.scss"
import "../../../@/components/tiptap-node/paragraph-node/paragraph-node.scss"
import moment from 'moment'

type TiptapNode = {
  type: string
  content?: TiptapNode[]
  text?: string
}

type TiptapDocument = {
  type: "doc"
  content: TiptapNode[]
}

interface TiptapRendererProps {
  content?: TiptapDocument 
  author?: string
  profileImage?: string
  className?: string
  DatePublished: string
}

export const TiptapRendererComment: React.FC<TiptapRendererProps> = ({
  content,
  author,
  profileImage,
  DatePublished,
  className = 'simple-editor-content'
}) => {

  const html = React.useMemo(() => {
    if (!content) return ''

    try {
      return generateHTML(content, [StarterKit])
    } catch{
      return ''
    }
  }, [content])

  const DateAgo = (DatePosted: string) =>{
      const datePosted = moment(DatePosted).fromNow();
      
      return datePosted
  }

  const stringToColor = (str: string) => {
    let hash = 0;
    for (let i = 0; i < str.length; i++) {
        hash = str.charCodeAt(i) + ((hash << 5) - hash);
    }

    const r = ((hash >> 16) & 0xff) % 140;
    const g = ((hash >> 8) & 0xff) % 140;
    const b = (hash & 0xff) % 140;
    return `#${r.toString(16).padStart(2,'0')}${g.toString(16).padStart(2,'0')}${b.toString(16).padStart(2,'0')}`;
  };

  return ( //this is funny, i'm tire and i don't fully understand the Renderer with json and the same styles, so i just implement the classNames of tiptap editor xD, i'm think is more simple
    <div className={`${className} !mt-5`}>
        <div className='flex items-center gap-5 border-b-1 border-white  pb-[30px] pl-[3rem] pt-[30px]'>
          {profileImage ? <img className="border-white border-1 w-[50px] h-[50px] rounded-[100%]" src={profileImage} alt="profileImage" /> : <p 
              style={{ backgroundColor: stringToColor(author ?? "?") }} 
              className="border-white p-4 border-1 w-[50px] h-[50px] rounded-[100%] text-white flex justify-center items-center"
          >
              {author?.[0] ?? ""}
          </p>}
          <div>
            <h2 className='text-2xl font-bold'>{author}</h2>
            <p className='text-gray-500'>Posted {DateAgo(DatePublished)}</p>
          </div>
        </div>
        <div className='tiptap ProseMirror simple-editor' dangerouslySetInnerHTML={{ __html: html }}>
        
        </div>
    </div>
    
  )
}