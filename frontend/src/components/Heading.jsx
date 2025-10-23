import React from 'react'

function Heading({props}){
  const level = props.level || 1
  const Tag = `h${Math.min(Math.max(level,1),6)}`
  return <Tag>{props.text}</Tag>
}

export default Heading