import React from 'react'
import styles from './HomePage.module.css'
import MainHeader from '../../components/headers/main/MainHeader'
import Dashboard from '../dashBoard/Dashboard'
function HomePage() {
  return (
    <div className={`${styles.homePage}`}>
         <MainHeader></MainHeader>
         <Dashboard></Dashboard>
    </div>
  )
}

export default HomePage