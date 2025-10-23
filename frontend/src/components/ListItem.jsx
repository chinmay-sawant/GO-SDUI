import React from 'react'
import Button from './Button'

function ListItem({props, onAction}){
  return (
    <div className="list-item">
      <h3>{props.title}</h3>
      <p>{props.description}</p>
      <Button props={{text: "View Detail", actionId: `view_detail_${props.id}`}} onAction={onAction} />
    </div>
  )
}

export default ListItem