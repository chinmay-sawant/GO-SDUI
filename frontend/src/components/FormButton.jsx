import React from 'react'

function FormButton({props, onAction}){
  const {text, actionId, variant} = props
  return (
    <div className="form-actions">
      <button className={variant === 'primary' ? 'primary' : ''} onClick={()=>onAction(actionId)} type="button">{text}</button>
    </div>
  )
}

export default FormButton