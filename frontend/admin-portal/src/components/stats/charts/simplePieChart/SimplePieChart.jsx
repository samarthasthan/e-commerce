import React from 'react'
import ReactApexChart from "react-apexcharts";
import styles from './SimplePieChart.module.css'

function SimplePieChart() {
  var options = {
    series: [44, 55, 13, 43, 22],
    options: {
      chart: {

        type: 'pie',
      },
      labels: ['Fashion', 'Electronics', 'Kids', 'Cosmetics', 'SmartPhones'],
      responsive: [{
        breakpoint: 480,
        options: {
          chart: {
            width: 200
          },
          legend: {
            position: 'bottom'
          }
        }
      }]
    },


  }

  return (
    <div className={`round-border ${styles.pieChart}`}>
      <h3 className='round-border-heading'>Sales by Category</h3>
      <ReactApexChart options={options} series={options.series} type="pie" height={"300px"} width={"100%"} />
    </div>
  )
}

export default SimplePieChart