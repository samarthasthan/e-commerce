import React from 'react'
import styles from './SellerPage.module.css'
import SellerSubHeader from '../../components/headers/sellerSubHeader/SellerSubHeader'

function SellerPage() {
    return (
        <div className={`${styles.sellerPage}`}>
            <SellerSubHeader></SellerSubHeader>
            <div>List</div>
            <div>Details</div>
        </div>
    )
}

export default SellerPage