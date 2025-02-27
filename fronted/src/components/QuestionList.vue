<template>
    <div class="question-list">
      <el-row>
        <el-col :span="24">
          <h2>问卷问题列表</h2>
          <div v-for="(question, index) in questions" :key="index" class="question">
            <el-card>
              <h3>{{ question.text }}</h3>
              <div v-if="question.type === 'radio' || question.type === 'checkbox'">
                <el-list>
                  <el-list-item v-for="(option, oIndex) in question.options" :key="oIndex">
                    {{ option }}
                  </el-list-item>
                </el-list>
              </div>
            </el-card>
          </div>
        </el-col>
      </el-row>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      questions: Array
    }
  };
  </script>
  
  <style scoped>
  .question-list {
    margin: 20px;
  }
  </style>
  