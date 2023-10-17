<template>
  <div class="q-pa-md">
    <q-layout view="hHh Lpr lff" container style="height: 500px" class="shadow-2 rounded-borders">
      <canvas id="myChart" style="height: 500px"></canvas>
    </q-layout>

    <div style="display: flex; justify-content: space-between;">
      <q-layout view="hHh Lpr lff" container style="height: 500px ; margin-top: 2%;" class="shadow-2 rounded-borders">
        <canvas id="productChart" style="height: 500px"></canvas>
      </q-layout>

    
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import Chart from 'chart.js/auto';
import { format, parseISO } from 'date-fns';

export default {
  data() {
    return {
      orderChartData: [],  
      orderChart: null,  
      productChartData: [], 
      productChart: null, 
    };
  },
  async created() {
    try {
      await this.fetchDataAndCreateChart();
    } catch (e) {
      console.error(e);
    }
  },
  methods: {
    async fetchDataAndCreateChart() {
      try {
        const token = localStorage.getItem('token');
        const headers = {
          'Authorization': `${token}`
        };

        const ordersResponse = await axios.get('http://localhost:8080/orders', { headers });
        const productsResponse = await axios.get('http://localhost:8080/products', { headers });
        const ordersData = ordersResponse.data;
        const productsData = productsResponse.data;
        console.log(ordersData)
        console.log(productsData)

        this.orderChartData  = ordersData.map((order) => ({
          ...order,
         
          timestamp: format(parseISO(order.create_date), 'MM/dd HH:mm'),
          customerName: `Customer Name： ${order.customer_name}`,
          productName: order.Product.product_name, 
            
        }));

        this.orderChartUpdate();


        this.productChartData = productsData.map((product) => ({
            ...product,
        }));
        this.productChartUpdate();

        
      } catch (e) {
        console.error(e);
      }
    },
    orderChartUpdate() {
      const ctx = document.getElementById('myChart').getContext('2d');

      if (this.orderChart && this.orderChart.data) {
        this.orderChart.data.labels = this.chartData.map((dataItem) => `${dataItem.customerName} - ${dataItem.timestamp}`);
        this.orderChart.data.datasets[0].data = this.orderChart.data.map((dataItem) => dataItem.price);
        this.orderChart.data.datasets[1].data = this.orderChart.data.map((dataItem) => dataItem.quantity);
        
        this.orderChart.update();
      } else {
        this.orderChart = new Chart(ctx, {
          type: 'bar',
          data: {
            labels: this.orderChartData.map((dataItem) => `${dataItem.customerName} ${dataItem.timestamp} `),
            datasets: [
              {
                label: 'Price',
                data: this.orderChartData.map((dataItem) => dataItem.price),
                backgroundColor: 'rgba(75, 192, 192, 0.2)',
                borderColor: 'rgba(75, 192, 192, 1)',
                borderWidth: 1,
              },
              {
                label: 'Quantity',
                data: this.orderChartData.map((dataItem) => dataItem.quantity),
                backgroundColor: 'rgba(255, 99, 132, 0.2)',
                borderColor: 'rgba(255, 99, 132, 1)',
                borderWidth: 1,
              },
             
            ],
          },
          options: {
            responsive: true,
            maintainAspectRatio: false,
            plugins: {
                title: {
                    display: true,
                    text: 'Custom Order',
                    font: {
                        size: 20
                    }
                }
            }
          },
        });
      }
    },
    productChartUpdate() {
      const ctx = document.getElementById('productChart').getContext('2d');

      if (this.productChart && this.productChart.data) {
        this.productChart.data.labels = this.productChartData.map((dataItem) => `${dataItem.product_name}`);
        this.productChart.data.datasets[0].data = this.productChartData.map((dataItem) => dataItem.quantity);

        this.productChart.update();
      } else {
        this.productChart = new Chart(ctx, {
          type: 'bar',
          data: {
            labels: this.productChartData.map((dataItem) => `${dataItem.product_name}`),
            datasets: [
              {
                label: 'Quantity',
                data: this.productChartData.map((dataItem) => dataItem.quantity),
                backgroundColor: 'rgba(255, 99, 132, 0.2)',
                borderColor: 'rgba(255, 99, 132, 1)',
                borderWidth: 1,
              },
            ],
          },
          options: {
            responsive: true,
            maintainAspectRatio: false,
            plugins: {
              title: {
                display: true,
                text: 'Product inventory',
                font: {
                  size: 20
                }
              }
            }
          },
        });
      }
    },
  },
};
</script>
