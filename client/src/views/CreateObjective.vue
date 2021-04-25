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
            <input type="number" class="input" v-model="duration" required placeholder="0 MIN">
          </div>
        </div>

        <!--body-->
        <div class="field">
          <label class="label">Body</label>
          <div class="control">
            <textarea class="textarea" placeholder="Textarea"></textarea>
          </div>
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

export default {
  name:'CreateObjective',
  components:{
    ModalView
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
</style>