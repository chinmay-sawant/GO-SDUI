import React from 'react'
import ComponentRenderer from '../ComponentRenderer'

function List({props, onAction}){
  const items = props.items || []
  return (
    <div className="list">
      {items.map((it, idx)=> (
        <ComponentRenderer key={idx} components={[it]} onAction={onAction} />
      ))}
    </div>
  )
}

export default List