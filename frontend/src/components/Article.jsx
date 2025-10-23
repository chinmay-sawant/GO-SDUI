import React from 'react'

function Article({props}){
  const {title, content, author, date} = props
  return (
    <article className="article">
      <header>
        <h3>{title}</h3>
      </header>
      <section>
        <p>{content}</p>
      </section>
      <footer className="article-meta muted">
        <span>By {author}</span>
        <span>{date}</span>
      </footer>
    </article>
  )
}

export default Article