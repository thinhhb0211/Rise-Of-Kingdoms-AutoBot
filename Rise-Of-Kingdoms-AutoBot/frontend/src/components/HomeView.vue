<script setup>
import { ref, reactive } from 'vue'

const props = defineProps({
  imageSrc: String,
  originalWidth: Number,
  originalHeight: Number
})

const showPositionSelector = ref(true)  // true = màn hình chọn vị trí

const selected = ref(null)
const imageRef = ref(null)

const buildings = reactive({
  "Nhà Chính": [-1, -1],
  "Bệnh viện": [-1, -1],
  "Trại": [-1, -1],
  "Công xưởng": [-1, -1],
  "Trại lính": [-1, -1],
  "Tháp canh": [-1, -1],
  "Kết giao": [-1, -1],
  "Trạm Trinh sát": [-1, -1],
  "Bảng Thông Báo": [-1, -1],
})

// Config reactive object
const config = reactive({
  takeBreakEveryEndRound: true,
  executeEveryBreak: 1,
  terminateWhenBreak: false,
  breakDuration: '3 Minute',

  useResourceBuyMysteryMerchant: false,
  openFreeChestInTavern: true,
  collectResourceTroopsHelpAlliance: true,
  produceMaterial: true,
  executeEveryProduceMaterial: 1,
  claimDailyVipPointChest: true,
  executeEveryVip: 1,
  claimQuestsDailyObjectives: true,
  executeEveryQuests: 1,
  collectAlliedResourceGiftsDonateTech: true,
  executeEveryCollectAllied: 1,

  autoUpgradeTrainTroops: true,
  troopsLevels: {
    barracksTrainingLv: 1,
    barracksUpgradeLv: 1,
    archeryTrainingLv: 1,
    archeryUpgradeLv: 1,
    stableTrainingLv: 1,
    stableUpgradeLv: 1,
    siegeTrainingLv: 1,
    siegeUpgradeLv: 1,
  },

  attackBarbarians: false,
})

function onImageClick(event) {
  if (!selected.value) {
    alert("Chọn công trình trước")
    return
  }

  const img = imageRef.value
  const rect = img.getBoundingClientRect()
  
  const xClicked = event.clientX - rect.left
  const yClicked = event.clientY - rect.top

  const scaleX = props.originalWidth / rect.width
  const scaleY = props.originalHeight / rect.height

  const xOriginal = Math.floor(xClicked * scaleX)
  const yOriginal = Math.floor(yClicked * scaleY)

  buildings[selected.value] = [xOriginal, yOriginal]
}

async function savePositions() {
  try {
    await window.go.main.App.SaveBuildingPositionsJSON(buildings)
    alert("Đã lưu vị trí công trình!")
    showPositionSelector.value = false // chuyển sang màn hình config
  } catch (err) {
    alert("Lỗi lưu vị trí: " + err)
  }
}

async function saveConfig() {
  try {
    await window.go.main.App.SaveConfigJSON(config)
    alert("Đã lưu cấu hình!")
  } catch (err) {
    alert("Lỗi lưu cấu hình: " + err)
  }
}

function editPositions() {
  showPositionSelector.value = true
}
</script>

<template>
  <div v-if="showPositionSelector" class="flex mt-4 gap-4">
    <div class="w-1/3 p-4">
      <h2 class="font-bold mb-2">Building Position</h2>
      <ul>
        <li
          v-for="(pos, name) in buildings"
          :key="name"
          @click="selected = name"
          :class="['cursor-pointer p-1', selected === name ? 'bg-blue-200' : '']"
        >
          {{ name }}: [{{ pos[0] }}, {{ pos[1] }}]
        </li>
      </ul>
      <button
        class="mt-4 px-4 py-2 bg-blue-600 text-white rounded"
        @click="savePositions"
      >
        Save Positions
      </button>
    </div>

    <div class="flex-1 p-4">
      <img
        ref="imageRef"
        :src="props.imageSrc"
        class="border max-w-full h-auto"
        @click="onImageClick"
      />
    </div>
  </div>

  <!-- Màn hình config chi tiết -->
  <div v-else class="p-4 max-w-3xl mx-auto">
    <h2 class="font-bold mb-4">Config</h2>

    <label class="block mb-2">
      <input type="checkbox" v-model="config.takeBreakEveryEndRound" />
      Take break at every end of round
    </label>

    <div class="mb-4">
      Execute at every
      <input
        type="number"
        v-model.number="config.executeEveryBreak"
        min="1"
        class="border px-2 py-1 w-16 mx-2"
      />
      round
    </div>

    <label class="block mb-4">
      <input type="checkbox" v-model="config.terminateWhenBreak" />
      Terminate when break
    </label>

    <label class="block mb-6">
      Break duration:
      <select v-model="config.breakDuration" class="border px-2 py-1 ml-2">
        <option>3 Minute</option>
        <option>5 Minute</option>
        <option>10 Minute</option>
      </select>
    </label>

    <label class="block mb-2">
      <input type="checkbox" v-model="config.useResourceBuyMysteryMerchant" />
      Use resource buy item in Mystery Merchant
    </label>

    <label class="block mb-2">
      <input type="checkbox" v-model="config.openFreeChestInTavern" />
      Open free chest in tavern
    </label>

    <label class="block mb-2">
      <input type="checkbox" v-model="config.collectResourceTroopsHelpAlliance" />
      Collecting resource, troops, and help alliance
    </label>

    <label class="block mb-2">
      <input type="checkbox" v-model="config.produceMaterial" />
      Produce material
    </label>

    <div class="mb-4 ml-6">
      Execute at every
      <input
        type="number"
        v-model.number="config.executeEveryProduceMaterial"
        min="1"
        class="border px-2 py-1 w-16 mx-2"
      />
      round
    </div>

    <label class="block mb-2">
      <input type="checkbox" v-model="config.claimDailyVipPointChest" />
      Claim daily vip point and chest
    </label>

    <div class="mb-4 ml-6">
      Execute at every
      <input
        type="number"
        v-model.number="config.executeEveryVip"
        min="1"
        class="border px-2 py-1 w-16 mx-2"
      />
      round
    </div>

    <label class="block mb-2">
      <input type="checkbox" v-model="config.claimQuestsDailyObjectives" />
      Claim quests and daily objectives
    </label>

    <div class="mb-4 ml-6">
      Execute at every
      <input
        type="number"
        v-model.number="config.executeEveryQuests"
        min="1"
        class="border px-2 py-1 w-16 mx-2"
      />
      round
    </div>

    <label class="block mb-2">
      <input type="checkbox" v-model="config.collectAlliedResourceGiftsDonateTech" />
      Collecting allied resource, gifts, and donate technology
    </label>

    <div class="mb-4 ml-6">
      Execute at every
      <input
        type="number"
        v-model.number="config.executeEveryCollectAllied"
        min="1"
        class="border px-2 py-1 w-16 mx-2"
      />
      round
    </div>

    <label class="block mb-2">
      <input type="checkbox" v-model="config.autoUpgradeTrainTroops" />
      Auto upgrade and train troops
    </label>

    <div class="grid grid-cols-2 gap-4 mb-4 ml-6">
      <div>
        Barracks: Training Lv.
        <input
          type="number"
          v-model.number="config.troopsLevels.barracksTrainingLv"
          min="1"
          class="border px-2 py-1 w-16 mx-2"
        />
      </div>
      <div>
        Upgrade Lv.
        <input
          type="number"
          v-model.number="config.troopsLevels.barracksUpgradeLv"
          min="1"
          class="border px-2 py-1 w-16 mx-2"
        />
      </div>

      <div>
        Archery: Training Lv.
        <input
          type="number"
          v-model.number="config.troopsLevels.archeryTrainingLv"
          min="1"
          class="border px-2 py-1 w-16 mx-2"
        />
      </div>
      <div>
        Upgrade Lv.
        <input
          type="number"
          v-model.number="config.troopsLevels.archeryUpgradeLv"
          min="1"
          class="border px-2 py-1 w-16 mx-2"
        />
      </div>

      <div>
        Stable: Training Lv.
        <input
          type="number"
          v-model.number="config.troopsLevels.stableTrainingLv"
          min="1"
          class="border px-2 py-1 w-16 mx-2"
        />
      </div>
      <div>
        Upgrade Lv.
        <input
          type="number"
          v-model.number="config.troopsLevels.stableUpgradeLv"
          min="1"
          class="border px-2 py-1 w-16 mx-2"
        />
      </div>

      <div>
        Siege: Training Lv.
        <input
          type="number"
          v-model.number="config.troopsLevels.siegeTrainingLv"
          min="1"
          class="border px-2 py-1 w-16 mx-2"
        />
      </div>
      <div>
        Upgrade Lv.
        <input
          type="number"
          v-model.number="config.troopsLevels.siegeUpgradeLv"
          min="1"
          class="border px-2 py-1 w-16 mx-2"
        />
      </div>
    </div>

    <label class="block mb-6">
      <input type="checkbox" v-model="config.attackBarbarians" />
      Attack Barbarians
    </label>

    <div class="flex gap-4">
      <button
        class="px-4 py-2 bg-green-600 text-white rounded"
        @click="saveConfig"
      >
        Save Config
      </button>

      <button
        class="px-4 py-2 bg-gray-600 text-white rounded"
        @click="editPositions"
      >
        Edit Positions
      </button>
    </div>
  </div>
</template>
