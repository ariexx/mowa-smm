<script setup>
import { computed } from 'vue';
import { useUserRole } from '../composables/useUserRole';
import AppMenuItem from './AppMenuItem.vue';

const { userRole } = useUserRole();

const model = computed(() => {
    let defaultMenu = [
       {
        label: 'Main Menu',
        items: [
            { label: 'Dashboard', icon: 'pi pi-fw pi-circle-fill', to: '/' },
            {
            label: 'Buat Pesanan',
            icon: 'pi pi-fw pi-cart-plus',
            items: [
                {
                    label: 'Social Media Single',
                    to: '/order/social-media'
                },
                {
                    label: 'Social Media Mass',
                    to: '/order/social-media-massal'
                },
            ]
        }
        ]
       }
    ];
    if (userRole.value === 'administrator') {
        return defaultMenu
    } else if (userRole.value === 'member') {
        return [
            {
                label: 'Member Dashboard',
                items: [{ label: 'Profile', icon: 'pi pi-fw pi-user', to: '/profile' }],
            },
        ];
    } else {
        return [];
    }
});
</script>

<template>
    <ul class="layout-menu">
        <template v-for="(item, i) in model" :key="i">
            <app-menu-item v-if="!item.separator" :item="item" :index="i"></app-menu-item>
            <li v-if="item.separator" class="menu-separator"></li>
        </template>
    </ul>
</template>
