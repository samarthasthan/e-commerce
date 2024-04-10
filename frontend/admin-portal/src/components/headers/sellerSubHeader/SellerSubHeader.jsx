import React from 'react'
import styles from './SellerSubHeader.module.css'
import SubNavBar from '../../cards/subNavBar/SubNavBar'

function SellerSubHeader() {
    var list = [1, 2, 3, 4, 5, 6, 7, 8, 9]
    return (
        <div className={`${styles.sellerSubHeader}`}>
            {list.map((item, index) => index === 0 ? <SubNavBar></SubNavBar> : <SubNavBar></SubNavBar>)}
        </div>
    )
}

export default SellerSubHeader