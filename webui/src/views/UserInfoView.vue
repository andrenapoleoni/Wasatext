<script>
export default {
    data: function(){
        return{
            userId: localStorage.userId,
            username: localStorage.username,
            photo: localStorage.photo,
            newUsername: "",
            newPhoto: null,
        };

    },

    methods: {
        async updateUsername()
        {
            console.log(this.newUsername);
            try{
                await this.$axios.put(
                     `/user/${sessionStorage.userID}/username`,
                     {username: this.newUsername},
                     {headers: {Authorization: sessionStorage.token}}

                );
                this.username = this.newUsername;
                sessionStorage.username = this.newUsername;
                this.newUsername = "";
                
                
            }catch(error){
                console.log(error);
            }
        },
        
    triggerFileInput() {
      this.$refs.fileInput.click();
    },
    changePhoto(event) {
      const file = event.target.files[0];
      if (file) {
        this.newPhoto = file;
        this.updatePhoto();
      }
    },
    async updatePhoto() {
      const formData = new FormData();
      formData.append('image', this.newPhoto);
      try {
        const response = await this.$axios.put(
          `/user/${sessionStorage.userID}/photo`,
          formData,
          {
            headers: {
              Authorization: sessionStorage.token,
            },
          }
        );
        this.photo = response.data.photo;
        sessionStorage.photo = response.data.photo;
        this.newPhoto = null;
        // Ricarica la pagina dopo aver aggiornato la foto
       
      } catch (error) {
        console.log(error);
      }
    },

},
};




</script>
<template> 
<div class="user-info">
    <img :src="`data:image/jpg;base64,${photo}`" alt="Profile Picture" class="user-photo" />
    <h2>{{ username }}</h2>
    <div class="change-username">
      <input
        v-model="newUsername"
        type="text"
        placeholder="Enter new username"
      />
      <button @click="updateUsername">Change Username</button>
    </div>
    <div class="change-photo">
      <input type="file" ref="fileInput" @change="changePhoto" style="display: none;" />
      <button @click="triggerFileInput">Change Photo</button>
    </div>
  </div>

</template>

<style>
</style>


