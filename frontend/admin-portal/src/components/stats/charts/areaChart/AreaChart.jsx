import React from 'react'
import ReactApexChart from "react-apexcharts";
import styles from './AreaChart.module.css'

function AreaChart() {


    const chartOptions = {
        // Define your chart options here
        chart: {
            type: "line",
            toolbar: {
                show: false,
            },
        },
        series: [
            {
                name: "Price",
                data: [10000, 30000, 5000, 40000, 20000, 25000],
            },
        ],
        xaxis: {
            labels: {
                show: false,
            },
            axisBorder: {
                show: false,
            },
            axisTicks: {
                show: false,
            },
        },
        yaxis: {
            labels: {
                show: false,
            },
        },
        tooltip: {
            enabled: false,
        },
        colors: ["#18181b"],
        stroke: {
            show: true,
            curve: "smooth",
            lineCap: "butt",
            colors: undefined,
            width: 2,
            dashArray: 0,
        },
        grid: {
            show: false,
        },
        selection: {
            enabled: true,
        },
    };
    return (
        <div className={`round-border ${styles.areaChart}`}>
            <h3 className='round-border-heading'>Total Users</h3>
            <ReactApexChart options={chartOptions} series={chartOptions.series} type="area" height={"300px"} width={"100%"} />
        </div>
    )
}

export default AreaChart