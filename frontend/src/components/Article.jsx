import React from 'react'

function Article({props}){
  const {title, content, author, date} = props
  return (
    <div className="article">
      <h3>{title}</h3>
      <p>{content}</p>
      <div className="article-meta">
        <span>By {author}</span> <span>{date}</span>
      </div>
    </div>
  )
}

export default Article