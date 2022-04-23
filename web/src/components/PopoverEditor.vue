<script setup>
import { Check, Close } from '@element-plus/icons-vue'
import { ref, unref } from 'vue';
import { ElMessage } from 'element-plus'
import axios from '@/utils/request'

const props = defineProps({
  name: {
    type: String,
    required: true
  },
  text: {
    required: true
  },
  type: {
    type: String,
    default: "string"
  },
  url: {
    type: String,
    required: true
  }
})
const emit = defineEmits(['saved'])

const popoverRef = ref()
const input = ref(props.text)

const hide = () => {
  unref(popoverRef).hide()
}

const submit = async () => {
  let val = input.value
  if (props.type === "number") {
    val = +val
  }
  await axios.put(props.url, { [props.name]: val })
  ElMessage.success('保存成功')
  hide()
  emit('saved', input.value)
}
</script>

<template>
  <el-popover placement="right" trigger="click" width="auto" ref="popoverRef">
    <template #reference>
      <a class="popover__reference" href="javascript:void(0)">{{ text }}</a>
    </template>
    <div class="popover__content">
      <el-input v-model="input" />
      <el-button type="primary" :icon="Check" size="small" @click="submit" />
      <el-button :icon="Close" size="small" @click="hide" />
    </div>
  </el-popover>
</template>

<style lang="scss" scoped>
.popover {
  &__reference {    
    color: #333;
    text-decoration: underline dotted var(--el-color-primary-light-1);
  }
  &__content {
    display: flex;
    flex-direction: row;
    button {
      width: 40px;
      height: 32px;
      + button {
        margin-left: 2px;
      }
    }
  }
}
</style>