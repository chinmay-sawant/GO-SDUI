import React from 'react'
import ComponentRenderer from '../ComponentRenderer'

function ProductList({props, onAction}){
  const items = props.items || []
  return (
    <div className="product-list">
      {items.map((it, idx)=> (
        <ComponentRenderer key={idx} components={[it]} onAction={onAction} />
      ))}
    </div>
  )
}

export default ProductList