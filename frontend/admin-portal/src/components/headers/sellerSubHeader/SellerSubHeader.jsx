import React from 'react'
import styles from './SellerSubHeader.module.css'
import SubNavBar from '../../cards/statsCard/subNavBar/SubNavBar'

function SellerSubHeader() {
    return (
        <div className={`${styles.sellerSubHeader}`}>
            <SubNavBar></SubNavBar>
        </div>
    )
}

export default SellerSubHeader