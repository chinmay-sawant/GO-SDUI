import React from 'react'
import ComponentRenderer from '../ComponentRenderer'

export default function ListDetail({components, onAction}){
  return (
    <div className="list-detail-page">
      <ComponentRenderer components={components} onAction={onAction} />
    </div>
  )
}
