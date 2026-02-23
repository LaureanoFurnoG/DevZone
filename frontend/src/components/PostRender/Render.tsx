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
import CategoriesTags from '../CategoriesTags/CategoriesTags'
import moment from 'moment'
type Category = {
  id: number
  name: string
}

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
  categories: Category[]
  author?: string
  profileImage?: string
  className?: string
  DatePublished: string
}

export const TiptapRenderer: React.FC<TiptapRendererProps> = ({
  content,
  categories,
  author,
  profileImage,
  DatePublished,
  className = 'simple-editor-content'
}) => {

  const html = React.useMemo(() => {
    if (!content) return ''

    try {
      return generateHTML(content, [StarterKit])
    } catch (error) {
      console.error('Error:', error)
      return ''
    }
  }, [content])

  const DateAgo = (DatePosted: string) =>{
      const datePosted = moment(DatePosted).fromNow();
      
      return datePosted
  }
  
  return ( //this is funny, i'm tire and i don't fully understand the Renderer with json and the same styles, so i just implement the classNames of tiptap editor xD, i'm think is more simple
    <div className={`${className} !mt-5`}>
        <div className='flex items-center gap-5 border-b-1 border-white  pb-[30px] pl-[3rem] pt-[30px]'>
          <img className='h-15 w-15 rounded-[100px] border-1 border-white' src={profileImage} alt="avatar of the user (profile image)" />
          <div>
            <h2 className='text-2xl font-bold'>{author}</h2>
            <p className='text-gray-500'>Posted {DateAgo(DatePublished)}</p>
          </div>
        </div>
        <div className='tiptap ProseMirror simple-editor' dangerouslySetInnerHTML={{ __html: html }}>
        
        </div>
        <div className='flex p-10 gap-10 border-t-1 border-white'>
          {categories.map((category) =>(
              <CategoriesTags Title={category.name} />
          ))}
        </div>
    </div>
    
  )
}