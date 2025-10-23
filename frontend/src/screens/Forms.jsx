import React from 'react'
import ComponentRenderer from '../ComponentRenderer'

export default function Forms({components, onAction}){
  return (
    <div className="page forms-page">
      <div className="card">
        <h2>Form</h2>
        <ComponentRenderer components={components} onAction={onAction} />
      </div>
    </div>
  )
}
