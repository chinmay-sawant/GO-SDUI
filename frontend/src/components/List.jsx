import React from 'react'
import ComponentRenderer from '../ComponentRenderer'

function List({props, onAction}){
  const items = props.items || []
  return (
    <div className="list">
      {items.map((it, idx)=> (
        <div key={idx} className="list-row card">
          <ComponentRenderer components={[it]} onAction={onAction} />
        </div>
      ))}
    </div>
  )
}

export default List