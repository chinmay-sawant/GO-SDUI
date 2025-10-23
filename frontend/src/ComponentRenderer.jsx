import React from 'react'
import Heading from './components/Heading'
import Paragraph from './components/Paragraph'
import Image from './components/Image'
import Button from './components/Button'
import ProductCard from './components/ProductCard'
import ProductList from './components/ProductList'
import StatItem from './components/StatItem'
import Divider from './components/Divider'
import List from './components/List'
import ListItem from './components/ListItem'
import FormInput from './components/FormInput'
import FormButton from './components/FormButton'
import Article from './components/Article'

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
          case 'divider': return <Divider key={i} />
          case 'list': return <List key={i} props={props} onAction={onAction} />
          case 'list_item': return <ListItem key={i} props={props} onAction={onAction} />
          case 'form_input': return <FormInput key={i} props={props} />
          case 'form_button': return <FormButton key={i} props={props} onAction={onAction} />
          case 'article': return <Article key={i} props={props} />
          default: return <div key={i} className="unknown-component">Unknown component: {c.type}</div>
        }
      })}
    </div>
  )
}