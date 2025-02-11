<script>
export default {
  data: function() {
    return {
      owner: sessionStorage.username, 
      ownerID: sessionStorage.userID,   
      errormsg: null,
      some_data: [],
      showDropdown: false,
      showPrivateChatModal: false,
      privateChatUserId: "", // Selected user ID
      searchTerm: "",
      searchResults: [],
      existingChatWarning: "",
      selectedUserId: null, // Track the selected user
      showGroupChatModal: false,
      groupChatMembers: [],
      groupName: "",

    };
  },
  methods: {
    async getMyconversations() {
      this.errormsg = null;
      console.log("USer:", sessionStorage.userID,"token:", sessionStorage.token);
      try {
        let response = await this.$axios.get(`/user/${sessionStorage.userID}/conversation`, {
          headers: { 'Authorization': sessionStorage.token }
      });

      this.some_data = response.data.map(chat => {
      return {
        ...chat,
        isGroup: !!chat.group, // Se `group` esiste, allora è una chat di gruppo
      };
      });
      } catch (e) {
        this.errormsg = e.toString();
  }
},

    handleConversationClick(response) {
      localStorage.clear();
    console.log("Conversation ID cliccato:", response.conversation.conversationID);
    console.log("Username cliccato:", response.user.username);
    console.log("Group ID cliccato:", response.group.groupID);
    console.log("Group name cliccato:", response.group.groupname);
    console.log("Group participants cliccato:", response.groupUsers);
    if (response.group.groupID!==0) {
      console.log("Group chat:", response.group.groupname);
      localStorage.groupname = response.group.groupname;
      localStorage.groupId = response.group.groupID;
      localStorage.groupMembers = JSON.stringify(response.groupUsers);
      localStorage.photo = response.group.photo;


     
    }else {
      console.log("Private chat:", response.user.username);
      
      localStorage.groupId = 0;
      localStorage.username = response.user.username;
      localStorage.photo = response.user.photo;
      
    }
    this.$router.push(`/conversation/${response.conversation.conversationID}`);
  },
  toggleDropdown(event) {
      this.showDropdown = !this.showDropdown;
      if (this.showDropdown) {
        event.stopPropagation();
      }
    },
  closeDropdown() {
      this.showDropdown = false;
    },
    newGroupChat() {
      this.showGroupChatModal = true;
      this.searchTerm = "";
      this.groupName = "";
      this.searchResults = [];
      this.groupChatMembers = [];
      
    },
    newPrivateChat() {
      this.showPrivateChatModal = true;
      this.searchTerm = ""; // Clear search term when opening modal
      this.searchResults = []; // Clear previous search results
      this.selectedUserId = null; // Reset selected user
    },
    async startPrivateChat() {
      
      const existingChat = this.some_data.find(chat => {
          // Verifica che "chat.participants" sia un array valido prima di usare includes
          return Array.isArray(chat.participants) && chat.participants.includes(this.privateChatUserId);
        });
      
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

  async startGroupChat() {
      if (this.groupChatMembers.length === 0) {
        alert("Seleziona almeno un utente per avviare la chat di gruppo.");
        return;
      }

      const members = this.groupChatMembers.map(user => user.username);
      localStorage.setItem("groupMembers", JSON.stringify(members)); // Salviamo gli ID nel localStorage
      console.log("Creating group chat as user:", sessionStorage.userID);
      console.log("Participants:", members);
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
  },
    closeModal() {
      this.showPrivateChatModal = false;
    },
    async searchUsers() {
      if (!this.searchTerm) {
        this.searchResults = [];
        return;
      }

      try {
        const response = await this.$axios.get(`/user`, {
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

    selectGroupUser(user) {
      const index = this.groupChatMembers.findIndex(u => u.username === user.username);
      if (index === -1) {
      this.groupChatMembers.push(user);
      } else {
        this.groupChatMembers.splice(index, 1);
      }
},

    
},
  mounted() {
    if (!sessionStorage.token) {
      this.$router.push("/login");
      return;
    }
    this.getMyconversations();
    document.addEventListener("click", this.closeDropdown);
  },
  beforeUnmount() {
    document.removeEventListener("click", this.closeDropdown);
  }

};


</script>

<template>
  <div>
    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <h1 class="h2">Home page</h1>
      <div class="btn-toolbar mb-2 mb-md-0">
        <div class="btn-group me-2">
          <button type="button" class="btn btn-sm btn-outline-secondary" @click="mounted">Refresh</button>
          <button type="button" class="btn btn-sm btn-outline-secondary" @click="exportList">Export</button>
        </div>
        <div class="btn-group me-2 position-relative" @click.stop>
          <button type="button" class="btn btn-sm btn-outline-primary" @click="toggleDropdown">New Chat</button>
          <div v-if="showDropdown" class="dropdown-menu show dropdown-menu-end">
            <button class="dropdown-item" @click="newGroupChat">New Group Chat</button>
            <button class="dropdown-item" @click="newPrivateChat">New Private Chat</button>
          </div>
        </div>
      </div>
    </div>

    <!-- Error message -->
    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

    <!-- Warning for existing chat -->
    <div v-if="existingChatWarning" class="alert alert-warning">
      {{ existingChatWarning }}
    </div>

    <!-- Modal for private chat -->
    <div v-if="showPrivateChatModal" class="modal-overlay">
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

    <div 
  class="conversations" 
  v-for="response in some_data" 
  :key="response.conversation.conversationID"
  @click="handleConversationClick(response)"
  style="cursor: pointer;"
>
  <!-- Se esiste un gruppo, mostra il nome del gruppo -->
  <span v-if="response.group && response.group.groupname">{{ response.group.groupname }}</span>
  
  <!-- Altrimenti, mostra il nome dell'utente -->
  <span v-else>{{ response.user.username }}</span>
  
  <hr>
</div>
  </div>
 <!-- Modale per la chat di gruppo -->
<div v-if="showGroupChatModal" class="modal-overlay">
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
    <div v-if="groupChatMembers.length" class="selected-users">
      <span v-for="user in groupChatMembers" :key="user.userID" class="selected-user">
        {{ user.username }}
        <span class="remove-user" @click="selectGroupUser(user)">✖</span>
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
        :key="user.userID"
        class="list-group-item"
        :class="{ 'selected-user': groupChatMembers.some(u => u.userID === user.userID) }"
        @click="selectGroupUser(user)"
      >
        {{ user.username }}
      </li>
    </ul>

    <div class="modal-actions">
      <button class="btn btn-primary" @click="startGroupChat">Start Group Chat</button>
      <button class="btn btn-secondary" @click="showGroupChatModal = false">Cancel</button>
    </div>
  </div>
</div>


</template>

<style>
.position-relative {
  position: relative;
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  right: 0; 
  z-index: 1000;
  display: block;
  float: left;
  min-width: 10rem;
  padding: .5rem 0;
  margin: .125rem 0 0;
  font-size: 1rem;
  color: #212529;
  text-align: left;
  background-color: #fff;
  border: 1px solid rgba(0, 0, 0, .15);
  border-radius: .25rem;
  box-shadow: 0 .5rem 1rem rgba(0, 0, 0, .175);
}

.dropdown-item {
  display: block;
  width: 100%;
  padding: .25rem 1.5rem;
  clear: both;
  font-weight: 400;
  color: #212529;
  text-align: inherit;
  white-space: nowrap;
  background-color: transparent;
  border: 0;
  cursor: pointer;
}

.dropdown-item:hover {
  color: #fff;
  background-color: #007bff;
}

/* Modal Styles */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1050;
}

.modal-content {
  background: #fff;
  padding: 20px;
  border-radius: 8px;
  width: 400px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  text-align: center;
}

.modal-actions {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}

/* List of matching users */
.list-group {
  max-height: 200px;
  overflow-y: auto;
  margin-bottom: 10px;
}

.list-group-item {
  cursor: pointer;
}

.selected-user {
  background-color: #007bff; /* Highlight color */
  color: white; /* Text color for contrast */
  font-weight: bold; /* Optional: make the text bold */
}

.selected-user:hover {
  background-color: #0056b3; /* Darker shade on hover */
}

.selected-users {
  display: flex;
  flex-wrap: wrap;
  gap: 5px;
  margin-bottom: 10px;
}

.selected-user {
  background-color: #007bff;
  color: white;
  padding: 5px 10px;
  border-radius: 15px;
  display: flex;
  align-items: center;
}

.remove-user {
  margin-left: 8px;
  cursor: pointer;
  font-weight: bold;
}

</style>