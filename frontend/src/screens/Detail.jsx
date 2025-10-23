import React from 'react'
import ComponentRenderer from '../ComponentRenderer'

export default function Detail({components, onAction}){
  return (
    <div className="detail-page">
      <ComponentRenderer components={components} onAction={onAction} />
    </div>
  )
}
