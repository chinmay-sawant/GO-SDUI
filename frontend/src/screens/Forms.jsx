import React from 'react'
import ComponentRenderer from '../ComponentRenderer'

export default function Forms({components, onAction}){
  return (
    <div className="forms-page">
      <ComponentRenderer components={components} onAction={onAction} />
    </div>
  )
}
