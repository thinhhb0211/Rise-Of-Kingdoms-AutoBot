<template>
    <div class="p-4 space-y-4">
        <!-- Window Size -->
        <div>
            <label class="block font-bold">Window Size (Restart to Effect)</label>
            <select v-model="screenSize" @change="updateScreenSize">
                <option v-for="size in windowSizes" :key="size" :value="size">{{ size }}</option>
            </select>
        </div>

        <!-- Method Selection -->
        <div>
            <label class="block font-bold">Pass Verification Method</label>
            <select v-model="config.method" @change="updateConfig">
                <option value="none">None</option>
                <option value="twocaptcha">2Captcha</option>
                <option value="haoi">Haoi</option>
            </select>
        </div>

        <!-- 2Captcha Config -->
        <div v-if="config.method === 'twocaptcha'" class="border p-2 rounded">
            <label>2Captcha Key</label>
            <input type="text" v-model="config.twocaptchaKey" @input="updateConfig" />
        </div>

        <!-- Haoi Config -->
        <div v-if="config.method === 'haoi'" class="border p-2 rounded">
            <label>User Key</label>
            <input type="text" v-model="config.haoiUser" @input="updateConfig" />
            <label>Software Key</label>
            <input type="text" v-model="config.haoiRebate" @input="updateConfig" />
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { LoadConfig, SetConfig } from '../../wailsjs/go/main/App'

const config = ref({
    method: '',
    haoiUser: '',
    haoiRebate: '',
    twocaptchaKey: '',
    screenSize: []
})

const screenSize = ref('470 x 450')
const windowSizes = ['470 x 450', '470 x 550', '470 x 650', '470 x 750', '470 x 850']

onMounted(async () => {
    config.value = await LoadConfig()
    screenSize.value = `${config.value.screenSize[0]} x ${config.value.screenSize[1]}`
})

function updateConfig() {
    SetConfig(config.value)
}

function updateScreenSize() {
    const [w, h] = screenSize.value.split('x').map(s => parseInt(s.trim()))
    config.value.screenSize = [w, h]
    updateConfig()
}
</script>