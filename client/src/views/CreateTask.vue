<template>
<div class="create-task is-clipped">
  <ModalView>
  <div class="card is-v-center">
    <div class="card-content">
      <h1 class="title">Create Task</h1>

      <!-- form elements -->
      <form id="create-task" @submit.prevent="createTask">
        <!-- name -->
        <div class="field">
          <label class="label">Name</label>
          <div class="control">
            <input type="text" class="input" v-model="name" required placeholder="Eat Sushi">
          </div>
        </div>

        <!-- duration -->
        <div class="field">
          <label class="label">Duration</label>
          <div class="control">
            <input type="number" class="input" v-model="duration" required placeholder="0 MIN">
          </div>
        </div>

      <!--tags -->
      <div class="field">
        <label class="label">Tags</label>
        <div class="control">
          <multiselect 
          v-model="selectedTags" track-by="_id" label="name" 
          :options="tagOptions" :allow-empty="true"
          :multiple="true" :taggable="true"
          @tag="createTag"
          placeholder="Select tags">
            <template slot="singleLabel" slot-scope="{ option }">{{ option.name }}</template>
          </multiselect>   
        </div>   
      </div>

        <!--highlight  -->
        <div class="field" id="checkbox">
          <label class="checkbox">
            <input type="checkbox" v-model="highlight">
            Highlight of the Day
          </label>
        </div>

        <!-- submit -->
        <div class="field is-grouped">
          <div class="control">
            <button class="button is-primary" :class="{'is-loading':isLoading}">Create</button>
          </div>
          <div class="control">
            <button class="button is-primary is-light" @click="$router.push({name:'Tasks'})">Back</button>
          </div>
        </div>

      </form>
    </div>
  </div> 
  </ModalView>
</div>  
</template>

<script>
import ModalView from '@/components/ModalView'
import {mapActions} from 'vuex'
import tags from '@/mixins/tags'

import Multiselect from 'vue-multiselect'

import postTask from '@/api/taskPost'

export default {
  name:'CreateTask',
  components:{
    ModalView,
    Multiselect
  },
  mixins:[tags],
  data(){
    return{
      name:'',
      duration:'',
      highlight:false,
      isLoading:false,
      selectedTags:[]
    }
  },
  methods:{
    createTask:async function(){

      let tagIds

      // only get tagIds
      if(this.selectedTags.length){
        tagIds = this.selectedTags.map(tag=>{return tag._id})
      }

      let task = {
        name:this.name,
        duration:parseInt(this.duration) || 0,
        highlight:this.highlight,
        tag_ids:tagIds
      }

      // check if task is for tomorrow
      if(this.$route.query.range=='tomorrow'){
        task.is_tomorrow = true
      }

      this.isLoading = true
      const savedTask = await postTask(task)
      this.isLoading = false
      // save to store
      this.storeTask(savedTask)
    },
    ...mapActions({
      storeTask:'task/storeTask'
    })
  }
}
</script>

<style src="vue-multiselect/dist/vue-multiselect.min.css"></style>

<style lang="scss" scoped>
@import '~bulma';

.card{
  width: 30vw!important;
}

.field:last-child{
  margin-top: 2rem;
}
</style>