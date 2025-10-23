import React from 'react'

function FormInput({props}){
  const {label, name, type, placeholder, value, onChange} = props
  if(type === 'textarea'){
    return (
      <div className="form-group">
        <label htmlFor={name}>{label}</label>
        <textarea id={name} name={name} placeholder={placeholder || ''} defaultValue={value || ''} />
      </div>
    )
  }
  return (
    <div className="form-group">
      <label htmlFor={name}>{label}</label>
      <input id={name} name={name} type={type || 'text'} placeholder={placeholder || ''} defaultValue={value || ''} />
    </div>
  )
}

export default FormInput