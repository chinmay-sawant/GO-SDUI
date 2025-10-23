import React from 'react'
import ComponentRenderer from '../ComponentRenderer'

export default function Detail({components, onAction}){
  return (
    <div className="page detail-page">
      <div className="detail card">
        <div className="detail-main">
          <ComponentRenderer components={components} onAction={onAction} />
        </div>
      </div>
    </div>
  )
}
