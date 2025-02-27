<template>
    <div class="survey-form">
      <el-row>
        <el-col :span="24">
          <h2>创建问卷问题</h2>
          <el-form :model="survey" ref="form" label-width="120px">
            <div v-for="(question, index) in survey.questions" :key="index">
              <el-form-item label="问题" :label-width="'120px'">
                <el-input
                  v-model="question.text"
                  placeholder="请输入问题"
                  clearable
                />
              </el-form-item>
  
              <el-form-item label="问题类型" :label-width="'120px'">
                <el-select v-model="question.type" placeholder="选择问题类型">
                  <el-option label="文本" value="text"></el-option>
                  <el-option label="单选" value="radio"></el-option>
                  <el-option label="多选" value="checkbox"></el-option>
                </el-select>
              </el-form-item>
  
              <div v-if="question.type === 'radio' || question.type === 'checkbox'">
                <el-form-item label="选项" :label-width="'120px'">
                  <div v-for="(option, oIndex) in question.options" :key="oIndex">
                    <el-input
                      v-model="question.options[oIndex]"
                      placeholder="选项"
                      clearable
                    />
                  </div>
                  <el-button type="primary" @click="addOption(index)">添加选项</el-button>
                </el-form-item>
              </div>
  
              <el-button type="danger" @click="removeQuestion(index)">删除问题</el-button>
            </div>
  
            <el-button type="primary" @click="addQuestion">添加问题</el-button>
            <el-button type="success" @click="saveSurvey">保存问卷</el-button>
          </el-form>
        </el-col>
      </el-row>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        survey: {
          questions: []
        }
      };
    },
    methods: {
      addQuestion() {
        this.survey.questions.push({
          text: '',
          type: 'text',
          options: [''] // 默认有一个选项
        });
      },
      removeQuestion(index) {
        this.survey.questions.splice(index, 1);
      },
      addOption(index) {
        this.survey.questions[index].options.push('');
      },
      saveSurvey() {
        console.log(this.survey.questions);
        // 你可以将数据发送到后端进行保存
      }
    }
  };
  </script>
  
  <style scoped>
  .survey-form {
    margin: 20px;
  }
  </style>
  