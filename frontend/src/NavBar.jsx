import React from 'react'

export default function NavBar({nav, onNavigate, active}){
  return (
    <nav className="nav">
      <div className="nav-inner">
        {nav.map((n,i)=> (
          <button key={i} className={active===n.path? 'nav-item active':'nav-item'} onClick={()=>onNavigate(n.path)}>{n.label}</button>
        ))}
      </div>
    </nav>
  )
}