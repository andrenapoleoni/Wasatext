<script>
export default{

    props:{
        show: Boolean,
        title: String,
        somedata: Array,

    },

    data:function() {
        return {
            searchTerm: "",
            searchResults: [],
            errormsg:"",
            owner: sessionStorage.username,
            privateChatUserId: "",
            selectedUserId: null,


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
            this.privateChatUserId = user.userID; // Set user ID for the chat
            this.selectedUserId = user.userID; // Highlight the selected user
        },
        async startPrivateChat() {
            const existingChat = this.somedata.find(chat=> chat.user.userID === this.selectedUserId);
            localStorage.groupname = null;
            localStorage.groupId = 0;
            localStorage.userID=this.selectedUserId;
            localStorage.username=this.searchResults.find(user => user.userID === this.selectedUserId).username;
            localStorage.photo=this.searchResults.find(user => user.userID === this.selectedUserId).photo;
            if (existingChat) {
            // Redirect to the existing chat
                this.$router.push(`/conversation/${existingChat.conversation.conversationID}`);
            } else {
                // Open a new chat view without creating a conversation yet
                this.$router.push(`/conversation/null`);
            }
        },

        closeModal() {
                this.searchTerm = "";
                this.searchResults = [];
                this.selectedUsers = [];
                this.$emit("close");
        },
    }
}




</script>
<template>
    <div v-if="show" class="modal-overlay">
      <div class="modal-content">
        <h3>Start Private Chat</h3>

        <!-- Search bar for user search -->
        <input
          type="text"
          v-model="searchTerm"
          placeholder="Search for a user"
          class="form-control mb-3"
          @input="searchUsers"
        />

        <!-- List of matching users -->
        <ul class="list-group">
          <li
            v-for="user in searchResults.filter(u => u.username !== owner)"
            :key="user.id"
            class="list-group-item"
            :class="{ 'selected-user': user.userID === selectedUserId }"
            @click="selectUser(user)"
          >
            {{ user.username }}
          </li>
        </ul>

        <div class="modal-actions">
          <button class="btn btn-primary" @click="startPrivateChat">Start chat</button>
          <button class="btn btn-secondary" @click="closeModal">Cancel</button>
        </div>
      </div>
    </div>




</template>

<style>


</style>