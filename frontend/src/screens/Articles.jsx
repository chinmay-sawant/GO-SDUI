import React from 'react'
import ComponentRenderer from '../ComponentRenderer'

export default function Articles({components, onAction}){
  return (
    <div className="page articles-page">
      <div className="card">
        <h2>Articles</h2>
        <ComponentRenderer components={components} onAction={onAction} />
      </div>
    </div>
  )
}
