<template>
  <svg :class="className" :style="{ width, height: _height }">
    <use :xlink:href="iconName"></use>
  </svg>
</template>

<script setup lang="ts">
import { computed } from "vue";
const props = defineProps({
  name: { type: String, require: true },
  width: { type: String, default: "16px" },
  height: { type: String },
  className: { type: String },
});
const _height = computed(() => {
  return props.height ? props.height : props.width;
});
const iconName = computed(() => {
  return `#${props.name}`;
});
</script>

<style lang="scss" scoped></style>