import React from 'react'
import styles from './HomePage.module.css'
import MainHeader from '../../components/headers/main/MainHeader'
import Dashboard from '../dashBoard/Dashboard'
import { Routes, Route } from 'react-router-dom'
import SellerPage from '../seller/SellerPage'


function HomePage() {
  return (
    <div className={`${styles.homePage}`}>
      <MainHeader className={`${styles.mainHeader}`}></MainHeader>
      <Routes>
        <Route path='/' element={<Dashboard></Dashboard>}></Route>
        <Route path='/sellers' element={<SellerPage></SellerPage>}></Route>
      </Routes>
    </div>
  )
}

export default HomePage