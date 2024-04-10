import React from 'react'
import styles from './SellerPage.module.css'
import SellerSubHeader from '../../components/headers/sellerSubHeader/SellerSubHeader'

function SellerPage() {
    var users = [
        {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        },
        {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        }
        , {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        }, {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        },
        {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        }, {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        }
        , {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        }, {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        },
        {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        }, {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        }
        , {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        }, {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        },
        {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        }, {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        }
        , {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        }, {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        },
        {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        }, {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        }
        , {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        }, {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        },
        {
            name: "Olivia Martin",
            mail: "olivia.martin@email.com",
            amount: 1999.00
        }

    ]
    return (
        <div className={`${styles.sellerPage}`}>
            <SellerSubHeader className={`${styles.subHeader}`}></SellerSubHeader>
            <div className={`${styles.list}`}>
                {users.map((user) => <div className={`${styles.userItem} round-border`}>
                    <div className={`${styles.userItemDetails} `}>
                        <img src="https://ui.shadcn.com/avatars/01.png" alt="" srcset="" className='icon' />
                        <div>
                            <p>{user.name}</p>
                            <p>{user.mail}</p>
                        </div>
                    </div>
                    <p>+${user.amount}.00</p>
                </div>)}

            </div>
            <div className={`${styles.details}`}>Details</div>
        </div>
    )
}

export default SellerPage