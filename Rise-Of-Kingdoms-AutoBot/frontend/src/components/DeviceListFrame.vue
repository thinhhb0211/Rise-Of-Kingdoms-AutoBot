<template>
    <div>
        <div>
            <input v-model="name" placeholder="name" />
            <input v-model="ip" placeholder="ip" />
            <input v-model="port" placeholder="port" />
            <button @click="addDevice">Add</button>
        </div>

        <div v-for="device in devices" :key="device.name" class="device-row">
            <span>{{ device.name }}</span>
            <span>{{ device.ip }}:{{ device.port }}</span>
            <span>connected</span>
            <button @click="displayDevice(device)">Display</button>
            <button @click="deleteDevice(device.name)">Delete</button>
        </div>

        <p class="footer">
            Welcome to use Rise of Kingdoms Bot, check updates on
            <a href="https://github.com">GitHub</a>
        </p>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { AddDevice, GetDevices, DeleteDevice } from '../../wailsjs/go/main/App'

const name = ref('')
const ip = ref('')
const port = ref('')
const devices = ref([])

const loadDevices = async () => {
    devices.value = await GetDevices()
}

const addDevice = async () => {
    devices.value = await AddDevice(name.value, ip.value, port.value)
}

const deleteDevice = async (name) => {
    devices.value = await DeleteDevice(name)
}

const displayDevice = (device) => {
    alert("Display device: " + device.name)
}

onMounted(loadDevices)
</script>