import React from 'react'
import { generateHTML } from '@tiptap/core'
import StarterKit from '@tiptap/starter-kit'
interface TiptapRendererProps {
  content: any 
  className?: string
}

export const TiptapRenderer: React.FC<TiptapRendererProps> = ({
  content,
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

  return ( //this is funny, i'm tire and i don't fully understand the Renderer with json and the same styles, so i just implement the classNames of tiptap editor xD, i'm think is more simple
    <div className={`prose max-w-none ${className}`}>
        <div className='tiptap ProseMirror simple-editor' dangerouslySetInnerHTML={{ __html: html }}>

        </div>
    </div>
    
  )
}