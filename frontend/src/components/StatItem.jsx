import React from 'react'

function StatItem({props}){
  return (
    <div className="stat-item">
      <span className="label">{props.label}:</span> <span className="value">{props.value}</span>
    </div>
  )
}

export default StatItem