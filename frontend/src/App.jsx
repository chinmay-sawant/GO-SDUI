import React, {useEffect, useState} from 'react'
import ComponentRenderer from './ComponentRenderer'
import NavBar from './NavBar'
import ListDetail from './screens/ListDetail'
import Detail from './screens/Detail'
import Forms from './screens/Forms'
import Articles from './screens/Articles'

const API_BASE = 'http://localhost:8080/api/ui'

export default function App(){
  const [nav, setNav] = useState([])
  const [currentPath, setCurrentPath] = useState('/home')
  const [uiComponents, setUiComponents] = useState([])

  useEffect(()=>{
    fetch(`${API_BASE}/nav`).then(r=>r.json()).then(setNav).catch(console.error)
  },[])

  useEffect(()=>{
    const endpoint = `${API_BASE}${currentPath}`
    fetch(endpoint).then(r=>r.json()).then(setUiComponents).catch(err=>{
      console.error(err)
      setUiComponents([{type:'paragraph', props:{text:'Failed to load'}}])
    })
  },[currentPath])

  const handleAction = (actionId)=>{
    if(!actionId) return
    if(actionId.startsWith('navigate_')){
      const dest = '/' + actionId.replace('navigate_','')
      setCurrentPath(dest)
    } else if(actionId.startsWith('view_detail_')){
      const id = actionId.replace('view_detail_','')
      setCurrentPath(`/detail/${id}`)
    } else if(actionId === 'user_logout'){
      alert('Logged out (demo)')
    } else if(actionId === 'form_submit'){
      alert('Form submitted (demo)')
    }
  }

  const renderScreen = () => {
    if(currentPath.startsWith('/detail/')){
      return <Detail components={uiComponents} onAction={handleAction} />
    }
    switch(currentPath){
      case '/list-detail': return <ListDetail components={uiComponents} onAction={handleAction} />
      case '/forms': return <Forms components={uiComponents} onAction={handleAction} />
      case '/articles': return <Articles components={uiComponents} onAction={handleAction} />
      default: return <ComponentRenderer components={uiComponents} onAction={handleAction} />
    }
  }

  return (
    <div className="app">
      <NavBar nav={nav} onNavigate={p=>setCurrentPath(p)} active={currentPath} />
      <main className={`container ${currentPath.slice(1).replace('/','')}-page`}>
        {renderScreen()}
      </main>
    </div>
  )
}