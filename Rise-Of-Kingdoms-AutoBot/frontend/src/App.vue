<template>
  <div class="p-4">
    <button @click="loadScreenshot" class="bg-blue-500 text-white px-4 py-2 rounded">
      Xem ảnh Screenshot
    </button>

    <HomeView
      v-if="imageSrc"
      :image-src="imageSrc"
      :original-width="1280"
      :original-height="720"
    />
  </div>
</template>

<script setup>
import HomeView from './components/HomeView.vue'
import { ref } from 'vue'

const imageSrc = ref(null)

function loadScreenshot() {
  window.go.main.App.Screenshot().then((data) => {
    if (data) {
      imageSrc.value = data
    } else {
      alert("Không thể tải ảnh.")
    }
  })
}
</script>
