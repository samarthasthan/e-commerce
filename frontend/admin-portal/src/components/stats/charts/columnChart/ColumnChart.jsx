import React from 'react'
import ReactApexChart from "react-apexcharts";
import styles from './ColumnChart.module.css'

function ColumnChart() {
    var options = {
        series: [{
            name: "Sales",
        data: [21, 22, 10, 28, 16, 21, 13, 30,34,12,37,23.10]
      }],
        chart: {
        type: 'bar',
        toolbar: {
            show: false,
          },
        events: {
          click: function(chart, w, e) {
            // console.log(chart, w, e)
          }
        }
      },
      
      tooltip: {
        enabled: true,
      },
      grid: {
        show: false,
      },
      colors: ['#18181b'],
      plotOptions: {
        bar: {
          columnWidth: '45%',
          
          distributed: true,
        }
      },
      dataLabels: {
        enabled: false
      },
      legend: {
        show: false
      },
      xaxis: {
        categories: [
          ['Jan'],
          ['Feb'],
          ['Mar'],
          ['Apr'],
          ['May'],
          ['Jun'],
          ['Jul'],
          ['Aug'],
          ['Sep'],
          ['Oct'],
          ['Nov'],
          ['Dec'],
        ],
        labels: {
          style: {
            colors: ['#18181b'],
            fontSize: '12px'
          }
        }
      }
      };
  return (
    <div className={`round-border `}>
        <ReactApexChart options={options} series={options.series} type="bar" height={400} width={`250%`} className={`${styles.columnChar}`} />
        <ReactApexChart options={options} series={options.series} type="bar" height={400} width={`100%`} className={`${styles.columnCharMobile}`} />
    </div>
  )
}

export default ColumnChart