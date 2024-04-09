import React from 'react'
import styles from './Dashboard.module.css'
import StatsCard from '../../components/cards/statsCard/StatsCard'
import ColumnChart from '../../components/stats/charts/columnChart/ColumnChart'

function Dashboard() {
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
    }
    , {
      name: "Olivia Martin",
      mail: "olivia.martin@email.com",
      amount: 1999.00
    }, {
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
    }
  ]
  return (
    <div className={`${styles.dashBoard}`}>
      <h2>Dashboard</h2>
      <div className={`${styles.cards}`}>
        <StatsCard></StatsCard>
        <StatsCard></StatsCard>
        <StatsCard></StatsCard>
        <StatsCard></StatsCard>
      </div>
      <div className={`${styles.stats}`}>
        <ColumnChart className={`${styles.columnChart}`}></ColumnChart>
        <div className={`${styles.sales} round-border `}>
          <h3>Recent Sales</h3>
          <p className='greyText'>You made 265 sales this month.</p>
          <div className={`${styles.list}`}>
            {users.map((user) => <div className={styles.userItem}>
            <div className={`${styles.userItemDetails}`}>
              <img src="https://ui.shadcn.com/avatars/01.png" alt="" srcset="" className='icon' />
              <div>
                <p>{user.name}</p>
                <p>{user.mail}</p>
              </div>
            </div>
            <p>+${user.amount}.00</p>
          </div>)}
          </div>
        </div>
      </div>

    </div>
  )
}

export default Dashboard