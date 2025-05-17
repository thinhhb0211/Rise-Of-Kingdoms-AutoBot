<template>
  <div>
    <h3>Config</h3>

    <label><input type="checkbox" v-model="takeBreak" /> Take break at every end of round</label>
    <div v-if="takeBreak" class="mb-4">
      Execute at every <input type="number" v-model.number="breakRound" min="1" /> round
      <br />
      <label><input type="checkbox" v-model="terminate" /> Terminate when break</label>
      <select v-model="breakTime">
        <option>3 Minute</option>
        <option>5 Minute</option>
        <option>10 Minute</option>
      </select>
    </div>

    <label><input type="checkbox" v-model="openChest" /> Open free chest in tavern</label>

    <label><input type="checkbox" v-model="autoCollectResources" /> Auto collect resources</label>

    <label><input type="checkbox" v-model="autoAttack" /> Auto attack enemies</label>
    <div v-if="autoAttack" class="mb-4">
      Attack power level: <input type="number" v-model.number="attackPower" min="1" max="100" />
      <br />
      <label><input type="checkbox" v-model="useSpecialSkill" /> Use special skill</label>
    </div>

    <label><input type="checkbox" v-model="autoHeal" /> Auto heal when low HP</label>
    <div v-if="autoHeal" class="mb-4">
      Heal threshold: <input type="number" v-model.number="healThreshold" min="1" max="100" /> %
    </div>

    <label><input type="checkbox" v-model="enableNotifications" /> Enable notifications</label>

    <label><input type="checkbox" v-model="autoUpgradeBuildings" /> Auto upgrade buildings</label>

    <button @click="$emit('editPositions')" class="mt-4 px-4 py-2 bg-blue-600 text-white rounded">
      Edit Positions
    </button>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const takeBreak = ref(true)
const breakRound = ref(1)
const terminate = ref(false)
const breakTime = ref('3 Minute')
const openChest = ref(true)

const autoCollectResources = ref(false)

const autoAttack = ref(false)
const attackPower = ref(50)
const useSpecialSkill = ref(false)

const autoHeal = ref(false)
const healThreshold = ref(30)

const enableNotifications = ref(false)

const autoUpgradeBuildings = ref(false)

// Bạn có thể emit hoặc dùng watch để gửi config lên cha hoặc backend
</script>
