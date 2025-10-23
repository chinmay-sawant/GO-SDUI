import React from 'react'

function FormInput({props}){
  const {label, name, type} = props
  if(type === 'textarea'){
    return (
      <div className="form-group">
        <label htmlFor={name}>{label}</label>
        <textarea id={name} name={name} />
      </div>
    )
  }
  return (
    <div className="form-group">
      <label htmlFor={name}>{label}</label>
      <input id={name} name={name} type={type || 'text'} />
    </div>
  )
}

export default FormInput