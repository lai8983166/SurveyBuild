<template>
  <div id="app">
    <!-- 引入并使用  组件 -->
     <QuestionForm />
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import QuestionForm from './components/QuestionForm.vue'
export default{
  name:'App',
  components:{
    QuestionForm,
  },
}
</script>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
