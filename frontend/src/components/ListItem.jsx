import React from 'react'
import Button from './Button'

function ListItem({props, onAction}){
  return (
    <div className="list-item">
      <div className="row">
        <div className="col">
          <h3>{props.title}</h3>
          <p className="muted">{props.description}</p>
        </div>
        <div className="col" style={{flex:'0 0 160px', display:'flex', alignItems:'center', justifyContent:'flex-end'}}>
          <Button props={{text: "View Detail", actionId: `view_detail_${props.id}`}} onAction={onAction} />
        </div>
      </div>
    </div>
  )
}

export default ListItem