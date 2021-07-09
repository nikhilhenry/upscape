import Vue from "vue";

export default {
  namespaced: true,

  state: {
    objectives: [],
    objectivesLoaded: false,
  },

  mutations: {
    STORE_OBJECTIVES: (state, objectivesList) => {
      const objectives = state.objectives.concat(objectivesList);
      state.objectives = objectives;
    },
    SET_OBJECTIVES: (state, objectivesList) => {
      state.objectives = objectivesList;
    },
    SET_OBJECTIVES_LOADED: (state, isLoaded) => {
      state.objectivesLoaded = isLoaded;
    },
    UPDATE_OBJECTIVE: (state, payload) => {
      Vue.set(state.objectives, payload.index, payload.updatedObjective);
    },
    STORE_OBJECTIVE: (state, newObjective) => {
      const oldObjectives = state.objectives;
      state.objectives = [newObjective, ...oldObjectives];
    },
    DELETE_OBJECTIVE: (state, objectiveIndex) => {
      state.objectives.splice(objectiveIndex, 1);
    },
  },

  actions: {
    storeObjectives: (context, userTasks) => {
      context.commit("STORE_OBJECTIVES", userTasks);
      if (userTasks.length) {
        context.commit("SET_OBJECTIVES_LOADED", true);
      }
    },
    setObjectives: (context, userTasks) => {
      context.commit("SET_OBJECTIVES", userTasks);
    },

    updateObjective: (context, updatedObjective) => {
      if (updatedObjective) {
        const index = context.state.objectives.findIndex(
          (objective) => objective._id === updatedObjective._id
        );
        context.commit("UPDATE_OBJECTIVE", { index, updatedObjective });
      }
    },

    storeObjective: (context, newObjective) => {
      if (newObjective) {
        context.commit("STORE_OBJECTIVE", newObjective);
      }
    },

    deleteObjective: (context, objectiveId) => {
      if (objectiveId) {
        const index = context.state.objectives.findIndex(
          (objective) => objective._id === objectiveId
        );
        context.commit("DELETE_OBJECTIVE", index);
      }
    },
  },

  getters: {
    getObjectives: (state) => {
      return state.objectives;
    },
    getObjective: (state) => (objectiveId) => {
      const objective = state.objectives.find(
        (objective) => objective._id === objectiveId
      );
      return objective;
    },
  },
};
