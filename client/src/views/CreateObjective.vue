<template>
<div class="create-objective is-clipped">
  <ModalView>
  <div class="card is-v-center">
    <div class="card-content">
      <h1 class="title">Create Objective</h1>

      <!-- form elements -->
      <form id="create-task" @submit.prevent="createObjective">
        <!-- name -->
        <div class="field">
          <label class="label">Name</label>
          <div class="control">
            <input type="text" class="input" v-model="name" required placeholder="Eat Sushi">
          </div>
        </div>

        <!-- duration -->
        <div class="field">
          <label class="label">Scheduled Date</label>
          <div class="control">
            <datetime v-model="scheduledFor" type="date" class="theme-dark"></datetime>
          </div>
        </div>

        <!--body-->
        <div class="field">
          <label class="label">Body</label>
          <div class="control">
            <textarea class="textarea" placeholder="Textarea" v-model="body"></textarea>
          </div>
        </div>

        <!-- submit -->
        <div class="field is-grouped">
          <div class="control">
            <button class="button is-primary" :class="{'is-loading':isLoading}">Create</button>
          </div>
          <div class="control">
            <button class="button is-primary is-light" @click="$router.push({name:'Objectives'})">Back</button>
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

import {Datetime} from 'vue-datetime'
import 'vue-datetime/dist/vue-datetime.css'

export default {
  name:'CreateObjective',
  components:{
    ModalView,
    Datetime,
  },
  data(){
    return {
      name:'',
      scheduledFor:'',
      completed:false,
      body:'',
      isLoading:false
    }
  }
}
</script>

<style lang="scss" scoped>
@import '~bulma';

.card{
  width: 30vw!important;
}

.field:last-child{
  margin-top: 2rem;
}

// datetime theme
/deep/ .theme-dark .vdatetime-popup__header,
/deep/.theme-dark .vdatetime-calendar__month__day--selected > span > span,
/deep/ .theme-dark .vdatetime-calendar__month__day--selected:hover > span > span {
  background: #FF9800;
}

/deep/ .theme-dark .vdatetime-year-picker__item--selected,
/deep/ .theme-dark .vdatetime-time-picker__item--selected,
/deep/ .theme-dark .vdatetime-popup__actions__button {
  color: #ff9800;
}
</style>