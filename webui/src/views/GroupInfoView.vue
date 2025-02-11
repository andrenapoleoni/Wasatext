<script>
export default{
    data: function(){
        return{
            groupId: localStorage.groupId,
            groupName: localStorage.groupName,
            groupPhoto: localStorage.groupPhoto,
            newGroupName: "",
            newGroupPhoto: null,
        };
    },

    methods:{
        async updateGroupName()
        {
            console.log(this.newGroupName);
            try{
                await this.$axios.put(
                    `user/${sessionStorage.userID}/groups/${localStorage.groupId}/groupname`,
                    {groupName: this.newGroupName},
                    {headers: {Authorization: sessionStorage.token}}
                );
                this.groupName = this.newGroupName;
                localStorage.groupName = this.newGroupName;
                this.newGroupName = "";
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
                this.newGroupPhoto = file;
                this.updatePhoto();
            }
        },
        async updatePhoto() {
            const formData = new FormData();
            formData.append('image', this.newGroupPhoto);
            try {
                const response = await this.$axios.put(
                    `user/${sessionStorage.userID}/groups/${localStorage.groupId}/groupphoto`,
                    formData,
                    {
                        headers: {
                            Authorization: sessionStorage.token,
                        },
                    }
                );
                this.groupPhoto = response.data.photo;
                localStorage.groupPhoto = response.data.photo;
                this.newGroupPhoto = null;
            }catch(error){
                console.log(error);
            }
        },
    }
}



</script>
<template> 
    <div class="container">
        <div class="row">
            <div class="col-12">
                <h1>Group Info</h1>
            </div>
        </div>
        <div class="row">
            <div class="col-12">
                <img :src="'data:image/jpg;base64,' + groupPhoto" class="user-avatar" />    
            </div>
        </div>
        <div class="row">
            <div class="col-12">
                <h2>{{groupName}}</h2>
            </div>
        </div>
        <div class="row">
            <div class="col-12">
                <button class="btn btn-primary" @click="triggerFileInput">Change Photo</button>
                <input type="file" ref="fileInput" @change="changePhoto" style="display: none">
            </div>
        </div>
        <div class="row">
            <div class="col-12">
                <input type="text" v-model="newGroupName" class="form-control">
                <button class="btn btn-primary" @click="updateGroupName">Change Group Name</button>
            </div>
        </div>
    </div>
    
    

</template>

<style>
.user-avatar {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  margin-right: 10px;
}
</style>



