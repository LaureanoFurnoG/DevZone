import { useEditor, EditorContent } from '@tiptap/react'
import StarterKit from '@tiptap/starter-kit'
import Underline from '@tiptap/extension-underline'
import TextAlign from '@tiptap/extension-text-align'
import Highlight from '@tiptap/extension-highlight'
import Link from '@tiptap/extension-link'
import Superscript from '@tiptap/extension-superscript'
import Subscript from '@tiptap/extension-subscript'
import Image from '@tiptap/extension-image'
import TableHeader from '@tiptap/extension-table-header'
import './style.css'

const Tiptap = () => {
  const editor = useEditor({
    extensions: [
      StarterKit,
      Underline,
      TextAlign.configure({
        types: ['heading', 'paragraph'],
      }),
      Highlight,
      Link.configure({
        openOnClick: false,
      }),
      Superscript,
      Subscript,
      Image,
      TableHeader,
    ],
    content: '',
    editorProps: {
      attributes: {
        class: 'outline-none focus:outline-none',
      },
    },
  })

  if (!editor) return null

  const addLink = () => {
    const url = window.prompt('Enter URL:')
    if (url) {
      editor.chain().focus().setLink({ href: url }).run()
    }
  }

  const addImage = () => {
    const url = window.prompt('Enter image URL:')
    if (url) {
      editor.chain().focus().setImage({ src: url }).run()
    }
  }

  return (
    <div className="border rounded border-white h-130 bg-[#1D1D1D]">
      <div className="flex gap-2 h-15 p-2 pl-8 pr-8 justify-between border-b border-white bg-[#1D1D1D] flex-wrap text-white">
        <select 
          onChange={(e) => {
            const level = parseInt(e.target.value)
            if (level === 0) {
              editor.chain().focus().setParagraph().run()
            } else {
              editor.chain().focus().toggleHeading({ level: 2 }).run()
            }
          }}
          className="bg-[#2D2D2D] text-white px-2 rounded"
        >
          <option value="0">Normal</option>
          <option value="1">H1</option>
          <option value="2">H2</option>
          <option value="3">H3</option>
          <option value="4">H4</option>
          <option value="5">H5</option>
          <option value="6">H6</option>
        </select>

        <button 
          onClick={() => editor.chain().focus().toggleBulletList().run()}
          className={editor.isActive('bulletList') ? 'bg-gray-600' : 'cursor-pointer'}
        >
          •
        </button>
        <button 
          onClick={() => editor.chain().focus().toggleOrderedList().run()}
          className={editor.isActive('orderedList') ? 'bg-gray-600' : 'cursor-pointer'}
        >
          1.
        </button>
        <button className= 'cursor-pointer' onClick={() => editor.chain().focus().undo().run()}>
          ↶
        </button>
        <button className= 'cursor-pointer' onClick={() => editor.chain().focus().redo().run()}>
          ↷
        </button>

        <button 
          onClick={() => editor.chain().focus().toggleBold().run()}
          className={editor.isActive('bold') ? 'bg-gray-600' : 'cursor-pointer'}
        >
          <strong>B</strong>
        </button>
        <button 
          onClick={() => editor.chain().focus().toggleItalic().run()}
          className={editor.isActive('italic') ? 'bg-gray-600' : 'cursor-pointer'}
        >
          <em>I</em>
        </button>
        <button 
          onClick={() => editor.chain().focus().toggleStrike().run()}
          className={editor.isActive('strike') ? 'bg-gray-600' : 'cursor-pointer'}
        >
          <s>S</s>
        </button>
        <button 
          onClick={() => editor.chain().focus().toggleCode().run()}
          className={editor.isActive('code') ? 'bg-gray-600' : 'cursor-pointer'}
        >
        </button>
        <button 
          onClick={() => editor.chain().focus().toggleUnderline().run()}
          className={editor.isActive('underline') ? 'bg-gray-600' : 'cursor-pointer'}
        >
          <u>U</u>
        </button>
        <button 
          onClick={() => editor.chain().focus().toggleHighlight().run()}
          className={editor.isActive('highlight') ? 'bg-gray-600' : 'cursor-pointer'}
        >
          🖍️
        </button>

        <button 
          onClick={addLink}
          className={editor.isActive('link') ? 'bg-gray-600' : 'cursor-pointer'}
        >
          🔗
        </button>
        {editor.isActive('link') && (
          <button onClick={() => editor.chain().focus().unsetLink().run()}>
            ❌ Link
          </button>
        )}

        <button 
          onClick={() => editor.chain().focus().toggleSuperscript().run()}
          className={editor.isActive('superscript') ? 'bg-gray-600' : 'cursor-pointer'}
        >
          x²
        </button>

        <button 
          onClick={() => editor.chain().focus().toggleSubscript().run()}
          className={editor.isActive('subscript') ? 'bg-gray-600' : 'cursor-pointer'}
        >
          x₂
        </button>

        <button 
          onClick={() => editor.chain().focus().setTextAlign('left').run()}
          className={editor.isActive({ textAlign: 'left' }) ? 'bg-gray-600' : 'cursor-pointer'}
        >
          ⬅️
        </button>
        <button 
          onClick={() => editor.chain().focus().setTextAlign('center').run()}
          className={editor.isActive({ textAlign: 'center' }) ? 'bg-gray-600' : 'cursor-pointer'}
        >
          ↔️
        </button>
        <button 
          onClick={() => editor.chain().focus().setTextAlign('right').run()}
          className={editor.isActive({ textAlign: 'right' }) ? 'bg-gray-600' : ''}
        >
          ➡️
        </button>
        <button onClick={addImage}>
          🖼️ Add Image
        </button>
      </div>

      <EditorContent editor={editor} className="p-4 text-white min-h-[400px]" />
    </div>
  )
}

export default Tiptap