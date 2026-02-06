<template>
  <div class="records-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>分析历史记录</span>
          <el-button type="primary" @click="loadRecords">刷新</el-button>
        </div>
      </template>

      <el-table :data="records" v-loading="loading" stripe>
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="image_type" label="图片类型" width="120" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 'success' ? 'success' : row.status === 'failed' ? 'danger' : 'info'">
              {{ row.status === 'success' ? '成功' : row.status === 'failed' ? '失败' : '处理中' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="分析时间" width="180" />
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button size="small" @click="viewDetail(row.id)">查看详情</el-button>
            <el-button size="small" type="danger" @click="handleDelete(row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :total="total"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="loadRecords"
        @current-change="loadRecords"
        style="margin-top: 20px; justify-content: center"
      />
    </el-card>

    <el-dialog v-model="detailVisible" title="分析详情" width="800px">
      <div v-if="currentRecord" v-loading="detailLoading">
        <el-descriptions :column="2" border>
          <el-descriptions-item label="记录ID">{{ currentRecord.id }}</el-descriptions-item>
          <el-descriptions-item label="用户ID">{{ currentRecord.user_id }}</el-descriptions-item>
          <el-descriptions-item label="图片类型">{{ currentRecord.image_type }}</el-descriptions-item>
          <el-descriptions-item label="状态">
            <el-tag :type="currentRecord.status === 'success' ? 'success' : 'danger'">
              {{ currentRecord.status === 'success' ? '成功' : '失败' }}
            </el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="分析时间" :span="2">{{ currentRecord.created_at }}</el-descriptions-item>
        </el-descriptions>

        <el-divider>分析结果</el-divider>
        <div v-if="currentRecord.status === 'success'" class="analysis-content">
          <pre>{{ currentRecord.analysis_result }}</pre>
        </div>
        <el-alert v-else type="error" :title="currentRecord.error_message" :closable="false" />
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getRecords, getRecordDetail, deleteRecord } from '@/api/policy'

const loading = ref(false)
const records = ref([])
const currentPage = ref(1)
const pageSize = ref(20)
const total = ref(0)

const detailVisible = ref(false)
const detailLoading = ref(false)
const currentRecord = ref(null)

const loadRecords = async () => {
  loading.value = true
  try {
    const res = await getRecords({
      page: currentPage.value,
      page_size: pageSize.value
    })
    records.value = res.data.records
    total.value = res.data.total
  } finally {
    loading.value = false
  }
}

const viewDetail = async (id) => {
  detailVisible.value = true
  detailLoading.value = true
  try {
    const res = await getRecordDetail(id)
    currentRecord.value = res.data
  } finally {
    detailLoading.value = false
  }
}

const handleDelete = (id) => {
  ElMessageBox.confirm('确定要删除这条记录吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    await deleteRecord(id)
    ElMessage.success('删除成功')
    loadRecords()
  })
}

onMounted(() => {
  loadRecords()
})
</script>

<style scoped>
.records-container {
  max-width: 1400px;
  margin: 0 auto;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.analysis-content {
  background-color: #f5f7fa;
  padding: 20px;
  border-radius: 4px;
  max-height: 400px;
  overflow-y: auto;
}

.analysis-content pre {
  white-space: pre-wrap;
  word-wrap: break-word;
  margin: 0;
  font-family: inherit;
  line-height: 1.8;
}
</style>
