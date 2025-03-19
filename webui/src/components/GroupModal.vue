<script>    
    export default{

        props: {
          show: Boolean,
          title: String,
        },

        data: function() {
            return {
                groupname: "",
                searchTerm: "",
                searchResults: [],
                selectedUsers: [],
                errormsg: "",
                owner: sessionStorage.username,

            }
            
        },

        methods:{
            async searchUsers() {
                if (!this.searchTerm) {
                    this.searchResults = [];
                    return;
                }

                try {
                    const response = await this.$axios.get(`/user/${sessionStorage.userID}`, {
                    params: { search: this.searchTerm },
                    headers: { 'Authorization': sessionStorage.token }
                    });

                    this.searchResults = response.data;
                } catch (e) {
                    this.errormsg = e.toString();
                }
            },
            selectUser(user) {
                const index = this.selectedUsers.findIndex(u => u.username === user.username);
                if (index === -1) {
                    this.selectedUsers.push(user);
                }else {
                    this.selectedUsers.splice(index, 1);
                }

            },

            closeModal() {
                this.searchTerm = "";
                this.searchResults = [];
                this.selectedUsers = [];
                this.groupName = "";
                this.$emit("close");
            },

            async createGroup(){
              if (this.selectedUsers.length === 0) {
                alert("Seleziona almeno un utente per avviare la chat di gruppo.");
                return;
              }

              const members = this.selectedUsers.map(user => user.username);
              localStorage.setItem("groupMembers", JSON.stringify(members)); // Salviamo gli ID nel localStorage
              console.log("Creating group chat as user:", sessionStorage.userID);
              console.log("Participants:", members);
              try {
                let response = await this.$axios.post(`/user/${sessionStorage.userID}/group`, {
                  groupname: this.groupName,
                  usernamelist: members
                }, {
                  headers: { 'Authorization': sessionStorage.token }
                });
                console.log("Group chat created:", response.data);
                let conviID= response.data.conversation.conversationID;
                localStorage.groupname = this.groupName;
                localStorage.groupId = response.data.conversation.groupID;
                localStorage.groupMembers = JSON.stringify(response.data.groupUsers);
                localStorage.photo = response.data.group.photo;
                this.$router.push(`/conversation/${conviID}`);

              }catch (e) {
                this.errormsg = e.toString();
              }

            }
    }
}

</script>
<template>
    <div v-if="show" class="modal-overlay">
  <div class="modal-content">
    <h3>Create Group Chat</h3>

    <!-- Input per il nome del gruppo -->
    <input
      type="text"
      v-model="groupName"
      placeholder="Enter group name"
      class="form-control mb-3"
    />

    <!-- Mostra gli utenti selezionati -->
    <div v-if="selectedUsers.length" class="selected-users">
      <span v-for="user in selectedUsers" :key="user.userID" class="selected-user">
        {{ user.username }}
        <span class="remove-user" @click="selectUser(user)">✖</span>
      </span>
    </div>

    <!-- Barra di ricerca per aggiungere altri utenti -->
    <input
      type="text"
      v-model="searchTerm"
      placeholder="Search for users"
      class="form-control mb-3"
      @input="searchUsers"
    />

    <!-- Lista utenti trovati -->
    <ul class="list-group">
      <li
        v-for="user in searchResults.filter(u => u.username !== owner)"
        :key="user.id"
        class="list-group-item"
        :class="{ 'selected-user': selectedUsers.some(u => u.userID === user.userID) }"
        @click="selectUser(user)"
      >
        {{ user.username }}
      </li>
    </ul>

    <div class="modal-actions">
      <button class="btn btn-primary" @click="createGroup">Select</button>
      <button class="btn btn-secondary" @click="closeModal">Cancel</button>
    </div>
  </div>
</div>


</template>
<style>

</style>