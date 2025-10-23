import React from 'react'

function ProductCard({props}){
  return (
    <div className="product-card">
      <img src={props.imageSrc} alt={props.name} />
      <div className="product-info">
        <strong>{props.name}</strong>
        <div>{props.price}</div>
      </div>
    </div>
  )
}

export default ProductCard