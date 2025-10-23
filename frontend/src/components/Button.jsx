import React from 'react'

function Button({props, onAction}){
  return <button onClick={()=>onAction(props.actionId)}>{props.text}</button>
}

export default Button