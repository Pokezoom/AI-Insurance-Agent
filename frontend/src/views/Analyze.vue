<template>
  <div class="analyze-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>上传保单图片</span>
        </div>
      </template>

      <el-upload
        class="upload-demo"
        drag
        :auto-upload="false"
        :on-change="handleFileChange"
        :limit="1"
        accept="image/*"
      >
        <el-icon class="el-icon--upload"><upload-filled /></el-icon>
        <div class="el-upload__text">
          将保单图片拖到此处，或<em>点击上传</em>
        </div>
        <template #tip>
          <div class="el-upload__tip">
            支持 jpg/png 格式，文件大小不超过 10MB
          </div>
        </template>
      </el-upload>

      <div v-if="imagePreview" class="preview-section">
        <el-divider>图片预览</el-divider>
        <el-image :src="imagePreview" fit="contain" style="max-height: 400px" />
      </div>

      <div class="action-section">
        <el-button type="primary" size="large" :loading="analyzing" :disabled="!imageFile" @click="handleAnalyze">
          开始分析
        </el-button>
        <el-button size="large" @click="handleReset">重置</el-button>
      </div>
    </el-card>

    <el-card v-if="result" class="result-card">
      <template #header>
        <div class="card-header">
          <span>分析结果</span>
          <el-tag :type="result.status === 'success' ? 'success' : 'danger'">
            {{ result.status === 'success' ? '成功' : '失败' }}
          </el-tag>
        </div>
      </template>

      <div v-if="result.status === 'success'">
        <el-descriptions title="保单信息" :column="2" border>
          <el-descriptions-item label="记录ID">{{ result.id }}</el-descriptions-item>
          <el-descriptions-item label="分析时间">{{ result.created_at }}</el-descriptions-item>
        </el-descriptions>

        <el-divider>分析详情</el-divider>
        <div class="analysis-content">
          <pre>{{ result.analysis_result }}</pre>
        </div>
      </div>

      <el-alert v-else type="error" :title="result.error_message" :closable="false" />
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { UploadFilled } from '@element-plus/icons-vue'
import { analyzePolicy } from '@/api/policy'

const imageFile = ref(null)
const imagePreview = ref('')
const analyzing = ref(false)
const result = ref(null)

const handleFileChange = (file) => {
  imageFile.value = file.raw
  imagePreview.value = URL.createObjectURL(file.raw)
  result.value = null
}

const handleAnalyze = async () => {
  if (!imageFile.value) {
    ElMessage.warning('请先上传保单图片')
    return
  }

  analyzing.value = true
  try {
    const reader = new FileReader()
    reader.readAsDataURL(imageFile.value)
    reader.onload = async () => {
      const base64 = reader.result.split(',')[1]
      const res = await analyzePolicy({
        image_base64: base64,
        image_type: imageFile.value.type
      })
      result.value = res.data
      ElMessage.success('分析完成')
    }
  } catch (error) {
    ElMessage.error('分析失败')
  } finally {
    analyzing.value = false
  }
}

const handleReset = () => {
  imageFile.value = null
  imagePreview.value = ''
  result.value = null
}
</script>

<style scoped>
.analyze-container {
  max-width: 1200px;
  margin: 0 auto;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.preview-section {
  margin-top: 20px;
  text-align: center;
}

.action-section {
  margin-top: 20px;
  text-align: center;
}

.result-card {
  margin-top: 20px;
}

.analysis-content {
  background-color: #f5f7fa;
  padding: 20px;
  border-radius: 4px;
}

.analysis-content pre {
  white-space: pre-wrap;
  word-wrap: break-word;
  margin: 0;
  font-family: inherit;
  line-height: 1.8;
}
</style>
