<script setup lang="ts">
import { ref } from "vue";
import config from "@/config"

const imgUrl = ref("");
const loading = ref(true);

async function getImage() {
  const res = await fetch(
    "https://api.waifu.im/search?&height=>=2000",
  );
  const data = await res.json();
  imgUrl.value = data.images[0].url;
  loading.value = false;
  console.log(data);
}
getImage();
</script>

<template>
  <div >
    <h1>This is an about page</h1>
	<p>My envs : {{config.apiUrl}}</p>
    <div v-if="loading">---0---</div>
    <img class="myImage" :src="imgUrl" v-else />
  </div>
</template>

<style>
.myImage {
  display: block;
  width: 10rem;
}

@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
