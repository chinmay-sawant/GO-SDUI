import React from 'react'
import ComponentRenderer from '../ComponentRenderer'

export default function ListDetail({components, onAction}){
  return (
    <div className="list-detail-page">
      <div className="page list-detail-page">
        <div className="card">
          <ComponentRenderer components={components} onAction={onAction} />
        </div>
      </div>
    </div>
  )
}
