<template>
  <div class="mb-6">
    <label class="label">Update Password</label>
    <label class="sub-label">Current Password</label>
    <input
      type="text"
      class="input-field"
      placeholder="New Name"
      required
      name="taskName"
    />
    <label class="sub-label">New Password</label>
    <div class="flex">
      <input
        type="text"
        class="input-field"
        placeholder="New Name"
        required
        name="taskName"
      />
      <button
        type="submit"
        class="
          inline-flex
          text-white
          font-medium
          rounded-lg
          text-sm
          w-full
          sm:w-auto
          ml-4
          px-5
          py-2.5
          text-center
          bg-blue-600
          hover:bg-blue-700
          focus:ring-blue-800
        "
      >
        Update
      </button>
    </div>
  </div>
</template>
