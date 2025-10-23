import React from 'react'
import ComponentRenderer from '../ComponentRenderer'

export default function Articles({components, onAction}){
  return (
    <div className="articles-page">
      <ComponentRenderer components={components} onAction={onAction} />
    </div>
  )
}
