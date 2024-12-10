<script setup>
import { onMounted, ref } from 'vue';
import { AdminService } from '@/service/AdminService';
import ProgressSpinner from 'primevue/progressspinner';

const statisticsData = ref(null);
const loading = ref(true);
const tableLoading = ref(false);
const orders = ref([]);

async function fetchStatisticsData() {
    try {
        const data = await AdminService.getStatisticsData().then(res => res.data);
        statisticsData.value = data;
    } catch (error) {
        console.error('Failed to fetch statistics data:', error);
    }
}

async function fetchLastOrders({ limit = 5 } = { limit: 5 }) {
    tableLoading.value = true;
    try {
        const data = await AdminService.getLastOrders({ limit: limit }).then(res => res.data);
        orders.value = data;
    } catch (error) {
        console.error('Failed to fetch last orders:', error);
    } finally {
        tableLoading.value = false;
    }
}

const getSeverity = (status) => {
    status = status.toLowerCase();
    switch (status) {
        case 'cancel':
            return 'danger';
        case 'completed':
            return 'success';
        case 'in progress':
            return 'info';
        case 'partials':
        case 'pending':
            return 'warn';
    }
};

onMounted(async () => {
    loading.value = true;
    await fetchStatisticsData();
    await fetchLastOrders();
    chartData.value = setChartData();
    chartOptions.value = setChartOptions();
    loading.value = false;
});

const chartData = ref();
const chartOptions = ref();

const setChartData = () => {
    return {
        labels: ['Total Transactions in 1 Month'],
        datasets: [
            {
                label: 'Transactions',
                data: [540],
                backgroundColor: ['rgba(249, 115, 22, 0.2)'],
                borderColor: ['rgb(249, 115, 22)'],
                borderWidth: 1
            }
        ]
    };
};

const setChartOptions = () => {
    const documentStyle = getComputedStyle(document.documentElement);
    const textColor = documentStyle.getPropertyValue('--p-text-color');
    const textColorSecondary = documentStyle.getPropertyValue('--p-text-muted-color');
    const surfaceBorder = documentStyle.getPropertyValue('--p-content-border-color');

    return {
        plugins: {
            legend: {
                labels: {
                    color: textColor
                }
            }
        },
        scales: {
            x: {
                ticks: {
                    color: textColorSecondary
                },
                grid: {
                    color: surfaceBorder
                }
            },
            y: {
                beginAtZero: true,
                ticks: {
                    color: textColorSecondary
                },
                grid: {
                    color: surfaceBorder
                }
            }
        }
    };
}

// Mock data for demonstration
const newDeposits = ref([
    { name: 'User 1', amount: 'Rp. 500.000' },
    { name: 'User 2', amount: 'Rp. 300.000' },
    { name: 'User 3', amount: 'Rp. 200.000' },
    { name: 'User 4', amount: 'Rp. 100.000' },
    { name: 'User 5', amount: 'Rp. 50.000' }
]);


const newUsers = ref([
    { username: 'user1', email: 'user1@example.com', registeredDate: '2023-10-01' },
    { username: 'user2', email: 'user2@example.com', registeredDate: '2023-10-02' },
    { username: 'user3', email: 'user3@example.com', registeredDate: '2023-10-03' },
    { username: 'user4', email: 'user4@example.com', registeredDate: '2023-10-04' },
    { username: 'user5', email: 'user5@example.com', registeredDate: '2023-10-05' }
]);
</script>

<template>
    <div v-if="loading" class="loading-screen">
        <ProgressSpinner />
    </div>
    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-8">
        <div class="col-span-1 md:col-span-1 lg:col-span-1 xl:col-span-1">
            <div class="card mb-0 border border-surface">
                <div class="flex justify-between mb-4">
                    <div>
                        <span class="block text-muted-color font-medium mb-4">Total Users</span>
                        <div class="text-surface-900 dark:text-surface-0 font-medium text-xl">
                            {{ statisticsData.total_users }}
                        </div>
                    </div>
                    <div class="flex items-center justify-center bg-blue-100 dark:bg-blue-400/10 rounded-border"
                         style="width: 2.5rem; height: 2.5rem">
                        <i class="pi pi-users text-blue-500 !text-xl"></i>
                    </div>
                </div>
            </div>
        </div>
        <div class="col-span-1 md:col-span-1 lg:col-span-1 xl:col-span-1">
            <div class="card mb-0 border border-surface">
                <div class="flex justify-between mb-4">
                    <div>
                        <span class="block text-muted-color font-medium mb-4">Total Orders</span>
                        <div class="text-surface-900 dark:text-surface-0 font-medium text-xl">
                            {{ statisticsData.total_orders }}
                        </div>
                    </div>
                    <div class="flex items-center justify-center bg-orange-100 dark:bg-orange-400/10 rounded-border"
                         style="width: 2.5rem; height: 2.5rem">
                        <i class="pi pi-shopping-cart text-orange-500 !text-xl"></i>
                    </div>
                </div>
            </div>
        </div>
        <div class="col-span-1 md:col-span-1 lg:col-span-1 xl:col-span-1">
            <div class="card mb-0 border border-surface">
                <div class="flex justify-between mb-4">
                    <div>
                        <span class="block text-muted-color font-medium mb-4">Total Revenue</span>
                        <div class="text-surface-900 dark:text-surface-0 font-medium text-xl">
                            {{ statisticsData.total_revenue }}
                        </div>
                    </div>
                    <div class="flex items-center justify-center bg-green-100 dark:bg-green-400/10 rounded-border"
                         style="width: 2.5rem; height: 2.5rem">
                        <i class="pi pi-dollar text-green-500 !text-xl"></i>
                    </div>
                </div>
            </div>
        </div>
        <div class="col-span-1 md:col-span-1 lg:col-span-1 xl:col-span-1">
            <div class="card mb-0 border border-surface">
                <div class="flex justify-between mb-4">
                    <div>
                        <span class="block text-muted-color font-medium mb-4">Total Tickets Open</span>
                        <div class="text-surface-900 dark:text-surface-0 font-medium text-xl">
                            {{ statisticsData.total_ticket_open }}
                        </div>
                    </div>
                    <div class="flex items-center justify-center bg-red-100 dark:bg-red-400/10 rounded-border"
                         style="width: 2.5rem; height: 2.5rem">
                        <i class="pi pi-ticket text-red-500 !text-xl"></i>
                    </div>
                </div>
            </div>
        </div>
        <div class="col-span-1 md:col-span-2 lg:col-span-2 xl:col-span-3">
            <Card>
                <template #title>
                    <h3 class="text-xl font-bold">Last 5 Transactions</h3>
                    <Divider/>
                </template>
                <template #content>
                    <div class="table-loading-overlay" v-if="tableLoading">
                        <ProgressSpinner />
                    </div>
                    <DataTable :value="orders" :loading="tableLoading" tableStyle="min-width: 50rem">
                        <Column field="name" header="Name" sortable />
                        <Column field="product" header="Product" sortable />
                        <Column field="total" header="Total" sortable />
                        <Column field="price" header="Price" sortable />
                        <Column field="status" header="Status">
                            <template #body="slotProps">
                                <Tag :value="slotProps.data.status" :severity="getSeverity(slotProps.data.status)" />
                            </template>
                        </Column>
                        <Column field="date" header="Date" sortable />
                    </DataTable>
                </template>
            </Card>
        </div>
        <div class="col-span- md:col-span-2 lg:col-span-1 xl:col-span-1 row-span-2">
            <Card>
                <template #title>
                    <h3 class="text-xl font-bold">Deposit</h3>
                    <Divider/>
                    <span class="text-bold text-muted-color">Total</span>
                    <h4 class="text-2xl font-bold">Rp. 1.500.000</h4>
                </template>
            </Card>
            <Card class="mt-4">
                <template #title>
                    <h3 class="text-xl font-bold">New User Registered</h3>
                    <Divider/>
                </template>
                <template #content>
                    <DataTable :value="newUsers">
                        <Column field="username" header="Username" sortable />
                        <Column field="email" header="Email" sortable />
                    </DataTable>
                </template>
            </Card>
        </div>
        <div class="col-span-1 md:col-span-1 lg:col-span-1 xl:col-span-2">
            <div class="col-span-1 md:col-span-1 lg:col-span-1 xl:col-span-1">
                <Card>
                    <template #title>
                        <h3 class="text-xl font-bold">Last 5 Deposits</h3>
                    </template>
                    <template #content>
                        <DataTable :value="newDeposits">
                            <Column field="name" header="Name"/>
                            <Column field="amount" header="Amount"/>
                            <Column field="status" header="Status"/>
                        </DataTable>
                    </template>
                </Card>
            </div>
        </div>
        <div class="col-span-1 md:col-span-1 lg:col-span-1 xl:col-span-1">
            <div class="col-span-1 md:col-span-1 lg:col-span-1 xl:col-span-1">
                <Card>
                    <template #title>
                        <h3 class="text-xl font-bold">Transactions</h3>
                        <Divider/>
                    </template>
                    <template #content>
                        <Chart type="bar" :data="chartData" :options="chartOptions" />
                    </template>
                </Card>
            </div>
        </div>
    </div>
</template>

<style scoped>
.loading-screen {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    font-size: 1.5rem;
    color: #333;
}

.table-loading-overlay {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    display: flex;
    justify-content: center;
    align-items: center;
    background: rgba(255, 255, 255, 0.8);
}
</style>
