import React from 'react'

function FormButton({props, onAction}){
  return <button onClick={()=>onAction(props.actionId)}>{props.text}</button>
}

export default FormButton