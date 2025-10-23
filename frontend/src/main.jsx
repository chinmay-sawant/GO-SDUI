import React from 'react'
import { createRoot } from 'react-dom/client'
import App from './App'
import './css/theme.css'
import './css/nav.css'
import './css/layout.css'
import './css/components.css'
import './css/pages.css'

createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
)