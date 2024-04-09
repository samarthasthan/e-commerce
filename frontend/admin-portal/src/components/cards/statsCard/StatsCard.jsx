import React from 'react'
import styles from './StatsCard.module.css'
import {CurrencyDollar}  from "@phosphor-icons/react";

function StatsCard() {
  return (
    <div className={`${styles.statsCard}`}>
        <div className={`${styles.title}`}>
            Total Revenue
            <div className={`icon`}> 
            <CurrencyDollar />
            </div>
        </div>
        <h3>$45,231.89</h3>
        <p className='greyText'>+20.1% from last month</p>
    </div>
  )
}

export default StatsCard