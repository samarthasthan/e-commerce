import React from 'react'
import styles from './SubNavBar.module.css'
import { Envelope } from "@phosphor-icons/react";

function SubNavBar() {
    return (
        <div className={`${styles.subNavBar} round-border`}>
            <div className={`${styles.details}`}>
                <Envelope className='icon' />
                <p>Inbox</p>
            </div>
            <p className={`${styles.count}`}>128</p>
        </div>
    )
}

export default SubNavBar