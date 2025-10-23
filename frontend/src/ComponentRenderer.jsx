import React from 'react'

function Heading({props}){
  const level = props.level || 1
  const Tag = `h${Math.min(Math.max(level,1),6)}`
  return <Tag>{props.text}</Tag>
}

function Paragraph({props}){
  return <p>{props.text}</p>
}

function Image({props}){
  return <img src={props.src} alt={props.alt||''} style={{maxWidth:'100%'}} />
}

function Button({props, onAction}){
  return <button onClick={()=>onAction(props.actionId)}>{props.text}</button>
}

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

function StatItem({props}){
  return (
    <div className="stat-item">
      <div className="label">{props.label}</div>
      <div className="value">{props.value}</div>
    </div>
  )
}

export default function ComponentRenderer({components, onAction}){
  if(!components) return null
  return (
    <div>
      {components.map((c, i)=>{
        const props = c.props || {}
        switch(c.type){
          case 'heading': return <Heading key={i} props={props} />
          case 'paragraph': return <Paragraph key={i} props={props} />
          case 'image': return <Image key={i} props={props} />
          case 'button': return <Button key={i} props={props} onAction={onAction} />
          case 'product_list': return <ProductList key={i} props={props} onAction={onAction} />
          case 'product_card': return <ProductCard key={i} props={props} />
          case 'stat_item': return <StatItem key={i} props={props} />
          case 'divider': return <hr key={i} />
          default: return <div key={i}>Unknown component: {c.type}</div>
        }
      })}
    </div>
  )
}