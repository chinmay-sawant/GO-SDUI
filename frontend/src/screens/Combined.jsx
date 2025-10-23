import React from 'react'
import ComponentRenderer from '../ComponentRenderer'

export default function Combined({components, onAction}){
  return (
    <div className="page combined-page">
      <div className="card">
        <h2>Combined</h2>
        <ComponentRenderer components={components} onAction={onAction} />
      </div>
    </div>
  )
}
